# v2 Plan — APC mini mk2 ↔ grandMA3 bridge

## Context

The current bridge is a Node.js → Go fork mid-rewrite. The Go bridge (`go/cmd/apc-mini-bridge/main.go`) is functional but hardcoded to exactly two devices, the new MA3 plugin `ma3/lua/apc_color2.lua` is ~30% complete (sequence layout works, OSC/LED/MA3-reads all stubbed), and the legacy `apc_color.lua` is single-device, single-cue, with stubbed reads. v2 must support:

1. **Multiple AKAIs** with per-device *role* (currently dev1/dev2 hardcoded).
2. **Color-AKAI role** — pads represent cues from named sequences. Pressing a pad **loads** that cue (`Load Sequence X Cue Y`). Loaded pads get a distinct LED state. A dedicated **fire-all-loaded** button executes `Go+ Sequence X` for every sequence with a loaded cue.
3. **Wing-AKAI role** — second AKAI is a transparent MA3 fader-wing surface (faders + buttons drive a configurable executor block, programmable inside MA3 like a normal wing).
4. Pad LED color auto-derived from the cue's `Appearance` (RGB) on the MA3 side.

Two cross-cutting issues must be fixed during this work because they're in the way: (a) the OSC contract is overloaded (`/Pad` is used both directions and silently accepts variable arity) and (b) the bridge has unsynchronized concurrent access to `apcDevice.apc` from scanner + OSC handler goroutines.

The audit feeding this plan: three Explore agents (Go bridge, MA3 Lua, origin/config/scripts) and an independent Codex (gpt-5.5) review.

---

## Architecture

### Hardware reference (APC mini mk2, confirmed from photo)

| Element | Count | MIDI | Notes |
|---|---|---|---|
| RGB grid pads | 8×8 = 64 | Note 0..63 | Multi-color, palette-driven via velocity |
| Top row "FADER CTRL" buttons (VOLUME, PAN, SEND, DEVICE, ↑, ↓, ←, →) | 8 | Note 100..107 | Single-color (red) |
| SHIFT button (top-right) | 1 | Note 122 | Single-color. Node.js origin treated this as "BO" — on mk2 it's labelled SHIFT; treat as a generic modal/extra button. |
| "SCENE LAUNCH" right column (CLIP STOP, SOLO, MUTE, REC ARM, SELECT, DRUM, NOTE, STOP ALL CLIPS) | 8 | Note 112..119 | Single-color. Used for page selection in current plugin. |
| Faders 1..8 | 8 | CC 48..55 | 0..127 |
| Master fader (rightmost) | 1 | CC 56 | 0..127 |

The grid LEDs are encoded by MIDI channel (Solid / Pulse-1/8 / Pulse-1/16 / true blink etc., see `docs/midi_and_apc_mini_mk2/05_led_protocol.md`). The single-color buttons accept velocity 0..127 mapping to brightness/blink modes.

### Per-device roles

Bridge config gains a `devices` array; each entry has `role: color | wing`. The Go bridge stays a generic MIDI ↔ OSC translator and is **role-agnostic** — it only knows `dev` ids, ports, and MIDI bytes. The MA3 plugin reads role from its own startup args and dispatches per device:

- `color` device → custom plugin UX (sequence layout, load tracking, fire-all). LED commands flow plugin → bridge over the v2 OSC schema.
- `wing` device → plugin emits MA3 built-in OSC (`/Page1/Fader…`, `/Page1/Key…`, `/cmd "Master 2.1 At …"`) using `SendOSC` to MA3's OSC config slot, *and* echoes the resulting executor state back to the bridge as LED commands. The wing's executor block is configurable (default: `wing_offset = 600` → faders 601-608, top-row keys 601-608 etc.).

This means **all multi-device complexity lives in the plugin**, not the bridge. The bridge just routes by `dev` id.

### State ownership (per Codex)

- **Loaded-cue set lives in the plugin only.** The bridge does not know what a cue is. If the bridge restarts, the plugin (which survives) repaints the entire LED state for each device that comes back online.
- The plugin tracks `loaded[seq_id] = cue_no` (one entry per sequence — MA3 load semantics already overwrite, so this matches).
- On plugin start, all loads are cleared (we explicitly do *not* persist across plugin reloads; `_G` reload is unreliable per `docs/grandMA3_practical_guide/08_gotchas.md:51`).

### OSC contract v2 (versioned, namespaced)

Codex's "must": don't extend `/Pad` in place. Add new namespaced routes; keep v1 routes accepted forever for legacy `apc_color.lua`.

**Bridge → Plugin** (input events; `dev` is mandatory and first):

