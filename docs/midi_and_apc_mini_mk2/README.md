# MIDI Protocol & AKAI APC mini mk2 Reference

This folder is the project's authoritative reference for the MIDI bytes that flow
between the AKAI APC mini mk2 hardware and our Go bridge / MA3 Lua plugin. Where the
official AKAI document and our code diverge, that's flagged with **DISCREPANCY** in
the topic file and consolidated in the [Known mismatches](#known-mismatches-between-our-code-and-the-official-spec)
section below.

Primary source is the official AKAI Communications Protocol PDF v1.0; everything else
verifies against it.

## Provenance

The reference was assembled from:

- **AKAI Pro APC mini mk2 — Communications Protocol v1.0** (PDF, primary source for
  byte-level message layouts, palette tables, SysEx envelopes):
  https://cdn.inmusicbrands.com/akai/attachments/APC%20mini%20mk2%20-%20Communication%20Protocol%20-%20v1.0.pdf
- **AKAI Pro APC mini mk2 — User Guide v1.7** (PDF, primary source for port naming
  and Live integration step-by-step):
  https://cdn.inmusicbrands.com/akai/apc-mini-mkii/APC%20mini%20mk2%20-%20User%20Guide%20-%20v1.7.pdf

Cross-checked against community references (Bome forum LED-mapping summary, Ardour
Discourse Linux thread, MADRIX MIDI map, AKAI APC40 mk2 sister-protocol document)
and against our own code in `go/internal/midi/apc_mk2.go`,
`go/internal/color/map.go`, `go/cmd/apc-mini-bridge/main.go`, and
`ma3/lua/apc_color.lua`.

## Sources cited

- AKAI Pro APC mini mk2 — Communications Protocol v1.0 (PDF):
  https://cdn.inmusicbrands.com/akai/attachments/APC%20mini%20mk2%20-%20Communication%20Protocol%20-%20v1.0.pdf
- AKAI Pro APC mini mk2 — User Guide v1.7 (PDF):
  https://cdn.inmusicbrands.com/akai/apc-mini-mkii/APC%20mini%20mk2%20-%20User%20Guide%20-%20v1.7.pdf
- AKAI Pro APC mini mk2 FAQ:
  https://support.akaipro.com/en/support/solutions/articles/69000826280-akai-pro-apc-mini-mk2-frequently-asked-questions
- AKAI Pro APC mini mk2 Troubleshooting Guide:
  https://support.akaipro.com/en/support/solutions/articles/69000826279-akai-pro-apc-mini-mk2-troubleshooting-guide
- Bome Forum — "New Akai Pro APC MINI MK2 initial LED mapping summary":
  https://forum.bome.com/t/new-akai-pro-apc-mini-mk2-initial-led-mapping-summary/4752
- Ardour Discourse — "AKAI APC mini mk2 (Linux)":
  https://discourse.ardour.org/t/akai-apc-mini-mk2/109765
- MADRIX — APC mini mk2 MIDI Map:
  https://help.madrix.com/m5/html/madrix/hidd_akai_professional_apcmini_mk2.html
- AKAI APC40 mk2 Communications Protocol v1.2 (sister-document, useful for SysEx
  envelope cross-checks):
  https://cdn.inmusicbrands.com/akai/attachments/apc40II/APC40Mk2_Communications_Protocol_v1.2.pdf
- OSC 1.0 Specification (CNMAT/UC Berkeley):
  https://opensoundcontrol.stanford.edu/spec-1_0.html
- USB Class Specification for MIDI Devices (USB-IF):
  https://www.usb.org/sites/default/files/midi10.pdf

## Topic index

| File                                             | Contents                                                                                     |
|--------------------------------------------------|----------------------------------------------------------------------------------------------|
| [01_hardware.md](01_hardware.md)                 | Top-down hardware layout, numeric ranges of all controls, differences from APC mini mk1.     |
| [02_port_enumeration.md](02_port_enumeration.md) | Multi-port USB MIDI enumeration, what Port 0 vs. Port 1 carry, OS-specific names, how the bridge picks ports. |
| [03_note_assignments.md](03_note_assignments.md) | Note numbers for clip-grid pads (0..63), Scene Launch (112..119), Track buttons (100..107), Shift, Stop All Clips. |
| [04_fader_ccs.md](04_fader_ccs.md)               | CC numbers `0x30..0x38` for 8 channel faders + master, all 7-bit, absolute position.         |
| [05_led_protocol.md](05_led_protocol.md)         | The 3-byte Note On LED message, channel-encoded behavior (brightness + pulse + blink), full 128-color velocity → palette table, single-LED button states. |
| [06_midi_primer.md](06_midi_primer.md)           | Standard MIDI byte-level reminder — status bytes, NoteOn vel=0 = NoteOff, 14-bit CCs, running status, USB-MIDI framing. |
| [07_sysex.md](07_sysex.md)                       | AKAI SysEx envelope, Device Inquiry, Introduction (mode-set), Direct 24-bit RGB lighting, the absence of a mode-set SysEx. |
| [08_multi_device.md](08_multi_device.md)         | USB enumeration, disambiguation strategies (best → worst), RtMidi / portmidi quirks.         |
| [09_osc_primer.md](09_osc_primer.md)             | OSC 1.0 wire format used between the MA3 plugin and the Go bridge (address, type-tag, args, bundles). |
| [10_code_vs_spec.md](10_code_vs_spec.md)         | Consolidated table of all 8 discrepancies between our code and the official spec.            |

## Known mismatches between our code and the official spec

These are surfaced inline in the topic files (search for `DISCREPANCY`); they are
also collected here for an at-a-glance view. The full table lives in
[10_code_vs_spec.md](10_code_vs_spec.md).

| # | Where                                  | Issue                                                                                                       | Severity            | Discussed in                                      |
|---|----------------------------------------|-------------------------------------------------------------------------------------------------------------|---------------------|---------------------------------------------------|
| 1 | `go/internal/midi/apc_mk2.go`          | `ChannelSlowBlink = 7` is actually a **pulse** (1/16), not a blink. `ChannelFastBlink = 8` is **pulse 1/8**. For real on/off blink use ch 11..15. | Cosmetic / naming   | [05_led_protocol.md §5.3](05_led_protocol.md#53-cross-check-against-our-code) |
| 2 | `go/internal/midi/apc_mk2.go::LightPage` | Comment says "bright green velocity 21" implying palette color 21, but side buttons are single-color and ignore the palette. Velocity 1..127 (except 2) all mean "on"; velocity 2 = blink. | Cosmetic / docs     | [05_led_protocol.md §5.7](05_led_protocol.md#57-verification-of-lightpage-velocity-21--bright-green) |
| 3 | `go/internal/color/map.go`             | Velocity 67 has `#0000FE` instead of `#0000FF`.                                                              | LSB typo            | [05_led_protocol.md §5.5](05_led_protocol.md#55-rgb-pad-color--velocity--palette-table-full-128-entries) |
| 4 | `go/internal/color/map.go`             | Velocity 70 has `#7F7F70` instead of `#7F7F7F`.                                                              | LSB typo            | [05_led_protocol.md §5.5](05_led_protocol.md#55-rgb-pad-color--velocity--palette-table-full-128-entries) |
| 5 | `go/internal/color/map.go`             | Velocity 72 has `#FF0001` instead of `#FF0000`.                                                              | LSB typo            | [05_led_protocol.md §5.5](05_led_protocol.md#55-rgb-pad-color--velocity--palette-table-full-128-entries) |
| 6 | `go/cmd/apc-mini-bridge/main.go`       | "2nd port-name occurrence" multi-device strategy is fragile on macOS/Windows. Better: SysEx Device Inquiry serial-based pin. | Reliability         | [02_port_enumeration.md](02_port_enumeration.md#how-our-bridge-addresses-ports), [08_multi_device.md §8.2](08_multi_device.md#82-disambiguation-strategies-best--worst) |
| 7 | Project README / Lua plugin            | "submit pad = note 63" — note 63 is the **top-right** pad. Bottom-right is note 7. If bottom-right was intended, the note number is wrong; if top-right is intended, the documentation should say so. | Doc / orientation   | [03_note_assignments.md §3.1](03_note_assignments.md#31-clip-grid-pads-notes-063) |
| 8 | `ma3/lua/apc_color.lua::parse_osc_packet` | Parses ints from fixed offsets at end of packet; ignores type-tag string. Adding any new arg type to the bridge breaks the plugin reader. | Maintainability     | [09_osc_primer.md §9.5](09_osc_primer.md#95-common-pitfalls-when-hand-rolling-and-what-apc_colorluaosc_msg-does) |

## How this folder is organized

- One topic per file, numbered `NN_topic.md` so the natural alphabetical order
  matches the original reading order.
- Each topic file is self-contained: it has an H1 title, a 1–2 line lead-in (where
  the source had one), the original H2/H3 sub-sections, all code blocks verbatim,
  and a footer link back to this README.
- Discrepancies between our code and the spec are kept inline in the topic where
  the relevant bytes / palette entries are defined (so you see the spec and the
  bug side-by-side), and also gathered into [10_code_vs_spec.md](10_code_vs_spec.md)
  and the [Known mismatches](#known-mismatches-between-our-code-and-the-official-spec)
  section above for quick cross-reference.
- Citations of the form `Source: ...` are preserved verbatim from the original
  monolithic document; raw URL references (e.g. to `/tmp/apc_mk2_docs/protocol.txt`)
  are pointers into the local PDF text-extraction the original document was built
  from and are kept for traceability.
