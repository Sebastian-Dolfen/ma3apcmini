# 10. Summary of discrepancies between project code and the official spec

| # | Where                                  | Issue                                                                                                       | Severity |
|---|----------------------------------------|-------------------------------------------------------------------------------------------------------------|----------|
| 1 | `go/internal/midi/apc_mk2.go`          | `ChannelSlowBlink = 7` is actually a **pulse** (1/16), not a blink. `ChannelFastBlink = 8` is **pulse 1/8**. For real on/off blink use ch 11..15. | Cosmetic / naming |
| 2 | `go/internal/midi/apc_mk2.go::LightPage` | Comment says "bright green velocity 21" implying palette color 21, but side buttons are single-color and ignore the palette. Velocity 1..127 (except 2) all mean "on"; velocity 2 = blink. | Cosmetic / docs |
| 3 | `go/internal/color/map.go`             | Velocity 67 has `#0000FE` instead of `#0000FF`.                                                              | LSB typo |
| 4 | `go/internal/color/map.go`             | Velocity 70 has `#7F7F70` instead of `#7F7F7F`.                                                              | LSB typo |
| 5 | `go/internal/color/map.go`             | Velocity 72 has `#FF0001` instead of `#FF0000`.                                                              | LSB typo |
| 6 | `go/cmd/apc-mini-bridge/main.go`       | "2nd port-name occurrence" multi-device strategy is fragile on macOS/Windows. Better: SysEx Device Inquiry serial-based pin. | Reliability |
| 7 | Project README / Lua plugin            | "submit pad = note 63" — note 63 is the **top-right** pad. Bottom-right is note 7. If bottom-right was intended, the note number is wrong; if top-right is intended, the documentation should say so. | Doc / orientation |
| 8 | `ma3/lua/apc_color.lua::parse_osc_packet` | Parses ints from fixed offsets at end of packet; ignores type-tag string. Adding any new arg type to the bridge breaks the plugin reader. | Maintainability |

---

[Back to README](README.md)