| Address | Type tag | Args | Notes |
|---|---|---|---|
| `/apc/v2/input/grid` | `iii` | dev, pad (0..63), value (1=press, 0=release) | Replaces `/Pad` outbound |
| `/apc/v2/input/button` | `iii` | dev, note (100..127), value | Side col, top row, BO — plugin decides |
| `/apc/v2/input/fader` | `iii` | dev, cc (48..56), value (0..127) | New; replaces nothing — was missing entirely |
| `/apc/v2/hello` | `i` | dev | Bridge announces device output came online; plugin must repaint |

**Plugin → Bridge** (LED + lifecycle):

| Address | Type tag | Args | Notes |
|---|---|---|---|
| `/apc/v2/led/grid` | `iiis` | dev, pad (0..63), state (bitmask), `#RRGGBB` | state bits: 1=loaded, 2=active, so 3=loaded+active, 0=idle |
| `/apc/v2/led/button` | `iiis` | dev, note (112..119 etc.), state, `#RRGGBB` | Page button + accessories |
| `/apc/v2/shutdown` | `i` | dev | Plugin going away — bridge restores logo |

Bridge maps `state` to MIDI channel. **Plugin never thinks in MIDI channels** — that's the bridge's job (and lets us fix the misnamed channel constants without touching the plugin).

| state | LED behavior | MIDI channel |
|---|---|---|
| 0 (idle) | Solid | `ChannelSolid` (6) |
| 1 (loaded) | Pulse 1/8 (fast pulse) | `ChannelPulse8` (currently misnamed `ChannelFastBlink` = 8) |
| 2 (active) | Pulse 1/16 (slow pulse) | `ChannelPulse16` (currently misnamed `ChannelSlowBlink` = 7) |
| 3 (loaded+active) | True blink | `ChannelBlink16` (11) — distinct, hardware-true blink |

Color hex is encoded to APC mk2 palette velocity via the existing LUT (`go/internal/color/map.go`). The three LSB typos at velocities 67/70/72 get fixed.

**Legacy v1 routes** (`/Pad`, `/Page`, `/Trigger`, `/Shutdown`) keep working in the bridge for the legacy `apc_color.lua`, but `apc_color2.lua` only emits v2.

### Per-device concurrency (per Codex "must")

`apcDevice` gains a `sync.Mutex`. The scanner replacing `apc *APCMK2` and the OSC handlers using it must coordinate. Wrap every `target.apc.…` call with the device's mutex; same for scanner mutations. This eliminates the documented race (`go/cmd/apc-mini-bridge/main.go:246` reads `target.apc` while `:425/:440` mutates it).

### Resync handshake

When a device's MIDI OUT opens (or re-opens after hot-plug), the bridge currently redraws only its boot logo (`main.go:426`). v2: after redrawing the logo, the bridge sends `/apc/v2/hello <dev>`. The plugin responds by repainting all LEDs for that device.

---

## Phase 0 — Foundation (do first; everything else depends on this)

**Why first:** every later phase touches the same multi-device wiring, OSC schema, and concurrency. Doing them once at the bottom avoids three mid-flight refactors.

### 0.1 Multi-device config

**Files:** `config.yaml`, `go/cmd/apc-mini-bridge/main.go:102-143` (Config struct + `loadConfig`).

Add an optional `devices` array. Backward-compat: if absent, synthesize from existing `midi_in/midi_out/midi_in2/midi_out2` keys.

```yaml
devices:
  - id: 1
    role: color
    midi_in: "APC mini mk2 1"
    midi_out: "APC mini mk2 3"
    midi_in_index: -1
    midi_out_index: -1
  - id: 2
    role: wing
    midi_in: "APC mini mk2 2"
    midi_out: "APC mini mk2 4"
    wing_offset: 600       # exec block base for this wing — faders 601..608
osc:
  local_ip: "127.0.0.1"
  local_port: 8001
  remote_ip: "127.0.0.1"
  remote_port: 8000
```

The `role` field is informational to the bridge (logged for diagnostics); routing decisions live in the plugin.

### 0.2 Replace dev1/dev2 with `map[int32]*apcDevice`

**Files:** `go/cmd/apc-mini-bridge/main.go:218-236, 241-292, 305-450, 457-461`.

Replace literal `dev1`, `dev2` with `devices map[int32]*apcDevice`. Loop on init to construct entries from the config array, loop on shutdown to clean up. OSC handlers look up by `dev` id and **reject + log** unknown ids (codex: "do not silently route to dev1"). Per-device scanner goroutine spun up in a loop.

`apcDevice` struct gains:
- `role string`
- `wingOffset int` (only meaningful for `role == "wing"`)
- `mu sync.Mutex` (Phase 0.4)

