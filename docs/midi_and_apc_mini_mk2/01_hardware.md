# 1. APC mini mk2 hardware overview

Layout, top-down (from the User Guide):

```
+-----------------------------------------------+
|  [Track Buttons 1..8 row above the grid]      |   << "Clip Stop / Soft Keys" row
|                                               |
|  +---+---+---+---+---+---+---+---+   +-----+  |
|  |56 |57 |58 |59 |60 |61 |62 |63 |   |Sc 1 |  |   row 7 (top)     side btn 0x70
|  +---+---+---+---+---+---+---+---+   +-----+  |
|  |48 |49 |50 |51 |52 |53 |54 |55 |   |Sc 2 |  |   row 6           side btn 0x71
|  +---+---+---+---+---+---+---+---+   +-----+  |
|  |40 |41 |42 |43 |44 |45 |46 |47 |   |Sc 3 |  |   row 5           side btn 0x72
|  +---+---+---+---+---+---+---+---+   +-----+  |
|  |32 |33 |34 |35 |36 |37 |38 |39 |   |Sc 4 |  |   row 4           side btn 0x73
|  +---+---+---+---+---+---+---+---+   +-----+  |
|  |24 |25 |26 |27 |28 |29 |30 |31 |   |Sc 5 |  |   row 3           side btn 0x74
|  +---+---+---+---+---+---+---+---+   +-----+  |
|  |16 |17 |18 |19 |20 |21 |22 |23 |   |Sc 6 |  |   row 2           side btn 0x75 (Drum mode)
|  +---+---+---+---+---+---+---+---+   +-----+  |
|  | 8 | 9 |10 |11 |12 |13 |14 |15 |   |Sc 7 |  |   row 1           side btn 0x76 (Note mode)
|  +---+---+---+---+---+---+---+---+   +-----+  |
|  | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 |   |Sc 8 |  |   row 0 (bottom)  side btn 0x77 (Stop All Clips)
|  +---+---+---+---+---+---+---+---+   +-----+  |
|                                               |
|  [F1] [F2] [F3] [F4] [F5] [F6] [F7] [F8] [FM] |   << 8 channel + master fader
|                                               |
|  [Shift]                                      |
+-----------------------------------------------+
```

Numeric ranges:

| Element                    | Count | Note range          | LED color/type       |
|----------------------------|-------|---------------------|----------------------|
| RGB clip pads (8x8)        | 64    | `0x00`–`0x3F` (0–63)| RGB, 128-color palette |
| Track buttons (top row)    | 8     | `0x64`–`0x6B` (100–107) | Single-color RED |
| Scene Launch buttons (right side) | 8 | `0x70`–`0x77` (112–119) | Single-color GREEN |
| Shift                      | 1     | `0x7A` (122)        | No LED               |
| Channel faders             | 8     | CC `0x30`–`0x37`    | n/a                  |
| Master fader               | 1     | CC `0x38`           | n/a                  |

Source: AKAI Communications Protocol v1.0, "Pad/Button Values", "Channel Faders" tables
(/tmp/apc_mk2_docs/protocol.txt lines 50–67, 595–606).

## Differences from mk1

The original APC mini (mk1) is widely documented to have:

- **Single-color** clip-grid LEDs (red/green/yellow only, with limited brightness).
- The mk1 grid uses the **same note numbers** as the mk2 (0..63, row-major from
  bottom-left) — confirmed by the Bome forum LED-mapping summary.
  Source: https://forum.bome.com/t/new-akai-pro-apc-mini-mk2-initial-led-mapping-summary/4752
- mk1 had no Note Mode / Drum Mode and no SysEx-based RGB control.
- mk1 enumerated as a single USB MIDI port; mk2 enumerates two ports (see §2).

The mk2 keeps the same physical button layout and CC/note numbers as mk1, but adds:

- Full RGB on the 64 grid pads (128-color palette via velocity).
- 16-channel LED-behavior encoding (brightness + pulse + blink).
- A Note Mode (chromatic keyboard layout) and Drum Mode (4×4 quadrants).
- A second USB MIDI port used for Note-Mode output.
- SysEx for arbitrary 24-bit RGB color and Device Inquiry.

---

[Back to README](README.md)
