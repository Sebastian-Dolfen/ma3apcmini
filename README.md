# GrandMA3 APC mini mk2 Color Control

Two-part solution to control fixture color selection from an AKAI APC mini mk2 with LED feedback matching cue Appearances in a GrandMA3 Sequence.

Components
- MA3 Lua plugin: `ma3/lua/apc_color.lua`
- Go OSC↔MIDI bridge: `go/cmd/apc-mini-bridge`

OSC Contract
- Bridge → Plugin (inputs)
  - `/Pad <i:int> <value:int>` where `i=0..63` in row-major order
  - `/Page <p:int>` where `p=1..8`
  - `/Trigger 1` bottom-right pad commit
- Plugin → Bridge (LED updates)
  - `/Pad <i:int> <state:int> <color:#RRGGBB>` where `state=0 idle, 1 preloaded, 2 active`
  - `/Page <p:int>`

Paging and Mapping
- `cueIndex = (page-1)*64 + padIndex + 1`
- Out-of-range pads on last page send `state=0` and `color="#000000"`

LED Behavior
- Color velocity is computed using closest Manhattan distance to the APC palette
- Channels: 6 solid (idle), 8 fast blink (preloaded), 7 slow blink (active)

Build
```
cd go
go build ./cmd/apc-mini-bridge
```

Run
1. Plug in APC mini mk2.
2. Edit `config.yaml` if device names or ports differ.
3. Start the bridge:
```
go\apc-mini-bridge.exe -config ..\config.yaml
```
4. In MA3, load `ma3/lua/apc_color.lua` and run:
```
Plugin "apc_color" "APCColor.Setup seq=5 page=1"
```
5. Call `APCColor.Poll` repeatedly (timer/macro) to process incoming OSC from the bridge.
6. Use side buttons for pages, grid to preload, bottom-right to commit.

Notes
- The Lua file includes placeholders for MA3 API reads (active cue, cue appearance). Replace the marked sections with actual calls to the MA3 object model in your environment. LED updates, paging, and commit flow are implemented.
- Color map is a compact subset; extend `go/internal/color/map.go` using the full table from the provided JS for perfect matching.



RUN BUILD SCRIPT:

```bash
pwsh -NoLogo -NoProfile -Command "./scripts/build.ps1 -Configuration Release -Arch amd64"
```