Logo arrays become a slice indexed by device id, with a default fallback for unconfigured ids.

### 0.3 OSC schema v2 in the bridge

**Files:** `go/internal/osc/server.go` (add new handlers), `go/cmd/apc-mini-bridge/main.go` (emit v2 routes outbound).

Add inbound v2 handlers: `/apc/v2/led/grid`, `/apc/v2/led/button`, `/apc/v2/shutdown`. Keep v1 handlers untouched. Each v2 handler validates argument count and types strictly (don't replicate v1's permissive parsing).

Outbound: when a MIDI event arrives from the controller, emit the v2 form: `/apc/v2/input/grid`, `/apc/v2/input/button`, `/apc/v2/input/fader`. Keep the v1 outbound (`/Pad`, `/Page`, `/Trigger`) gated behind a `legacy_osc: true` per-device flag for one release cycle so we don't break a running `apc_color.lua` mid-show.

Add `/apc/v2/hello <dev>` send on output port open.

### 0.4 Per-device mutex around `d.apc`

**Files:** `go/cmd/apc-mini-bridge/main.go:241-294, 425-446`.

Wrap every read/write of `d.apc` with `d.mu`. The scanner sets/clears `d.apc` under the lock; OSC handlers read+call under the lock. MIDI output is naturally serialized as a side effect (correct: `LightGrid` and `LightPage` could race today).

### 0.5 Channel constant rename + LSB color typos

**Files:** `go/internal/midi/apc_mk2.go:7-12`, `go/internal/color/map.go` (velocities 67, 70, 72).

Rename `ChannelSolid` → keep, `ChannelSlowBlink` → `ChannelPulse16`, `ChannelFastBlink` → `ChannelPulse8`. Add `ChannelBlink16 = 11` for true blink. Update call sites in `main.go:250-258`. (Documented misnaming per `docs/midi_and_apc_mini_mk2/10_code_vs_spec.md`.)

Fix LSB typos: `#0000FE` → `#0000FF`, `#7F7F70` → `#7F7F7F`, `#FF0001` → `#FF0000`.

### 0.6 Pad-orientation question

Before plumbing the fire-all button, run the one-line console test in `docs/midi_and_apc_mini_mk2/03_note_assignments.md` to resolve whether note 63 is top-right (AKAI spec) or bottom-right (existing comments). Pin the answer in code comments.

---

## Phase 1 — Color AKAI (apc_color2.lua)

**Why second:** the user's primary v2 feature. Builds on Phase 0's schema + multi-device.

### 1.1 Port the OSC codec from apc_color.lua → apc_color2.lua

**Files:** `ma3/lua/apc_color2.lua`, reference `ma3/lua/apc_color.lua:67-102, 245-296`.

Lift verbatim and adapt to v2 schema:
- `osc_pad4`, `be32`, `osc_msg` (lines 67-102) — generic OSC encoder.
- `parse_osc_packet` (lines 245-284) — **rewrite to be type-tag-driven, not position-based**. Codex flagged this as a maintainability bug; it'll bite us when we add new arg types.
- `poll_once` (lines 286-331) — drain UDP. Codex "nice-to-have": loop until socket would block, not just one datagram per poll.

Address constants update from `/Pad`/`/Page` → `/apc/v2/input/grid`/`/apc/v2/led/grid` etc.

### 1.2 Sequence + cue discovery (configurable name filters)

**Files:** `ma3/lua/apc_color2.lua` (rewrite `Setup`, line 40-90 is the existing skeleton).

Plugin args (passed via `Plugin "apc_color2" "init seq_prefix=APCColor_ cue_prefix=APCColor_ devices=1,2"`):

- `seq_prefix` — sequence name filter (default `APCColor_`).
- `cue_prefix` — cue name filter (default `APCColor_`). Cues without this prefix are skipped, **even within a matching sequence**.
- `devices` — comma-separated device ids this plugin instance handles.

Layout (per device, generalizing the existing `state.dev1` skeleton to `state.devices[id]`):

- One sequence per *starting* row.
- Within a sequence, cues are placed left-to-right; if more than 8 mappable cues exist, wrap onto the next row (still that sequence). The sequence may span multiple rows.
- Next sequence starts on a fresh row.
- 8 rows = one page. Right side buttons (notes 112-119) page through.
- Per-device `loaded[seq_id] = cue_no` table (Codex "must": set, not single).

Skip the `OffCue`/`CueZero` synthetic children (already handled by `Count_display_cues`). Remove the debug `cue:Dump()` at `apc_color2.lua:134` (will spam logs in production).

### 1.3 Auto-color from cue Appearance

**Files:** `ma3/lua/apc_color2.lua` (new `read_cue_color` to replace the legacy stub at `apc_color.lua:119-124`).

