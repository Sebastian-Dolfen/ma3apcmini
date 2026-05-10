# 5. LED protocol

## 5.1 The 3-byte Note On message

```
status        data1     data2
9X            <note>    <velocity>
```

- `9X` → status nibble `0x9` = Note On, low nibble `X` (= channel 0..15) selects
  LED **behavior** (RGB pads) or is always 0 for single-color buttons.
- `<note>` → which pad/button (see §3).
- `<velocity>` → for RGB pads: palette index 0..127. For single-color buttons:
  on/off/blink (0=off, 1 or 3..127=on, 2=blink).

A Note Off (`8X` status) or a Note On with velocity 0 turns the LED off.

## 5.2 RGB pad behavior — channel mapping

From AKAI Communications Protocol v1.0, RGB LED Behavior table
(/tmp/apc_mk2_docs/protocol.txt lines 79–109):

| MIDI ch (0-based) | Status byte | Function                     |
|-------------------|-------------|------------------------------|
| 0                 | `0x90`      | On — 10% brightness (solid)  |
| 1                 | `0x91`      | On — 25% brightness          |
| 2                 | `0x92`      | On — 50% brightness          |
| 3                 | `0x93`      | On — 65% brightness          |
| 4                 | `0x94`      | On — 75% brightness          |
| 5                 | `0x95`      | On — 90% brightness          |
| **6**             | **`0x96`**  | **On — 100% brightness (solid full)** |
| 7                 | `0x97`      | Pulsing 1/16                 |
| 8                 | `0x98`      | Pulsing 1/8                  |
| 9                 | `0x99`      | Pulsing 1/4                  |
| 10                | `0x9A`      | Pulsing 1/2                  |
| 11                | `0x9B`      | Blinking 1/24                |
| 12                | `0x9C`      | Blinking 1/16                |
| 13                | `0x9D`      | Blinking 1/8                 |
| 14                | `0x9E`      | Blinking 1/4                 |
| 15                | `0x9F`      | Blinking 1/2                 |

Pulse and blink rates are synced to an **external MIDI clock**; without clock the
device falls back to a free-running internal rate.

## 5.3 Cross-check against our code

`go/internal/midi/apc_mk2.go`:

```go
const (
    ChannelSolid     = uint8(6) // steady          ← 100% brightness (correct)
    ChannelSlowBlink = uint8(7) // active          ← actually "Pulsing 1/16"
    ChannelFastBlink = uint8(8) // preloaded       ← actually "Pulsing 1/8"
)
```

**DISCREPANCY (cosmetic / naming)**:

- `ChannelSolid = 6` is correct — that *is* full-brightness solid.
- `ChannelSlowBlink = 7` is **actually a pulse**, not a blink. The official term is
  "Pulsing 1/16" — a smooth fade in/out at 1/16 notes. If the operator expects a
  hard on/off blink, channel 11..15 should be used instead.
- `ChannelFastBlink = 8` similarly is "Pulsing 1/8" — also a pulse.
- For a true on/off blink, use ch 11 (1/24 — fastest), 12 (1/16), 13 (1/8),
  14 (1/4), 15 (1/2).

Suggested rename to match firmware semantics:
`ChannelSolid = 6` (keep), `ChannelPulseFast = 7`, `ChannelPulseSlow = 10`,
`ChannelBlinkFast = 11`, `ChannelBlinkSlow = 15`. The visual difference matters when
the LED is dim/colorful — pulsing reads as "alive but stable", blinking reads as
"alarming/active".

## 5.4 Worked examples

```
# Solid red, full brightness, on pad 0 (bottom-left):
0x96 0x00 0x05
#   |    |    +-- velocity 5 = #FF0000 (red) from palette
#   |    +------- pad 0 (bottom-left)
#   +------------ Note On, channel 6 = solid 100%

# Light pad 17 with palette color 21 (#00FF00 = green), solid 100%:
0x96 0x11 0x15
#   ^^^   ^^   ^^
#   ch6   17   21

# Pulse pad 4 in palette color 5 (red) at 1/16:
0x97 0x04 0x05

# Blink pad 0 in palette color 13 (#FFFF00 yellow) at 1/8:
0x9D 0x00 0x0D

# Turn off pad 0:
0x90 0x00 0x00
# (or use Note Off 0x80 0x00 0x00; either works)
```

## 5.5 RGB pad color — velocity → palette table (full 128 entries)

Cross-checked against `go/internal/color/map.go`. Format `velocity → #RRGGBB`:

