# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this project is

A **fork** of the original `ma3apcminimk2color` Node.js project, being rebuilt as a Go bridge + MA3 Lua plugin under the name v2. Goal: drive one or more AKAI APC mini mk2 controllers from grandMA3 with per-device **roles**:

- **Color role** — pads represent cues from named sequences; pressing a pad *loads* it; a dedicated button *fires all loaded* cues at once. LED color auto-derived from cue Appearance.
- **Wing role** — controller is a transparent fader-wing surface mapped to a configurable MA3 executor block; programmable inside MA3 like any normal wing.

The current implementation plan — multi-device support, OSC schema v2, color/wing roles, verification checklist, open decisions — is the **authoritative direction document**: [`docs/v2-plan.md`](docs/v2-plan.md). Read it before starting non-trivial work.

### Trustworthy vs. stale files

Some files in the repo are out of date or were AI-generated and got it wrong. Before editing, know which is which:

| Path | Status | Notes |
|---|---|---|
| `go/` | **Active.** This is the bridge (`apc-mini-bridge.exe`). MIDI ↔ OSC translator. Already supports two devices. | Real, current code. |
| `ma3/lua/apc_color2.lua` | **Active / WIP.** This is the *new* MA3 plugin — the one to extend. | Sequence enumeration started; OSC I/O **not yet wired**. |
| `ma3/lua/apc_color.lua` | **Legacy.** Earlier prototype of the MA3 plugin. | OSC + LED + commit flow are wired here, but reads from MA3's object model are stubbed (`read_active_cue`, `read_cue_color`, `rescan` placeholders). Useful as a reference for the OSC wiring; **do not build new features on top of it.** |
| `ma3apcmini/ma3apcminimk2color.js` | **Origin.** The original Node.js project this is forked from. | Reference for the full APC mk2 color→velocity LUT, original OSC scheme (`/Key101..208`, `/Fader201..208`, `/cmd`), and original feature behavior (faders, BO, page buttons firing MA3 commands). Don't edit. |
| `README.md` | **Active.** v2-forward project page; describes color/wing roles, components, build/run, and links the docs tree. | Authoritative. Mirrors the v2 plan's framing. |
| `LICENSE`, `NOTICE` | **Active.** AGPLv3 with a § 7 additional permission for MTN Media Group and Nemorit UG. | Don't change without legal review. |
| `docs/v2-plan.md` | **Active.** The v2 implementation plan: phases, OSC schema v2, file-by-file changes, pre-flight verification checklist, open decisions. | **Read this before non-trivial work.** All MA3 Lua API claims in it have been audited against `docs/grandMA3_*` (see the plan's "API verification audit" section). |
| `docs/hardware.md` | **Active.** APC mini mk2 physical layout reference, with `docs/akai.jpg`. | The "what does the controller look like" page; complements the byte-level protocol reference under `docs/midi_and_apc_mini_mk2/`. |
| `ma3/macros/APCColorSetup.xml` | **AI-generated junk.** | A macro that an earlier AI session produced; it does not reflect how the plugin is meant to be wired. Ignore unless you're rewriting it. |
| `config.yaml` | Active. | Bridge runtime config (MIDI port names + OSC endpoints). v2 will introduce a `devices` array — see [`docs/v2-plan.md` §Phase 0.1](docs/v2-plan.md). |
| `dist/apc-mini-bridge.exe` | Active. | Prebuilt Windows binary; this is what gets run on the show machine. |
| `shared/shows/Akai APCmini mk2 OSC demo.show` | Reference. | Demo MA3 showfile. |
| `scripts/build.ps1`, `install-startup.ps1`, `uninstall-startup.ps1` | Active. | Windows-targeted PowerShell. |
| `docs/` | **Active reference tree.** MA3 Lua API references, MA3 plugin-dev practical guide, MIDI/APC mini mk2 protocol reference, the v2 plan, the hardware page. | **Start at [`docs/README.md`](docs/README.md).** Before non-trivial changes, read at least the relevant section there. |

The **build target is Windows.** Development happens on Linux but the bridge ships as a Windows .exe.

## Build & run

```powershell
# Build the Go bridge for Windows (writes dist/apc-mini-bridge.exe)
pwsh -NoLogo -NoProfile -Command "./scripts/build.ps1 -Configuration Release -Arch amd64"
```

Plain Go build (Linux dev sanity check; cgo via go-rtmidi requires ALSA dev libs):

```bash
cd go && go build ./cmd/apc-mini-bridge
```

Run on the target machine:

```
dist\apc-mini-bridge.exe -config config.yaml
```

There are no tests in this repo.

`scripts/install-startup.ps1` registers a Windows Scheduled Task to launch the bridge at logon; `uninstall-startup.ps1` removes it.

## Architecture

### Go bridge (`go/`)

Entry point `go/cmd/apc-mini-bridge/main.go`. Internal packages:

- `internal/midi` — APC mk2 abstraction. Encodes LED behavior as the **MIDI channel** of a Note On. The code defines `ChannelSolid=6`, `ChannelSlowBlink=7`, `ChannelFastBlink=8` — but **per the AKAI spec, channels 7/8 are *pulse* modes, not blink**; true blink lives on channels 11..15. The names mislead; the runtime behavior on hardware is "pulse." See [`docs/midi_and_apc_mini_mk2/05_led_protocol.md`](docs/midi_and_apc_mini_mk2/05_led_protocol.md) for the full channel table.
- `internal/osc` — OSC server (LED commands inbound from plugin) with `/Pad`, `/Page`, `/Shutdown` handlers.
- `internal/color` — APC mk2 palette → velocity LUT (128 entries) plus a Manhattan-distance fallback for unknown colors. Source of truth for the full table is the original `ma3apcmini/ma3apcminimk2color.js`; the Go map mirrors it. **Three LSB typos against the official spec** at velocities 67 / 70 / 72 (`#0000FE` → `#0000FF`, `#7F7F70` → `#7F7F7F`, `#FF0001` → `#FF0000`) — see [`docs/midi_and_apc_mini_mk2/10_code_vs_spec.md`](docs/midi_and_apc_mini_mk2/10_code_vs_spec.md).

The bridge **always opens two APC instances** (`dev1`, `dev2`). If `midi_in2`/`midi_out2` aren't set, dev2 falls back to the **2nd port-name occurrence** of the same name (a single APC mk2 enumerates as multiple USB MIDI ports). When extending to >2 devices, this is where the work starts: the device set is hardcoded to two in `main.go`.

A scanner goroutine runs every 3 s, opening MIDI ports as the controller appears/disappears. When an output opens, it draws an **8×8 logo** (`Logo8x8Dev1` / `Logo8x8Dev2` matrices in `main.go`) which is overwritten by plugin LED traffic and redrawn on `/Shutdown`.

Logging uses ANSI-colored prefixes via `logInfo` / `logMIDI` / `logOSCIn` / `logOSCOut` — use those helpers, not `log.Printf`, to keep categories consistent. `-no-color` disables ANSI.

### OSC contract (the load-bearing interface)

The Go bridge and Lua plugin are decoupled through OSC. Both sides need to stay in sync — the Lua-side encoder/parser in `apc_color.lua` is **hand-rolled** and parses arguments by trailing-byte position, not by type tag. Don't add float args or mixed type tags without updating that parser.

> **Schema status: v1.** What follows describes the *current code* (legacy `/Pad`, `/Page`, `/Trigger`, `/Shutdown` routes). The v2 plan introduces versioned namespaced routes (`/apc/v2/input/grid`, `/apc/v2/led/grid`, etc.) with mandatory `dev` first arg, plus a `/apc/v2/hello` resync handshake — see [`docs/v2-plan.md` §"OSC contract v2"](docs/v2-plan.md). v1 routes will keep working as legacy compat for `apc_color.lua`.

**Bridge → Plugin** (controller events, defined in `main.go`'s `registerCB`):
- `/Pad <dev:int> <i:int> <state:int>` — `dev=1|2`, `i=0..63`, `state=1` press / `0` release. Pad indexing: code and old comments call this "row-major from top-left," but the AKAI spec says note 0..63 is row-major **from bottom-left** (so note 63 is **top-right**, not bottom-right). This needs verification on hardware — see [`docs/midi_and_apc_mini_mk2/03_note_assignments.md`](docs/midi_and_apc_mini_mk2/03_note_assignments.md).
- `/Page <dev:int> <p:int>` — side button → page 1..8.
- `/Trigger 1` — fires when pad note 63 is pressed. Documented as "bottom-right" in the existing code/README, but per the AKAI spec note 63 is the **top-right** pad — open discrepancy, do not assume the comment is correct.

**Plugin → Bridge** (LED + lifecycle, accepted by `internal/osc/server.go`):
- `/Pad <dev:int> <i:int> <state:int> <color:#RRGGBB>` — `state=0` idle solid / `1` preloaded (fast blink) / `2` active (slow blink). Legacy 3-arg form `(i, state, color)` without `dev` is also accepted and routes to dev1.
- `/Page <dev:int> <p:int>` — light the page side button. Legacy 1-arg form accepted.
- `/Shutdown <dev:int>` — plugin going away; bridge restores the device's logo. Legacy 0-arg form accepted.

In `apc_color.lua`, `cfg.dev` controls whether the plugin emits the legacy form (`dev <= 1`) or the new dev-prefixed form. New work in `apc_color2.lua` should always emit the dev-prefixed form.

The pad/cue mapping used by `apc_color.lua` is `cueIndex = (page-1)*64 + padIndex + 1`. The legacy plugin sends `state=0`, `color=#000000` for indices past `state.total`.

### MA3 Lua plugin (`ma3/lua/apc_color2.lua`) — where new work goes

`apc_color2.lua` is structured around enumerating **sequences** named `APCColor*` from the MA3 `DataPool().Sequences` and laying them out on the grid:

- 8 buttons per **line** (one row of the APC).
- 8 lines per **page**.
- Each sequence starts on a fresh line, taking `ceil(cue_count / 8)` lines.
- `Count_display_cues(seq)` excludes cues named `OffCue` and `CueZero`.
- `state.dev1` / `state.dev2` skeletons exist for multi-device routing — symmetric to the bridge's two-device design.

What's **not** done in `apc_color2.lua` yet:
- No UDP socket setup, no OSC encode/decode, no `Poll`. The `mainloop` just yields. To reach feature parity with `apc_color.lua` and then surpass it, the OSC plumbing has to be ported in (referencing `apc_color.lua`'s `osc_msg`/`osc_pad4`/`be32` and `parse_osc_packet`), and the LED render + commit (`Load Sequence N Cue M` then `Go+ Sequence N`) flows reattached to the new sequence-layout model.
- `read_active_cue` / appearance-color reads against the real MA3 API still need to be implemented (the legacy plugin had them as TODOs too — solving these is part of the "load and execute sequences" feature).

When adding **multi-APC** support: the Go bridge is partly there (dev1/dev2 hardcoded, see above); the plugin side needs `state.dev2` filled in symmetrically and the layout/render routines parameterized by device. The v2 plan §Phase 0.2 generalizes this to N devices via a `map[int32]*apcDevice` keyed by device id.

## References (`docs/`)

`docs/` is a wayfinding tree. **Always start at [`docs/README.md`](docs/README.md)** — it has an "I want to do X" map, the MA3 version pin (**2.3.2.0**, released 2025-11-03), a consolidated table of known code-vs-spec mismatches, and a checklist of unverified items still needing console testing.

Most-important entry points by task:

- **Current implementation direction:** [`docs/v2-plan.md`](docs/v2-plan.md). The phased breakdown of all v2 work, with file:line citations for each change and a verified MA3 Lua API audit. Always cross-check against this before deciding how to implement something new.
- **Hardware layout:** [`docs/hardware.md`](docs/hardware.md). Photo + MIDI assignment table for the APC mini mk2.
- **Plugin-dev how-to (read top-to-bottom):** [`docs/grandMA3_practical_guide/`](docs/grandMA3_practical_guide/README.md). The high-leverage subsections for current work are `02_cmd_function.md` (`Cmd` / `CmdIndirect` / `CmdIndirectWait` semantics, undo handles), `03_sequence_cue_executor.md` (cue enumeration — including the `OffCue`/`CueZero` indices 1-2 gotcha and `cue.No / 1000` cue numbering), `04_osc_in_ma3.md` (built-in OSC config slots, `SendOSC` syntax), `06_hooks_and_changes.md` (`HookObjectChange` for sync without polling), `08_gotchas.md`.
- **Lua API lookup (look up specific functions):** [`docs/grandMA3_objectfree_api/`](docs/grandMA3_objectfree_api/README.md) (189 free functions like `Cmd`, `GetExecutor`, `DataPool`, `HookObjectChange` — plugin-relevant ones tagged `[APC]`) and [`docs/grandMA3_object_api/`](docs/grandMA3_object_api/README.md) (20 methods on handles like `:Children()`, `:Get(prop)`, `:HasActivePlayback()` — plugin-relevant ones tagged `★`).
- **Signature source-of-truth:** [`docs/grandMA3_lua_functions.txt`](docs/grandMA3_lua_functions.txt). If a Markdown reference disagrees with this file, prefer the .txt — it's the raw `HelpLua` dump.
- **MIDI / APC mk2 protocol:** [`docs/midi_and_apc_mini_mk2/`](docs/midi_and_apc_mini_mk2/README.md). Critical files: `05_led_protocol.md` (full 128-entry palette + channel-encoded LED behaviors), `09_osc_primer.md` (the wire format our hand-rolled Lua encoder targets), `10_code_vs_spec.md` (every place this codebase disagrees with the official AKAI v1.0 protocol PDF, with severity).

These were captured at a moment in time against MA3 2.3 / 2.3.2.0 docs and the AKAI mk2 protocol PDF v1.0. If MA3 ships 2.4 or AKAI ships a new PDF, expect drift; the practical guide has a "what's likely to change" pointer.

## Workflow rules

- **Never commit or push without an explicit user instruction.** Edits, file writes, and local builds are fine without asking. `git commit` and `git push` happen only when the user says so in plain words ("commit and push", "ship it", etc.). After making changes, describe what was changed and wait — do not pre-commit "to be helpful." This applies even when the change is small, even when previous commits in the same session were authorized.

## Conventions specific to this repo

- All Lua-side OSC bytes are built by hand (`osc_msg`, `osc_pad4`, `be32` in `apc_color.lua`); the parser is similarly minimal and **position-based**, not type-tag-driven. Keep new messages int/string-only or upgrade the codec.
- `config.yaml` is allowed to be missing; defaults in `loadConfig()` apply. Keep that property.
- The bridge logs are color-coded by category. Use the existing `logInfo` / `logMIDI` / `logOSCIn` / `logOSCOut` helpers.
- When reading the Node.js origin (`ma3apcmini/ma3apcminimk2color.js`), note that its OSC vocabulary (`/Key101..208`, `/Fader201..208`, `/cmd`) is the **MA3 built-in executor OSC scheme** — not the new bridge↔plugin protocol. Don't conflate them.