Read `cue.Appearance.BackR/G/B` and convert to `#RRGGBB`. Per `docs/grandMA3_practical_guide/03_sequence_cue_executor.md` §3.3, the float range is unverified — handle both 0..255 and 0..1 by sniffing max value and scaling. Cite that this needs console verification.

If `cue.Appearance` is nil, fall back to the *sequence*'s appearance, then to a configured default. Skip `nil`-deref crashes.

### 1.4 LED render

**Files:** `ma3/lua/apc_color2.lua` (port `render_leds_for_page` from `apc_color.lua:154-175`, restructure for per-device + bitmask state).

For each pad on each device's current page:

```lua
local cue, seq = pad_to_cue_and_sequence(dev, page, pad)
if not cue then send_led(dev, pad, 0, "#000000"); return end
local color = read_cue_color(cue)
local active = (read_active_cue(seq) == cue.idx)
local loaded = (state.devices[dev].loaded[seq.id] == cue.no)
local st = (loaded and 1 or 0) | (active and 2 or 0)   -- bitmask
send_led(dev, pad, st, color)
```

Codex "nice-to-have": diff against last sent state, emit only changed pads. 64 pads × 2 devices × 200ms = 640 packets/s in the worst case; diffing typically drops this by ~100×.

### 1.5 Active-cue read

**Files:** `ma3/lua/apc_color2.lua` (new `read_active_cue`).

`seq:CurrentChild()` (verified in `docs/grandMA3_lua_functions.txt:214` and `docs/grandMA3_practical_guide/03_sequence_cue_executor.md:61-67`). Cache per polling cycle to avoid 64 reads per pad-render. Return the child handle's index in the sequence's `Children()` array, mapping to a cue display number via `child.No / 1000`.

**Cue number formatting note:** cue numbers are **decimal** (stored as `display × 1000`, so cue "1.5" is `No=1500`). When formatting for `Cmd()`, use `%g` after dividing — `string.format("Load Sequence %d Cue %g", seq_id, cue.No / 1000)` — not `%d`, which truncates fractional cues.

**Behavior when sequence is Off** is documented as unverified (`docs/grandMA3_practical_guide/03_sequence_cue_executor.md:66-67`). Pre-flight test: read `seq:CurrentChild()` while sequence is off — expect either `nil` or the OffCue handle. Code defensively for both.

### 1.6 Pad press → load

**Files:** `ma3/lua/apc_color2.lua` (new `handle_grid_input` replacing `handle_pad`).

```lua
local cue, seq = pad_to_cue_and_sequence(dev, page, pad)
if not cue then return end
local cue_no = cue.No / 1000   -- display number; may be fractional (e.g. 1.5)
if state.devices[dev].loaded[seq.id] == cue_no then
  state.devices[dev].loaded[seq.id] = nil   -- toggle off (unload)
else
  state.devices[dev].loaded[seq.id] = cue_no
  -- %g preserves fractional cue numbers (1.5 stays 1.5; 1.0 prints as 1)
  Cmd(string.format("Load Sequence %d Cue %g", seq.id, cue_no))
end
render_leds_for_page(dev, state.devices[dev].page)
```

`Cmd()` returns "OK" / "Syntax Error" / "Illegal Command" (`docs/grandMA3_objectfree_api/command_execution.md:52-69`). For load failures, log the result; don't silently swallow.

The Load command is issued on press so MA3's executor visualization shows the loaded cue immediately. Toggle-off (deselect) does *not* issue an Unload — there is no Lua-API "unload" command; the loaded cue is naturally superseded on the next Load or Go+. This is a design choice worth surfacing to the user during testing.

### 1.7 Fire-all-loaded button

**Files:** `ma3/lua/apc_color2.lua` (new `handle_fire_all`).

Default button: configurable in `init` args (`fire_all_pad=63` for top-right grid pad after orientation resolves). Per-device: each color device has its own fire-all button; firing only affects that device's loaded set, *or* global firing across all color devices if `fire_all_global=true`.

```lua
local function fire_all_loaded(dev)
  local loaded = state.devices[dev].loaded
  for seq_id, _cue_no in pairs(loaded) do
    -- One CmdIndirect call per sequence — fire-and-forget, non-blocking.
    -- CmdIndirect is documented as taking a single command string
    -- (docs/grandMA3_objectfree_api/command_execution.md:90-106);
    -- semicolon-batched multi-command strings are NOT documented.
    CmdIndirect(string.format("Go+ Sequence %d", seq_id))
  end
  state.devices[dev].loaded = {}
  render_leds_for_page(dev, state.devices[dev].page)
end
```

