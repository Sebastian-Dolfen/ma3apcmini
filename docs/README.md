# `docs/` — APC mini mk2 ↔ grandMA3 Bridge Reference

This folder is the wayfinding root for everything a plugin/bridge developer needs to work on the
APC mini mk2 ↔ grandMA3 bridge in this repo. It collects MA3 Lua API references, a practical
plugin-development guide, and the MIDI / APC mini mk2 protocol reference. The audience is anyone
(human or AI) touching `go/` (the bridge `.exe`) or `ma3/lua/apc_color2.lua` (the WIP successor
plugin; `ma3/lua/apc_color.lua` is legacy).

> **MA3 version pin: targets grandMA3 2.3.2.0** (released 2025-11-03, the current public release).
> Help URL base: `https://help.malighting.com/grandMA3/2.3/HTML/...`. Version 2.4 was teased at
> ISE 2026 but is not released; re-check anything in the practical guide's "What's likely to
> change in 2.4" list when it ships. See [`grandMA3_practical_guide/README.md`](grandMA3_practical_guide/README.md)
> for the full version-confirmation note.

---

## Quick start: "I want to do X"

| I want to… | Go to |
|---|---|
| Call a MA3 command from Lua | [`grandMA3_practical_guide/02_cmd_function.md`](grandMA3_practical_guide/02_cmd_function.md) |
| Enumerate the cues of a sequence | [`grandMA3_practical_guide/03_sequence_cue_executor.md`](grandMA3_practical_guide/03_sequence_cue_executor.md) |
| Detect the active cue at runtime | [`grandMA3_practical_guide/03_sequence_cue_executor.md`](grandMA3_practical_guide/03_sequence_cue_executor.md) |
| Know what `:Children()`, `:Count()`, `:Ptr()` return | [`grandMA3_object_api/traversal.md`](grandMA3_object_api/traversal.md) |
| Read or write fader values from Lua | [`grandMA3_object_api/faders.md`](grandMA3_object_api/faders.md) |
| Look up an exact function signature | [`grandMA3_lua_functions.txt`](grandMA3_lua_functions.txt) (source-of-truth) |
| Find an APC LED channel number / behaviour | [`midi_and_apc_mini_mk2/05_led_protocol.md`](midi_and_apc_mini_mk2/05_led_protocol.md) |
| Map a clip-grid pad note number to its position | [`midi_and_apc_mini_mk2/03_note_assignments.md`](midi_and_apc_mini_mk2/03_note_assignments.md) |
| Hand-roll OSC bytes in Lua | [`midi_and_apc_mini_mk2/09_osc_primer.md`](midi_and_apc_mini_mk2/09_osc_primer.md) |
| Port OSC code from the legacy plugin | [`midi_and_apc_mini_mk2/09_osc_primer.md`](midi_and_apc_mini_mk2/09_osc_primer.md) + [`grandMA3_practical_guide/04_osc_in_ma3.md`](grandMA3_practical_guide/04_osc_in_ma3.md) |
| Convert handles ↔ address strings | [`grandMA3_object_api/addressing.md`](grandMA3_object_api/addressing.md), [`grandMA3_practical_guide/05_addressing_handles.md`](grandMA3_practical_guide/05_addressing_handles.md) |
| Subscribe to object-change events | [`grandMA3_practical_guide/06_hooks_and_changes.md`](grandMA3_practical_guide/06_hooks_and_changes.md) |
| Understand multi-device USB-MIDI port disambiguation | [`midi_and_apc_mini_mk2/02_port_enumeration.md`](midi_and_apc_mini_mk2/02_port_enumeration.md) + [`midi_and_apc_mini_mk2/08_multi_device.md`](midi_and_apc_mini_mk2/08_multi_device.md) |

---

## Folder map

| Path | Type | Purpose |
|---|---|---|
| [`grandMA3_lua_overview.md`](grandMA3_lua_overview.md) | narrative (single file) | Light orientation for the Lua API surface and the landing-page split into Object-Free vs Object API. |
| [`grandMA3_lua_functions.txt`](grandMA3_lua_functions.txt) | reference (source-of-truth) | Raw HelpLua dump. **Authoritative for signatures.** If a Markdown reference disagrees with this file, prefer the .txt. |
| [`grandMA3_objectfree_api/`](grandMA3_objectfree_api/README.md) | reference | Object-Free Lua API (189 functions, 17 categories). Look up specific free functions like `Cmd`, `GetExecutor`, `HookObjectChange`, `SetLED`. Functions used by the APC plugin are tagged **[APC]**. Start at the [folder README](grandMA3_objectfree_api/README.md). |
| [`grandMA3_object_api/`](grandMA3_object_api/README.md) | reference | Object API — methods callable on handles (sequences, cues, executors, pool entries). 20 documented methods across 8 categories. Plugin-relevant methods are marked ★. Start at the [folder README](grandMA3_object_api/README.md). |
| [`grandMA3_practical_guide/`](grandMA3_practical_guide/README.md) | narrative (read top-to-bottom) | The plugin developer's how-to: runtime model, `Cmd()` deep-dive, sequence/cue/executor model, OSC inside MA3, addressing, hooks, community plugins, gotchas. Targets MA3 2.3.2.0. Start at the [folder README](grandMA3_practical_guide/README.md). |
| [`midi_and_apc_mini_mk2/`](midi_and_apc_mini_mk2/README.md) | reference + narrative | MIDI byte-level reference for the APC mini mk2 hardware: hardware layout, port enumeration, note/CC assignments, the LED protocol, SysEx, the OSC primer, and a consolidated code-vs-spec discrepancy log. Start at the [folder README](midi_and_apc_mini_mk2/README.md). |

