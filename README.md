# ma3apcmini

A bridge between **AKAI APC mini mk2** controllers and **grandMA3** lighting consoles, supporting multiple controllers in distinct roles.

> **Status: in development.** This is a fresh v2 build, not the original `ArtGateOne/ma3apcmini` Node.js project that the repo was forked from. The legacy Lua plugin (`ma3/lua/apc_color.lua`) and Node.js source (`ma3apcmini/`) are kept in-tree as reference only. The v2 plan and current state of work live at [`docs/v2-plan.md`](docs/v2-plan.md).

## What it does

Drives one or more APC mini mk2 controllers from grandMA3, with per-device **roles** chosen at config time:

- **Color role.** Grid pads represent cues from named MA3 sequences. Each pad's LED color is derived automatically from the cue's `Appearance`. Pressing a pad **loads** that cue (`Load Sequence X Cue Y`); a dedicated button fires every loaded cue at once (`Go+ Sequence X` per loaded sequence). Loaded pads pulse, active pads pulse differently, and pads that are both blink — so the controller surface always shows current state at a glance.
- **Wing role.** The controller behaves as a transparent fader-wing surface mapped to a configurable MA3 executor block. The user programs the wing inside MA3 like any normal fader wing — the bridge just relays MIDI events to MA3's built-in OSC executor scheme and reflects executor state back to the LEDs.

A single bridge process can drive multiple APCs in any mix of roles.

## Components

```
APC mini mk2 ─MIDI─▶ apc-mini-bridge (Go) ─OSC─▶ apc_color2 (MA3 Lua plugin) ─Cmd/SendOSC─▶ grandMA3
APC mini mk2 ◀─MIDI─ apc-mini-bridge (Go) ◀─OSC─ apc_color2 (MA3 Lua plugin)
```

| Component | Path | Purpose |
|---|---|---|
| Go bridge | `go/cmd/apc-mini-bridge` | MIDI ↔ OSC translator. Generic, role-agnostic. Built as `dist/apc-mini-bridge.exe` for Windows show machines. |
| MA3 plugin | `ma3/lua/apc_color2.lua` | Owns sequence layout, cue load tracking, fire-all-loaded UX, and wing executor passthrough. WIP — see [`docs/v2-plan.md` Phase 1](docs/v2-plan.md). |
| Build / install scripts | `scripts/` | PowerShell helpers for the Windows show machine. |

## Build

Windows release (the show-machine target):

```powershell
pwsh -NoLogo -NoProfile -Command "./scripts/build.ps1 -Configuration Release -Arch amd64"
```

Output: `dist/apc-mini-bridge.exe`.

Linux dev sanity check (requires ALSA dev libs for the `go-rtmidi` cgo build):

```bash
cd go && go build ./cmd/apc-mini-bridge
```

## Configuration

Bridge runtime config is `config.yaml`. The v2 schema uses an optional `devices` array (with backward-compat for the legacy four-key form). See [`docs/v2-plan.md` §Phase 0.1](docs/v2-plan.md) for the schema and per-device role/`wing_offset` fields.

## Run

```
dist\apc-mini-bridge.exe -config config.yaml
```

Then in grandMA3, load `ma3/lua/apc_color2.lua` and start the plugin (invocation arguments documented in [`docs/v2-plan.md` §Phase 1.2](docs/v2-plan.md)).

To launch the bridge automatically at user logon on Windows:

```powershell
./scripts/install-startup.ps1     # register a scheduled task
./scripts/uninstall-startup.ps1   # remove it
```

## Documentation

The `docs/` tree is the wayfinding root for everything a developer or operator needs.

| Path | What it covers |
|---|---|
| [`docs/README.md`](docs/README.md) | Top-level index — "I want to do X" map across all the references below. |
| [`docs/v2-plan.md`](docs/v2-plan.md) | The current implementation plan: multi-device, color/wing roles, OSC schema v2, verification checklist. |
| [`docs/hardware.md`](docs/hardware.md) | APC mini mk2 physical layout reference (with photo). |
| [`docs/grandMA3_practical_guide/`](docs/grandMA3_practical_guide/README.md) | MA3 Lua plugin development how-to: `Cmd()` semantics, sequence/cue/executor model, OSC inside MA3, hooks. |
| [`docs/grandMA3_objectfree_api/`](docs/grandMA3_objectfree_api/README.md) | MA3 free Lua functions (189 entries, plugin-relevant ones tagged). |
| [`docs/grandMA3_object_api/`](docs/grandMA3_object_api/README.md) | MA3 object methods (handles → `:Children()`, `:CurrentChild()`, `:HasActivePlayback()`, etc.). |
| [`docs/grandMA3_lua_functions.txt`](docs/grandMA3_lua_functions.txt) | Source-of-truth raw `HelpLua` dump. If a Markdown reference disagrees with this file, prefer the .txt. |
| [`docs/midi_and_apc_mini_mk2/`](docs/midi_and_apc_mini_mk2/README.md) | APC mini mk2 byte-level MIDI protocol reference. |

## License

This project is licensed under the [GNU Affero General Public License v3.0](LICENSE), with an [additional permission](NOTICE) under AGPLv3 § 7 for MTN Media Group and Nemorit UG. All other use is governed by the unmodified AGPLv3.

## Origin

Forked from [`ArtGateOne/ma3apcmini`](https://github.com/ArtGateOne/ma3apcmini) (Node.js bridge) and rebuilt as a Go bridge + a new MA3 Lua plugin.
