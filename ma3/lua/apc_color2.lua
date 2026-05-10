-- Plugin for controlling ColorSequences via an APC Mini Mk2 MIDI controller

-- @Author: Sebastian Dolfen
-- @Date: 2024-06-01
-- @Description: This plugin polls the status and color of ColorSequences on the MA3 console and sends updates via OSC to a supplementary executable that translates the OSC messages to MIDI. It also detects page changes and sends updates accordingly.

-- #region Libraries
local socket_ok, socket = pcall(require, "socket")

-- #endregion

-- #region Configuration:
local osc_config = {
    remote_ip = "127.0.0.1", -- IP address of the machine running the supplementary executable
    remote_port = 8000,    -- Port number for OSC communication
    local_ip = "0.0.0.0",  -- IP address for receiving OSC messages
    local_port = 8001      -- Port number for receiving OSC messages from the supplementary executable
}

-- #endregion

-- #region MAIN PLUGIN LOGIC

local APCColor = {}

APCColor.state = {
    enabled = false,
    connected = false,
    dev1 = {
        page = 1,
        sequences = {},
        sequence_start_lines = {}, -- Map of sequence name -> start page and line
    },
    dev2 = {
        page = 1,
        sequences = {},
    },
}

APCColor.Setup = function()
    local sequences = DataPool().Sequences:Children()

    Filter_inplace(sequences, function(seq)
        return string.starts(seq.Name, "APCColor")
    end)

    -- Map of sequence name -> start line index (1-based)
    APCColor.state.dev1.sequence_start_lines = {}
    APCColor.state.dev1.sequences = sequences

    -- Placement rules:
    -- - 8 buttons per line
    -- - A sequence needs one button per cue (excluding OffCue and CueZero)
    -- - Next sequence always starts on a new line, even if previous didn't fill the line
    local buttons_per_line = 8
    local current_line = 1
    local current_page = 1

    for _, seq in ipairs(sequences) do
        local seq_name = seq.Name or "<unnamed>"

        -- Determine how many lines this sequence consumes
        local cue_count = Count_display_cues(seq)
        local lines_used = math.ceil((cue_count > 0 and cue_count or 1) / buttons_per_line)

        -- Check if lines_used would exceed the current page (8 lines per page)
        if current_line + lines_used - 1 > 8 then
            current_page = current_page + 1
            current_line = 1
        end

        -- Record the starting line for this sequence
        APCColor.state.dev1.sequence_start_lines[seq_name] = { page = current_page, line = current_line }

        -- Advance to the next starting line for the following sequence
        current_line = current_line + lines_used
    end

    -- Optional: log the computed mapping for diagnostics
    if APCColor.state.dev1.sequence_start_lines then
        for name, line in pairs(APCColor.state.dev1.sequence_start_lines) do
            Printf(string.format("APCColor: Sequence '%s' starts at page %d, line %d", tostring(name),
                tonumber(line.page),
                tonumber(line.line)))
        end
    end

    Printf("APCColor: Plugin setup complete, entering main loop")
    APCColor.mainloop() -- Start the main loop to keep the plugin running
end

APCColor.mainloop = function()
    while APCColor.state.enabled do
        coroutine.yield(0.1)
    end
end

-- #endregion

-- #region Utility functions

function Filter_inplace(arr, func)
    local new_index = 1
    local size_orig = #arr
    for old_index, v in ipairs(arr) do
        if func(v, old_index) then
            arr[new_index] = v
            new_index = new_index + 1
        end
    end
    for i = new_index, size_orig do arr[i] = nil end
end

function CleanNils(t)
    local ans = {}
    for _, v in pairs(t) do
        ans[#ans + 1] = v
    end
    return ans
end

function string.starts(String, Start)
    return string.sub(String, 1, string.len(Start)) == Start
end

-- #region MA3-specific utility functions

-- Helper: count cues that require buttons (exclude OffCue and CueZero)
function Count_display_cues(sequence)
    local cnt = 0
    local cues = sequence:Children()
    for _, cue in ipairs(cues) do
        local n = cue and cue.Name or ""
        cue:Dump()
        if n ~= "OffCue" and n ~= "CueZero" then
            cnt = cnt + 1
        end
    end
    return cnt
end

-- #endregion
-- #endregion

local function maintoggle()
    if APCColor.state.enabled then
        Printf("APCColor: Disabling plugin")
        APCColor.state.enabled = false
    else
        Printf("APCColor: Enabling plugin")
        APCColor.state.enabled = true
        APCColor.Setup()
    end
end

return maintoggle