The per-folder READMEs are the **second-level entry points**. Always start there, not at the
numbered topic files — they carry the alphabetical / categorical indexes you need to find
anything inside.

---

## Known code-vs-spec mismatches

The MIDI agent surfaced these while crawling the AKAI protocol PDF against the bridge code.
**Do not miss these.** Full table, byte-level evidence, and the remaining four lower-priority
items are in [`midi_and_apc_mini_mk2/10_code_vs_spec.md`](midi_and_apc_mini_mk2/10_code_vs_spec.md);
each is also surfaced inline (search `DISCREPANCY`) in the relevant topic file.

| # | Affected file | Issue (one line) | Severity |
|---|---|---|---|
| 1 | [`go/internal/color/map.go`](../go/internal/color/map.go) | Velocity 67 LSB typo: `#0000FE` should be `#0000FF`. | LSB typo (bug) |
| 2 | [`go/internal/color/map.go`](../go/internal/color/map.go) | Velocity 70 LSB typo: `#7F7F70` should be `#7F7F7F`. | LSB typo (bug) |
| 3 | [`go/internal/color/map.go`](../go/internal/color/map.go) | Velocity 72 LSB typo: `#FF0001` should be `#FF0000`. | LSB typo (bug) |
| 4 | [`go/internal/midi/apc_mk2.go`](../go/internal/midi/apc_mk2.go) | `ChannelSlowBlink = 7` / `ChannelFastBlink = 8` are **pulse**, not blink. Real blink lives on channels 11..15. | Naming (misleading) |
| 5 | [`go/internal/midi/apc_mk2.go::LightPage`](../go/internal/midi/apc_mk2.go) | "bright green velocity 21" comment misleads — side buttons are single-colour and ignore the palette. | Doc / comment |
| 6 | Project README / [`ma3/lua/apc_color.lua`](../ma3/lua/apc_color.lua) | "submit pad = note 63" but note 63 is the **top-right** pad; bottom-right is note 7. Either the note number or the doc is wrong. | Pad-orientation question (needs verification) |
| 7 | [`go/cmd/apc-mini-bridge/main.go`](../go/cmd/apc-mini-bridge/main.go) | "2nd port-name occurrence" multi-device strategy is fragile on macOS/Windows. Better: SysEx Device Inquiry serial pin. | Reliability |
| 8 | [`ma3/lua/apc_color.lua::parse_osc_packet`](../ma3/lua/apc_color.lua) | Parses ints from fixed offsets at end of packet; ignores OSC type-tag string. Adding any new arg type breaks the plugin reader. | Maintainability (potentially bug) |

---

## Unverified items

The practical guide tracks an 8-item checklist of "unverified — needs testing on a real console"
flags. These are claims sourced from forum posts or older help-page versions that have not yet
been verified against MA3 2.3.2.0 on hardware. The list (with deep-links into the topic files
where the original wording lives) is in
[`grandMA3_practical_guide/README.md` § Unverified items](grandMA3_practical_guide/README.md#unverified-items).
Treat that checklist as the queue for future console testing.

---

## Conventions

- **`★`** in [`grandMA3_object_api/`](grandMA3_object_api/README.md): method is used directly by the
  APC plugin in this repo (see each function's "Used by APC plugin for:" line).
- **`[APC]`** in [`grandMA3_objectfree_api/`](grandMA3_objectfree_api/README.md): function has a
  `Used by APC plugin for:` note in its entry.
- **`DISCREPANCY`** in [`midi_and_apc_mini_mk2/`](midi_and_apc_mini_mk2/README.md): inline marker
  next to a place where this repo's code disagrees with the official spec.
- **`[ ]` / `[x]`**: unverified-item checklist in the practical guide.
- **Source-of-truth ordering for Lua signatures:**
  [`grandMA3_lua_functions.txt`](grandMA3_lua_functions.txt) > the Markdown references. If they
  disagree, prefer the .txt.
- **Always start at the folder README, not at the topic files.** The folder READMEs hold the
  per-category and alphabetical indexes; the numbered topic files are designed to be linked
  into, not browsed flat.
- **Crawl-time snapshot.** All four reference subfolders were crawled at a moment in time
  against MA3 2.3 / 2.3.2.0 docs and the AKAI v1.0 protocol PDF. If MA3 ships 2.4 (or AKAI ships
  a v1.1 PDF), expect drift; the practical guide's "What's likely to change in 2.4" list is the
  starting point for re-checks.

---

## What's NOT in here yet

This index covers what currently lives in `docs/`. There is no per-feature how-to for the
multi-APC support that's being built toward, no `apc_color2.lua` walkthrough (the WIP plugin
is still being rebuilt), and no integration tests captured against a live console — the
unverified checklist above is the placeholder for that work. The "load/execute sequences from
the controller" goal is also not yet documented as a feature; the building blocks for it live
across [`grandMA3_practical_guide/03_sequence_cue_executor.md`](grandMA3_practical_guide/03_sequence_cue_executor.md)
and [`grandMA3_object_api/traversal.md`](grandMA3_object_api/traversal.md).
