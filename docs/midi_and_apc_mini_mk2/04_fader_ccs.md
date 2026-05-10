# 4. Fader CCs

All faders are CC on Channel 0, USB Port 0:

| Fader        | Physical position | CC# (hex) | CC# (dec) | Range  |
|--------------|-------------------|-----------|-----------|--------|
| Channel 1    | leftmost          | `0x30`    | 48        | 0..127 |
| Channel 2    |                   | `0x31`    | 49        | 0..127 |
| Channel 3    |                   | `0x32`    | 50        | 0..127 |
| Channel 4    |                   | `0x33`    | 51        | 0..127 |
| Channel 5    |                   | `0x34`    | 52        | 0..127 |
| Channel 6    |                   | `0x35`    | 53        | 0..127 |
| Channel 7    |                   | `0x36`    | 54        | 0..127 |
| Channel 8    |                   | `0x37`    | 55        | 0..127 |
| Master       | rightmost         | `0x38`    | 56        | 0..127 |

- All faders are **7-bit** (single CC byte, not 14-bit MSB/LSB pairs). No High-Res
  Velocity / 14-bit CC is implemented in firmware.
- Value is the **absolute position** of the fader (not relative deltas).
- Faders are physically vertical; 0 = bottom, 127 = top.

Source: AKAI Communications Protocol v1.0, "Channel Faders / Master Fader" table
(/tmp/apc_mk2_docs/protocol.txt lines 597–606).

---

[Back to README](README.md)