Use `CmdIndirect` (one call per sequence — async, fire-and-forget) per `docs/grandMA3_practical_guide/08_gotchas.md:60` so a 5-cue fire doesn't stall the UI. Each loaded entry fires `Go+ Sequence X`, which advances the sequence to the previously-Loaded cue (matches user spec exactly).

### 1.8 Hello-handshake response

**Files:** `ma3/lua/apc_color2.lua`.

When `/apc/v2/hello <dev>` arrives, repaint that device's full grid + page LEDs. Lets the bridge restart freely without requiring plugin restart.

### 1.9 Color-AKAI faders (open design — pick one before Phase 1 ships)

The 9 faders on the color AKAI are unassigned in the current spec. Two viable mappings, each with merit:

**Option A — Per-row sequence master.** Each of the 8 lower faders drives the master/intensity of the sequence currently shown on that row of the grid. Master fader (CC 56) drives `Master 2.1` (matches Node.js origin).
- Pros: spatially intuitive — fader 3 controls the sequence whose cues are on row 3.
- Cons: when sequences span multiple rows (cue count > 8), one fader controls the *first* row's sequence; the wrap rows have no fader. Fader assignments shift as the user pages.
- Implementation: use the **object API** `seq:SetFader({value = pct})` where `pct = math.floor(midi_cc / 127 * 100)` (MA3 fader values are 0..100, not 0..127). Signature documented at `docs/grandMA3_object_api/faders.md:99-135`. Avoid the cmdline form `Cmd("Sequence X At …")` — that is **not a valid MA3 command** (no doc support; would silently fail).
- Alternative if sequences are bound to executors: `SendOSC 1 "/Page1/Fader{exec},i,{pct}"` (still 0..100). Both reach the same target but the object API is more direct.

**Option B — Separate wing block.** Faders map to a configurable executor block (e.g. `wing_offset = 700` → faders 701-708, master = `Master 2.1`). User programs that block freely inside MA3 — speed, intensity, effects, anything.
- Pros: maximally flexible. Doesn't compete with the row-wise grid layout.
- Cons: less "automatic" — needs setup inside MA3.
- Implementation: identical to Phase 2's wing role, scoped to faders only on this device.

**Option C — Both, configurable.** Plugin args choose: `color_fader_mode = sequence_master | wing` with `wing_offset` populated for the latter. Defaults to `sequence_master` so the out-of-the-box experience is wired.

**Recommendation:** Option C with `sequence_master` default. Cheap to implement on top of Option A, defers the user's "I'm not sure yet" without blocking.

### 1.10 Periodic rescan

Either polling at 200-500ms (matches `MA3_OSC_FEEDBACK` design per `docs/grandMA3_practical_guide/01_plugin_runtime.md:80-88`) or `HookObjectChange` per sequence (caveat: spurious wake-ups, `docs/grandMA3_practical_guide/06_hooks_and_changes.md:36`). Recommendation: **polling for v2** — simpler, proven by upstream community plugins. Switch to hooks later if perf is a problem.

---

## Phase 2 — Wing AKAI (executor passthrough)

**Why third:** independent of color UX; can ship after Phase 1 lands without blocking.

### 2.1 Plugin-side wing role

**Files:** `ma3/lua/apc_color2.lua` (new module `wing_role.lua` or section).

For devices with `role=wing`, the plugin acts as a translator between APC MIDI events (received from bridge as `/apc/v2/input/...`) and MA3 built-in OSC (sent via `SendOSC <slot> "/Page1/Fader{n},i,{val}"` etc.). Reverse direction: poll executor state every 200ms (or hook), send LED commands back to the bridge as `/apc/v2/led/grid` / `/apc/v2/led/button`.

Mapping (default, with `wing_offset = 600`). All OSC sends use `SendOSC <slot> "<address>,<types>,<value>..."` per `docs/grandMA3_practical_guide/04_osc_in_ma3.md:54-74`. **Scale APC MIDI values 0..127 → MA3 percent 0..100** before sending.

