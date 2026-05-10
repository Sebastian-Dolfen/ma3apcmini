# 3. Note assignments

## 3.1 Clip-grid pads (notes 0..63)

**Layout: row-major, origin at the bottom-left.**

```
row 7 (top):     56 57 58 59 60 61 62 63
row 6:           48 49 50 51 52 53 54 55
row 5:           40 41 42 43 44 45 46 47
row 4:           32 33 34 35 36 37 38 39
row 3:           24 25 26 27 28 29 30 31
row 2:           16 17 18 19 20 21 22 23
row 1:            8  9 10 11 12 13 14 15
row 0 (bottom):   0  1  2  3  4  5  6  7
```

Source: same convention as the original APC mini (the mk2 protocol PDF says only
`0x00–0x3F`, but the LED bulk table in the protocol confirms the contiguous 0..63
range; the orientation is established by mk1 documentation that the mk2 inherits).
Cross-check: VirtualDJ / MXWendler / Box of Stops community references all agree note
0 is bottom-left.
Source: https://forum.bome.com/t/new-akai-pro-apc-mini-mk2-initial-led-mapping-summary/4752

**Bottom-right pad ("submit" in our bridge) = note 7**, NOT note 63.
Note 63 is the **top-right** pad.

**DISCREPANCY**: `go/internal/midi/apc_mk2.go` and the Lua plugin use `note := index`
where `index` ranges 0..63 with no explicit orientation comment. This is fine *if*
the OSC sender uses the same row-0-at-bottom convention. The plugin (`apc_color.lua`,
`render_leds_for_page`) iterates `pad = 0..63` and the bridge emits MIDI note number
== pad index. As long as the operator considers "pad 0" to be the bottom-left, the
visual result is correct. If users intuit "pad 0 = top-left", the rendering is
mirrored vertically.

The README claim that note 63 is the "submit" pad needs to specify *physical
location* — by AKAI's mapping it is the **top-right** pad, not bottom-right. If the
intended button is the bottom-right (next to the rightmost fader), that's note 7.

## 3.2 Side buttons / Scene Launch (notes 112..119)

```
side button physical (top to bottom):    Sc1  Sc2  Sc3  Sc4  Sc5  Sc6  Sc7  Sc8
note number (hex):                       0x70 0x71 0x72 0x73 0x74 0x75 0x76 0x77
note number (dec):                       112  113  114  115  116  117  118  119
default LED color:                       Green (single-color)
shift function (hold Shift +):           ClipStop Solo Mute RecArm Select Drum Note StopAll
```

Source: AKAI Communications Protocol v1.0, table at lines 60–67 / 568–575.

Our `LightPage(p)` lights `note = 111 + p` for `p = 1..8` — that's `0x70..0x77`,
which **is correct**, indexing top-to-bottom in physical order.

## 3.3 Track buttons / Clip Stop row (notes 100..107)

```
button (left to right):     Trk1  Trk2  Trk3  Trk4  Trk5  Trk6  Trk7  Trk8
note number (hex):          0x64  0x65  0x66  0x67  0x68  0x69  0x6A  0x6B
note number (dec):          100   101   102   103   104   105   106   107
default LED color:          Red (single-color)
labels under buttons:       Volume Pan Send Device  ←  →  ↑  ↓
                            (when held with Shift, these select fader-CTRL or
                             scroll the 8x8 viewport in Live's Session View)
```

Source: AKAI Communications Protocol v1.0, table at lines 51–59 / 557–564.

Our bridge does not currently consume the track-button row.

## 3.4 Shift

- Note `0x7A` (122). Channel 0. **No LED.**
- Sends Note On / Note Off to the host. Hold-modifier: pad-press combinations let
  the user choose Soft Keys, Drum Mode, Note Mode, Stop All Clips, etc.

## 3.5 Stop All Clips

- This is the **eighth Scene Launch button** (Sc8, note `0x77`) when pressed with
  Shift held — there is **no separate "Stop All Clips" note**. The mk2 firmware
  sends Note On `0x77` and Note On `0x7A` (Shift) and Live decodes the combo.

---

[Back to README](README.md)
