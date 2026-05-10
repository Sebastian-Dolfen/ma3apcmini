-- APC Color Plugin for GrandMA3
-- Maps APC mini mk2 pads/pages to Sequence cues and drives LEDs via OSC bridge
-- Note: This plugin uses LuaSocket for UDP. Ensure 'socket' is available to MA3 Lua.

local socket_ok, socket = pcall(require, 'socket')

local APCColor = {}

-- Flag to control the main loop
local enabled = false

-- Logging helpers using MA3's Printf/ErrPrintf if available
local function logf(fmt, ...)
  if Printf then
    if select('#', ...) > 0 then
      Printf(string.format(fmt, ...))
    else
      Printf(fmt)
    end
  end
end

local function errf(fmt, ...)
  if ErrPrintf then
    if select('#', ...) > 0 then
      ErrPrintf(string.format(fmt, ...))
    else
      ErrPrintf(fmt)
    end
  elseif Printf then
    if select('#', ...) > 0 then
      Printf('ERROR: ' .. string.format(fmt, ...))
    else
      Printf('ERROR: ' .. fmt)
    end
  end
end

-- Function to pause the script for a given number of seconds
local function pause(seconds)
  coroutine.yield(seconds)
end

-- Config defaults
local cfg = {
  seqId = 1,
  currentPage = 1,
  remote_ip = '127.0.0.1', -- Bridge in
  remote_port = 8001,
  local_ip = '0.0.0.0',    -- Plugin in
  local_port = 8000,
  dev = 1,                 -- Device id to target on the bridge (1 or 2). If 1, legacy 3-arg OSC is used.
}

-- State
local state = {
  selected = nil,  -- preloaded cue index (global index, 1..N) or nil
  active = nil,    -- active cue index read from Sequence
  cues = {},       -- map[cueIndex] = { color = '#RRGGBB' }
  total = 0,
}

-- UDP sockets
local udp_out, udp_in