| APC element | MIDI | MA3 OSC out (note: faders are percent 0..100, keys use `si` type tag) |
|---|---|---|
| Faders 1-8 | CC 48-55 | `SendOSC 1 "/Page1/Fader601,i,<pct>"` … `Fader608,i,<pct>` where `pct = math.floor(cc/127*100)` |
| Master fader | CC 56 | `SendOSC 1 "/cmd,s,Master 2.1 At <pct>"` (or skip if user wants Master 2.1 independent of the wing) |
| Top-row "FADER CTRL" buttons 1-8 (VOLUME/PAN/SEND/DEVICE/↑/↓/←/→) | Note 100-107 | Press: `SendOSC 1 "/Page1/Key601,si,Press,1"` … `Key608`. Release: `,si,Release,1`. (Built-in key OSC uses `si` type tag with "Press"/"Release" string per `docs/grandMA3_practical_guide/04_osc_in_ma3.md:35`.) |
| Grid pads (8 rows × 8 col) | Note 0-63 | Configurable; default = `/Page1/Key{wing_offset+row*100+col+1},si,Press\|Release,1` for an 8-row exec block |
| Scene-launch col 1-8 (CLIP STOP / SOLO / MUTE / REC ARM / SELECT / DRUM / NOTE / STOP ALL CLIPS) | Note 112-119 | Page-switch via `Cmd("Page +")`/`Cmd("Page 5")` (verified `docs/grandMA3_practical_guide/02_cmd_function.md:60`) or `/Page1/KeyXXX` executor keys, configurable |
| SHIFT (top-right) | Note 122 | Default: `SendOSC 1 "/cmd,s,Master 2.1 At 0"` on press, `SendOSC 1 "/cmd,s,Master 2.1 At <stored>"` on release (mimics Node.js BO). Configurable as a plain button if user wants modal use. |

`SendOSC` syntax detail: the address, type tags, and values are one comma-separated string with **no internal spaces** around the commas (`"/cmd,s,Master 2.1 At 0"`, not `"/cmd ,s,..."`).

Same `wing_offset`-driven mapping is used for LED feedback in reverse.

### 2.2 LED feedback for wing executors

Plugin polls each mapped executor and reads playback state via the executor's `.Object` property, then `:HasActivePlayback()` on that:

```lua
local exec = GetExecutor(exec_no)
if not exec then return false end
local obj = exec.Object
if not obj then return false end
return obj:HasActivePlayback()
```

This is the documented pattern (`docs/grandMA3_practical_guide/03_sequence_cue_executor.md:100-110`) and matches the user's `plugin.txt` (line 22: `exec.Object and exec.Object:HasActivePlayback()`). **Do not call `:HasActivePlayback()` directly on the executor handle** — it must be resolved through `.Object` first.

Cue Appearance for the LED color is read from the active cue's `cue.Appearance.BackR/G/B` (same caveat as §1.3 about float range — pre-flight test). Generates `/apc/v2/led/grid|button` packets back to the bridge.

This is essentially the user's `plugin.txt` pattern (the `executor_table` poll loop) generalized over a wing offset.

### 2.3 Operational note

The wing AKAI does **not** participate in load/fire-all. It's a transparent surface — if the user programs a Go button on a wing exec, that Go fires immediately on press, no load semantics.

---

## Phase 3 — Hardening

Tackle these once Phase 1 and 2 are running on a real console; they're reliability work, not feature work.

### 3.1 Type-tag-driven OSC parser in plugin

Replace `parse_osc_packet`'s position-based reads with a real type-tag walker. Codex + audit agent both flagged this; it's a latent bug and adding any new arg type to bridge → plugin breaks the parser silently.

### 3.2 SysEx Device Inquiry for serial-pinning

Replace the fragile "second port-name occurrence" multi-device strategy (`main.go:222-223, 364-372`) with a SysEx Device Inquiry handshake on each opened port — store the device serial and pin the (in, out) pair by serial, not by enumeration order. Survives unplug/replug and Windows port-name reshuffling.

Reference: `docs/midi_and_apc_mini_mk2/08_multi_device.md` §8.2.

### 3.3 Diff-only LED updates

Plugin keeps a `last_sent[dev][pad] = {state, color}` table; only sends `/apc/v2/led/grid` if (state, color) changed. Big win for the 200ms repaint cadence.

### 3.4 Stale-doc cleanup