```
  0  #000000    32  #4CFFB7    64  #033900    96  #FF7F00
  1  #1E1E1E    33  #00FF99    65  #005735    97  #B9B000
  2  #7F7F7F    34  #005935    66  #00547F    98  #90FF00
  3  #FFFFFF    35  #001912    67  #0000FF    99  #835D07
  4  #FF4C4C    36  #4CC3FF    68  #00454F   100  #392b00
  5  #FF0000    37  #00A9FF    69  #2500CC   101  #144C10
  6  #590000    38  #004152    70  #7F7F7F   102  #0D5038
  7  #190000    39  #001019    71  #202020   103  #15152A
  8  #FFBD6C    40  #4C88FF    72  #FF0000   104  #16205A
  9  #FF5400    41  #0055FF    73  #BDFF2D   105  #693C1C
 10  #591D00    42  #001D59    74  #AFED06   106  #A8000A
 11  #271B00    43  #000819    75  #64FF09   107  #DE513D
 12  #FFFF4C    44  #4C4CFF    76  #108B00   108  #D86A1C
 13  #FFFF00    45  #0000FF    77  #00FF87   109  #FFE126
 14  #595900    46  #000059    78  #00A9FF   110  #9EE12F
 15  #191900    47  #000019    79  #002AFF   111  #67B50F
 16  #88FF4C    48  #874CFF    80  #3F00FF   112  #1E1E30
 17  #54FF00    49  #5400FF    81  #7A00FF   113  #DCFF6B
 18  #1D5900    50  #190064    82  #B21A7D   114  #80FFBD
 19  #142B00    51  #0F0030    83  #402100   115  #9A99FF
 20  #4CFF4C    52  #FF4CFF    84  #FF4A00   116  #8E66FF
 21  #00FF00    53  #FF00FF    85  #88E106   117  #404040
 22  #005900    54  #590059    86  #72FF15   118  #757575
 23  #001900    55  #190019    87  #00FF00   119  #E0FFFF
 24  #4CFF5E    56  #FF4C87    88  #3BFF26   120  #A00000
 25  #00FF19    57  #FF0054    89  #59FF71   121  #350000
 26  #00590D    58  #59001D    90  #38FFCC   122  #1AD000
 27  #001902    59  #220013    91  #5B8AFF   123  #074200
 28  #4CFF88    60  #FF1500    92  #3151C6   124  #B9B000
 29  #00FF55    61  #993500    93  #877FE9   125  #3F3100
 30  #00591D    62  #795100    94  #D31DFF   126  #B35F00
 31  #001F12    63  #436400    95  #FF005D   127  #4B1502
```

Source: AKAI Communications Protocol v1.0, "Velocity to RGB Color Chart"
(/tmp/apc_mk2_docs/protocol.txt lines 128–183).

### Cross-check against `go/internal/color/map.go`

Walking the 128 entries between the protocol PDF and our code:

- Velocities 0–63: **identical** (✓).
- Velocities 64–71 in our code map to:
  `#033900`, `#005735`, `#00547F`, `#0000FE`, `#00454F`, `#2500CC`, `#7F7F70`, `#202020`.
  The official PDF says 67 = `#0000FF` (not `#0000FE`) and 70 = `#7F7F7F`
  (not `#7F7F70`). **Two off-by-one cosmetic typos** in our table:
  - velocity 67: ours `#0000FE`, official `#0000FF` — last byte FE→FF.
  - velocity 70: ours `#7F7F70`, official `#7F7F7F` — last byte 70→7F.
- Velocity 72: ours `#FF0001`, official `#FF0000`. Another last-byte typo (01→00).
  (Note: 72 is a duplicate of pure red; intentional in the AKAI palette as a paired
  "darker red" partner of the 73-row warm greens.)
- Velocities 73–127: **identical** (✓).

**DISCREPANCY** — three typo'd hex values in `go/internal/color/map.go`:

| Velocity | Our value | Official  | Fix             |
|----------|-----------|-----------|-----------------|
| 67       | `#0000FE` | `#0000FF` | last byte FE→FF |
| 70       | `#7F7F70` | `#7F7F7F` | last byte 70→7F |
| 72       | `#FF0001` | `#FF0000` | last byte 01→00 |

These are visually negligible (one LSB off in one channel) but they break
exact-match lookup for code that asks for `#0000FF` and expects velocity 67. Our
`LookupVelocity` falls back to nearest-Manhattan, so a query for `#0000FF` returns
velocity 45 (also `#0000FF` — exact match), masking the typo. Worth fixing for
correctness.

## 5.6 Single-LED buttons (track / scene)

Track and Scene Launch buttons are **single-color** LEDs (red and green
respectively) with three states:

```
status      data1       data2
0x90        <0x64..0x77> <0x00 = off | 0x01 = on | 0x02 = blink | 0x03..0x7F = on>
```

- **Status is always `0x90`** (channel 0). Channel-encoded behavior does **not**
  apply. Source: protocol PDF, lines 200–214.
- Velocity 0 = off, 1 or 3..127 = on (steady), 2 = blink.
- There is **no per-button color choice**: track row is hard-wired red, scene-launch
  column is hard-wired green.

## 5.7 Verification of `LightPage` ("velocity 21 = bright green")

`apc_mk2.go::LightPage` does:

```go
return a.noteOn(uint8(111+p), 21, 0)  // ch 0, velocity 21
```

Per the protocol, scene-launch buttons (`0x70..0x77`) are **single-color green**,
and the only velocity values that matter are:

- `0x00` → off
- `0x01` or `0x03..0x7F` → on
- `0x02` → blink

So velocity **21** simply means "on". The "bright green" intuition came from the
RGB pad palette where 21 = `#00FF00`, but the side buttons **ignore** the palette —
they just look at on/off/blink. **Functionally `LightPage` works** (any value
≥ 1 except 2 yields steady-on green) but the comment is misleading.

**DISCREPANCY (cosmetic)**: comment in `LightPage` says "bright green velocity 21"
implying palette color 21, but for single-color side buttons there is no palette;
velocity 21 just means "on". Use `0x01` or `0x7F` for clarity.

To make a side button **blink**, send velocity `0x02` instead.

---

[Back to README](README.md)