-- Utility: pad OSC string to 4-byte boundary
local function osc_pad4(s)
  local rem = (#s + 1) % 4
  if rem == 0 then return s .. string.char(0) end
  local pad = 4 - rem
  return s .. string.char(0) .. string.rep('\0', pad)
end

-- Utility: int32 big-endian
local function be32(i)
  local b1 = math.floor(i / 16777216) % 256
  local b2 = math.floor(i / 65536) % 256
  local b3 = math.floor(i / 256) % 256
  local b4 = i % 256
  return string.char(b1, b2, b3, b4)
end

-- Encode OSC message with address and arguments (ints and strings only)
local function osc_msg(address, ...)
  local parts = {}
  table.insert(parts, osc_pad4(address))
  local types = {','}
  local args_bin = {}
  local args = {...}
  for _,v in ipairs(args) do
    if type(v) == 'number' then
      table.insert(types, 'i')
      table.insert(args_bin, be32(math.floor(v)))
    else
      table.insert(types, 's')
      table.insert(args_bin, osc_pad4(tostring(v)))
    end
  end
  table.insert(parts, osc_pad4(table.concat(types)))
  table.insert(parts, table.concat(args_bin))
  return table.concat(parts)
end

local function cueIndexFor(page, padIndex0)
  return (page - 1) * 64 + padIndex0 + 1
end

local function pageForCueIndex(idx)
  return math.floor((idx - 1) / 64) + 1
end

-- Read active cue index from Sequence; fallback to selected if unknown
local function read_active_cue()
  -- TODO: integrate with MA3 API to read the current cue index of the sequence.
  -- Placeholder: keep previous value.
  return state.active
end

-- Read appearance color for specified cue index; returns '#RRGGBB'
local function read_cue_color(idx)
  -- TODO: integrate with MA3 API to access Appearance of Cue <idx> of Sequence <seqId>
  -- Placeholder to black; user should assign appearances and run Rescan after integrating API.
  return '#000000'
end

local function send_led_for_pad(padIndex0, ledState, hex)
  if not udp_out then
    errf('send_led_for_pad: udp_out not initialized')
    return
  end
  local msg
  if cfg.dev and cfg.dev > 1 then
    msg = osc_msg('/Pad', cfg.dev, padIndex0, ledState, hex)
  else
    msg = osc_msg('/Pad', padIndex0, ledState, hex)
  end
  udp_out:sendto(msg, cfg.remote_ip, cfg.remote_port)
end

local function send_page_led()
  if not udp_out then
    errf('send_page_led: udp_out not initialized')
    return
  end
  local msg
  if cfg.dev and cfg.dev > 1 then
    msg = osc_msg('/Page', cfg.dev, cfg.currentPage)
  else
    msg = osc_msg('/Page', cfg.currentPage)
  end
  udp_out:sendto(msg, cfg.remote_ip, cfg.remote_port)
end

local function render_leds_for_page(page)
  for pad = 0, 63 do
    local idx = cueIndexFor(page, pad)
    local color = '#000000'
    local st = 0
    if idx <= state.total and state.cues[idx] then
      color = state.cues[idx].color or '#000000'
      if state.active and state.active == idx then
        st = 2 -- active
      elseif state.selected and state.selected == idx then
        st = 1 -- preloaded
      else
        st = 0 -- idle solid
      end
    else
      color = '#000000'
      st = 0
    end
    send_led_for_pad(pad, st, color)
  end
  send_page_led()
end

local function rescan()
  -- Discover number of cues and their appearance colors
  -- TODO: hook to MA3 object model to list cues of Sequence cfg.seqId
  -- Placeholder: assume 64 cues with default black unless previously known
  local prev_total = state.total or 0
  local assumed = math.max(state.total, 64)
  state.total = assumed
  for i = 1, state.total do
    if not state.cues[i] then
      state.cues[i] = { color = read_cue_color(i) }
    end
  end
  state.active = read_active_cue()
  render_leds_for_page(cfg.currentPage)
  if state.total ~= prev_total then
    logf('Rescan complete: sequence=%d total_cues=%d page=%d', cfg.seqId, state.total, cfg.currentPage)
  end
end

local function set_page(p)
  if p < 1 then p = 1 end
  cfg.currentPage = p
  render_leds_for_page(cfg.currentPage)
  -- Also send current page indicator to bridge explicitly to clear any stale logos
  send_page_led()
  logf('Set page: %d', cfg.currentPage)
end

local function commit_selected()
  if not state.selected then return end
  -- Preferred: Load then Go+
  local cmdLoad = string.format('Load Sequence %d Cue %d', cfg.seqId, state.selected)
  local cmdGo = string.format('Go+ Sequence %d', cfg.seqId)
  if Cmd then
    Cmd(cmdLoad)
    Cmd(cmdGo)
    logf('Committed cue: seq=%d cue=%d', cfg.seqId, state.selected)
  else
    errf('Cmd not available; cannot commit cue seq=%d cue=%d', cfg.seqId, state.selected)
  end
  state.active = state.selected
  state.selected = nil
  render_leds_for_page(cfg.currentPage)
end

local function handle_pad(padIndex0, value)
  if value ~= 1 then return end -- only on press
  local idx = cueIndexFor(cfg.currentPage, padIndex0)
  if idx > state.total then return end
  if state.selected == idx then
    state.selected = nil
    logf('Pad press: pad=%d deselect cue=%d', padIndex0, idx)
  else
    state.selected = idx
    logf('Pad press: pad=%d select cue=%d', padIndex0, idx)
  end
  render_leds_for_page(cfg.currentPage)
end

local function handle_page(p)
  set_page(p)
end

local function handle_trigger(_)
  commit_selected()
end

-- Minimal OSC message parser for our inputs: expects address in first packet bytes
local function parse_osc_packet(pkt)
  if not pkt or #pkt < 4 then return nil end
  -- Address is a null-terminated string starting at 1
  local addr = pkt:match('([^%z]+)') or ''
  local args = {}
  -- Very limited parser: supports 2 or 3 trailing int32 for /Pad, and 1 or 2 for /Page
  local function u32(s)
    local b1,b2,b3,b4 = s:byte(1,4)
    return b1*16777216 + b2*65536 + b3*256 + b4
  end
  if addr == '/Pad' then
    if #pkt >= 12 then
      -- try 3 ints: dev, pad, state
      local a = pkt:sub(#pkt-11, #pkt-8)
      local b = pkt:sub(#pkt-7,  #pkt-4)
      local c = pkt:sub(#pkt-3,  #pkt)
      table.insert(args, u32(a))
      table.insert(args, u32(b))
      table.insert(args, u32(c))
    elseif #pkt >= 8 then
      -- fallback 2 ints: pad, state
      local b = pkt:sub(#pkt-7, #pkt-4)
      local c = pkt:sub(#pkt-3, #pkt)
      table.insert(args, u32(b))
      table.insert(args, u32(c))
    end
  elseif addr == '/Page' or addr == '/Trigger' then
    if #pkt >= 8 then
      -- possibly dev, page
      local a = pkt:sub(#pkt-7, #pkt-4)
      local b = pkt:sub(#pkt-3, #pkt)
      table.insert(args, u32(a))
      table.insert(args, u32(b))
    elseif #pkt >= 4 then
      local b = pkt:sub(#pkt-3, #pkt)
      table.insert(args, u32(b))
    end
  end
  return { address = addr, args = args }
end

local function poll_once()
  if not udp_in then return end
  udp_in:settimeout(0)
  local pkt, rip, rport = udp_in:receivefrom()
  if not pkt then
    -- ignore timeouts, log other socket errors
    local _, err = udp_in:receive()
    if err and err ~= 'timeout' then
      errf('UDP receive error: %s', tostring(err))
    end
    return
  end
  -- Diagnostic: log first bytes and raw address string
  local function hexdump(s, max)
    max = max or 64
    local out = {}
    local n = math.min(#s, max)
    for i=1,n do
      out[#out+1] = string.format('%02X', s:byte(i))
    end
    if #s > max then out[#out+1] = '…' end
    return table.concat(out, ' ')
  end
  local raw_addr = pkt:match('([^%z]+)') or ''
  logf('UDP pkt %dB from %s:%s addr_guess="%s" bytes=%s', #pkt, tostring(rip), tostring(rport), raw_addr, hexdump(pkt, 48))
  local msg = parse_osc_packet(pkt)
  if not msg then return end
  logf('OSC recv: %s from %s:%s', msg.address or 'nil', tostring(rip), tostring(rport))
  if msg.address == '/Pad' then
    if #msg.args >= 3 then -- dev, pad, state
      handle_pad(tonumber(msg.args[2]) or 0, tonumber(msg.args[3]) or 0)
    else -- pad, state
      handle_pad(tonumber(msg.args[1]) or 0, tonumber(msg.args[2]) or 0)
    end
  elseif msg.address == '/Page' then
    if #msg.args >= 2 then -- dev, page
      handle_page(tonumber(msg.args[2]) or 1)
    else
      handle_page(tonumber(msg.args[1]) or 1)
    end
  elseif msg.address == '/Trigger' then
    handle_trigger(1)
  else
    logf('Unhandled OSC address: %s', tostring(msg.address))
  end
end

-- Public API
function APCColor.Setup(params)
  -- params like: "seq=5 page=1 remote=127.0.0.1:8001 local=0.0.0.0:8000"
  if type(params) == 'string' then
    for k,v in string.gmatch(params, '(%w+)=([^%s]+)') do
      if k == 'seq' then cfg.seqId = tonumber(v) or cfg.seqId
      elseif k == 'page' then cfg.currentPage = tonumber(v) or cfg.currentPage
      elseif k == 'remote' then local ip,port = v:match('([^:]+):(%d+)'); if ip and port then cfg.remote_ip, cfg.remote_port = ip, tonumber(port) end
      elseif k == 'local' then local ip,port = v:match('([^:]+):(%d+)'); if ip and port then cfg.local_ip, cfg.local_port = ip, tonumber(port) end
      elseif k == 'dev' then cfg.dev = tonumber(v) or cfg.dev
      end
    end
  end
  if socket_ok then
    udp_out = socket.udp()
    udp_in = socket.udp()
    local ok, bind_err = pcall(function()
      udp_in:setsockname(cfg.local_ip, cfg.local_port)
    end)
    if not ok then
      errf('UDP bind failed on %s:%d: %s', tostring(cfg.local_ip), tonumber(cfg.local_port), tostring(bind_err))
    end
    logf('Setup: seq=%d page=%d remote=%s:%d local=%s:%d dev=%d socket_ok=%s',
      cfg.seqId, cfg.currentPage, tostring(cfg.remote_ip), tonumber(cfg.remote_port), tostring(cfg.local_ip), tonumber(cfg.local_port), tonumber(cfg.dev or 0), tostring(socket_ok))
    -- Log the local address/port used by udp_out (may be assigned after first send)
    local out_ip, out_port = udp_out:getsockname()
    logf('udp_out local address (pre-send): %s:%s', tostring(out_ip), tostring(out_port))
  end
  if not socket_ok then
    errf('LuaSocket not available; OSC I/O disabled')
  end
  rescan()
  if socket_ok and udp_out then
    -- After initial LED/page sends in rescan, udp_out should have an ephemeral local port
    local out_ip2, out_port2 = udp_out:getsockname()
    logf('udp_out local address (post-send): %s:%s', tostring(out_ip2), tostring(out_port2))
  end
end

function APCColor.Rescan()
  while enabled do
    rescan()
    pause(0.1)
  end
end

function APCColor.SetPage(p)
  set_page(tonumber(p) or 1)
end

function APCColor.Commit()
  commit_selected()
end

-- Non-blocking poll; call periodically from a Macro or timer to handle OSC input
function APCColor.Poll()
  poll_once()
end

-- Explicit shutdown call to inform the bridge to restore logos for this device
function APCColor.Shutdown()
  if not udp_out then
    errf('Shutdown: udp_out not initialized')
    return
  end
  local msg
  if cfg.dev and cfg.dev > 1 then
    msg = osc_msg('/Shutdown', cfg.dev)
  else
    msg = osc_msg('/Shutdown')
  end
  udp_out:sendto(msg, cfg.remote_ip, cfg.remote_port)
  logf('Shutdown sent to %s:%d (dev=%d)', tostring(cfg.remote_ip), tonumber(cfg.remote_port), tonumber(cfg.dev or 0))
end

-- Export to global so other macros can call APCColor.* if needed
_G.APCColor = APCColor


local function maintoggle() 
  if enabled then
    -- Disable the loop and clear history when turning off
    enabled = false
    logf('APCColor disabled')
  else
    -- Enable the loop and start it when turning on
    enabled = true
    logf('APCColor enabled')
    APCColor.Setup()
  end
end

-- Return the toggle function for external control
return maintoggle