- Update `README.md` (currently ~60% wrong per audit) to reflect v2 schema, multi-device, build target, plugin args.
- Delete or rewrite `ma3/macros/APCColorSetup.xml` (AI-generated, doesn't reflect actual plugin invocation).
- Update inline comments in `main.go` that misdescribe pad orientation, page button colors, channel behaviors.

---

## Verification

End-to-end testing requires a real MA3 console (or onPC) and at least one APC mini mk2.

**Pre-flight (resolve before merging Phase 0; each is flagged "unverified" in `docs/grandMA3_practical_guide/README.md` §Unverified items):**

1. **Pad orientation**: press the four corner pads with the bridge logging notes. Confirm whether note 63 = top-right or bottom-right. Pin in code.
2. **MA3 Load+Go+ semantics**: in MA3, run `Load Sequence 1 Cue 3` then `Go+ Sequence 1` from the cmdline; confirm cue 3 fires (not cue 4). Repeat with sequence already playing a different cue — confirm hot-swap.
3. **Appearance float range** (docs flag: `docs/grandMA3_practical_guide/03_sequence_cue_executor.md:48-49`): read `someCue.Appearance.BackR` from a cue with a known red color. If value is 255, scale is 0..255; if 1.0, scale is 0..1. Pin in plugin.
4. **`seq:CurrentChild()` while sequence is Off** (docs flag: `docs/grandMA3_practical_guide/03_sequence_cue_executor.md:66-67`): turn off a sequence and read `seq:CurrentChild()`. Expect either `nil` or the OffCue handle. Code defensively for both.
5. **`CmdIndirect` semicolon batching**: confirm whether `CmdIndirect("Go+ Sequence 1 ; Go+ Sequence 2")` actually fires both, or whether one-call-per-command is required. (Plan currently uses one-call-per-command, which is unambiguously safe; this test only matters if we ever want to optimize.)
6. **Cue numbering with fractional values**: create a cue numbered 1.5, verify `cue.No == 1500` and `Cmd(string.format("Load Sequence 1 Cue %g", 1.5))` actually fires it.

**Phase 0 — bridge build sanity:**

```bash
# Linux dev sanity
cd /mnt/devdrive/MTN/Events/ma3apcmini/go && go build ./cmd/apc-mini-bridge

# Windows release for the show machine
pwsh -NoLogo -NoProfile -Command "./scripts/build.ps1 -Configuration Release -Arch amd64"
```

Run the bridge with the new array config + a single device entry; confirm logo draws, OSC server starts, scanner finds the port, and `/apc/v2/hello 1` is sent on output open. Then with two devices.

**Phase 1 — color AKAI on console:**

1. Create a sequence `APCColor_PARs` with cues `APCColor_Red`, `APCColor_Blue`, `APCColor_Green` each with distinct `Appearance` colors.
2. Run plugin, verify the three pads light in the cue's appearance color.
3. Press one — confirm pulse-1/8 LED, MA3 cmdline shows `Load Sequence X Cue Y`.
4. Press the same pad again — confirm LED returns to solid (toggle off).
5. Press a different pad — confirm new pad pulses, MA3 shows new Load.
6. Add a second sequence `APCColor_Spots` on row 2; load one cue in each.
7. Press fire-all button — confirm both cues fire (sequences advance to loaded cues simultaneously).
8. Verify pad LED for an actively-firing cue shows pulse-1/16; pads loaded *and* active show true blink.
9. With sequence having 10 cues, verify wrap onto row 2; remaining sequences shift down.
10. Disconnect+reconnect APC USB cable — verify `/apc/v2/hello 1` fires and grid repaints.

**Phase 2 — wing AKAI:**

1. Configure second device with `role: wing, wing_offset: 600`.
2. In MA3, Page 1, program executors 601-608 with simple chase / dimmer cues.
3. Push fader 1 on the wing AKAI; confirm exec 601 fader value changes in MA3 (`/Page1/Fader601` arrives).
4. Press the row-1 top button; confirm exec 601 Goes.
5. Verify LED feedback: when exec 601 is active, the corresponding APC LED is lit; when off, dim.

**Regression:**

Run a show with `apc_color.lua` (legacy plugin) against the v2 bridge. Confirm v1 routes still work, no warnings. (Once confirmed stable, the legacy code path can be removed in a future release.)

---

## Critical files to modify

**Go bridge:**

- `go/cmd/apc-mini-bridge/main.go` — Config struct, device map, OSC handlers, scanner loop, hello-on-open, per-device mutex.
- `go/internal/osc/server.go` — add v2 handlers (`/apc/v2/led/grid`, `/apc/v2/led/button`, `/apc/v2/shutdown`); keep v1.
- `go/internal/midi/apc_mk2.go` — rename channel constants; add `ChannelBlink16 = 11`.
- `go/internal/color/map.go` — fix three LSB typos (velocities 67, 70, 72).
- `config.yaml` — add example `devices` array; document role + wing_offset.

**MA3 plugin:**

- `ma3/lua/apc_color2.lua` — full rewrite. Lift OSC codec from `apc_color.lua:67-102, 245-331`. New: per-device state, role dispatch, `read_cue_color` from Appearance, `read_active_cue` from `seq:CurrentChild()`, fire-all-loaded with `CmdIndirect`, hello-handshake.
- (Optional in Phase 2) `ma3/lua/apc_color2_wing.lua` — split file for wing-role logic, or kept inline.

**Docs / scripts:**

- `README.md` — rewrite (currently ~60% inaccurate).
- `ma3/macros/APCColorSetup.xml` — delete or rewrite.
- `scripts/install-startup.ps1` — accept `$TaskName` properly so multiple bridges can co-exist if ever needed (probably not, but cheap).

**Reference files (do not modify):**

- `ma3/lua/apc_color.lua` — keep as legacy, source of OSC codec.
- `ma3apcmini/ma3apcminimk2color.js` — Node.js origin, source of fader CC + BO mapping.
- `plugin.txt` — executor-poll snippet, source for wing role's polling loop.
- `docs/` — read for API references during implementation; do not modify unless a doc is wrong.

---

## Phasing summary

| Phase | What ships | Approx effort | Blocks Phase…|
|---|---|---|---|
| 0 — Foundation | Multi-device config, OSC v2 schema, per-device mutex, channel rename, color typo fix | 1.5 days | 1, 2 |
| 1 — Color AKAI | Cue layout, Appearance-color LEDs, load tracking, fire-all-loaded, hello handshake | 2-3 days | — |
| 2 — Wing AKAI | Per-device role dispatch, fader/button passthrough, LED feedback poll | 2 days | — |
| 3 — Hardening | Type-tag OSC parser, SysEx serial-pin, diff-LED, README rewrite | 1-2 days | — |

Parallel-ready: Phases 1 and 2 can be implemented concurrently after Phase 0 lands (different files, different code paths).

---

---

## API verification audit — summary

A second-pass review (Explore agent + independent Codex/gpt-5.5 against the docs tree) verified the plan's MA3 Lua usage. Result:

- **Verified against docs** (with file:line citations): `Cmd("Load Sequence X Cue Y")`, `Cmd("Go+ Sequence X")`, `seq:CurrentChild()`, `seq:Children()` (with OffCue/CueZero at indices 1-2), `cue.No / 1000`, `cue.Appearance.BackR/G/B` (property names verified; float range still unverified per docs), `DataPool().Sequences:Children()`, `HookObjectChange` 3-arg form, `SendOSC <slot> "/addr,types,values"`, `Cmd("Page +")`, `Cmd("FaderMaster Page X.Y At Z")`, `coroutine.yield(seconds)`, `Object:Get(prop)`, built-in OSC inbound addresses, 1-based OSC config slot.
- **Hallucinations corrected** in this revision:
  - `Cmd("Sequence X At <value>")` was not a valid command — replaced with `seq:SetFader({value=pct})` (object API) in §1.9.
  - `CmdIndirect(table.concat(cmds, ";"))` semicolon-batching was undocumented — replaced with a loop of separate `CmdIndirect()` calls in §1.7.
  - Built-in key OSC type tag is `,si,Press|Release,1`, not `,i,1` — fixed in §2.1.
  - MA3 OSC fader values are percent (0..100), not 0..127 — added scaling note in §1.9 and §2.1.
  - `SendOSC` formatting requires no spaces around commas in the type-tag string — fixed in §2.1.
  - `:HasActivePlayback()` must be called on `exec.Object`, not the executor handle directly — fixed in §2.2.
  - Cue numbers are decimal (display × 1000) — `string.format` switched to `%g` after dividing in §1.5/§1.6.
- **Unverified-with-doc-support** (called out in pre-flight):
  - Appearance color float range (0..255 vs 0..1).
  - `seq:CurrentChild()` return value when sequence is Off.
  - LuaSocket availability on physical console (sidestepped — plan uses MA3's built-in OSC where the plugin needs to hit MA3, and LuaSocket only for the bridge UDP link, which always runs on the same machine as MA3 onPC or via the show-console's loopback).

---

## Open items needing user decisions before/during implementation

1. **Sequence + cue name prefixes** — defaults `APCColor_` for both. Confirm or override per show. Configurable via plugin args either way.
2. **Wing offset (wing AKAI)** — default `600` (faders 601-608, etc.). User said "Wing 6"; this matches MA3's executor numbering convention. Confirm before Phase 2.
3. **Color-AKAI fader mapping** — user explicitly flagged this as undecided. See §1.9. Recommendation: implement both modes (sequence-master default + wing fallback) with a config switch so the user can flip after testing. If "speed" knobs end up wanted, the wing mode covers it.
4. **Fire-all button placement** — default top-right grid pad after orientation resolution. User can move via config.
5. **Master fader on wing AKAI** — drive `Master 2.1` (matches Node.js origin) or leave independent? Default: drive `Master 2.1`, with a config flag to disable.
6. **SHIFT button (note 122)** — Node.js origin used the same note as a blackout (BO). On mk2 it's labelled SHIFT. Default to BO behavior on the wing AKAI, generic configurable button on the color AKAI; user can repurpose either.
7. **Legacy plugin support window** — keep v1 OSC routes accepted indefinitely or set a removal target? Recommendation: indefinite for now, revisit once `apc_color2.lua` is field-proven.
