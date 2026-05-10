# 7. Community plugins worth referencing

> Targets MA3 2.3.2.0

All four meet the criteria: source visible, demonstrably used in 2.x, activity within ~2 years.

## 7.1 `ArtGateOne/ma3apcmini` — our upstream

The repo this project forked from. Node.js + Lua hybrid. Tracks 6 wings of executors and reflects appearance colours back to APC LEDs over OSC. 30 commits, last activity 2025.

- <https://github.com/ArtGateOne/ma3apcmini>

## 7.2 `ArtGateOne/MA3_OSC_FEEDBACK` — same author, just the Lua side

Pure Lua plugin. Polls executors 201–215 every 0.5 s, sends OSC out on a configurable config slot. **The cleanest reference for the polling pattern we want.** Co-developed with "Mr-Button" and Andreas.

- <https://github.com/ArtGateOne/MA3_OSC_FEEDBACK>

## 7.3 `MacTirney/GrandMA3-API-Documentation`

Community-maintained API reference — Object API, Object-Free API, plus tutorials. **Fills the gap between the bare help pages and how things actually compose.** Updated recently.

- <https://github.com/MacTirney/GrandMA3-API-Documentation>

## 7.4 `hossimo/GMA3Plugins`

Older but historically important — its **Appearance Builder** is a reference for programmatic Appearance creation, and its API wiki was the only public docs for years. Last release May 2020 — older but still cited by MA staff in forum answers.

- <https://github.com/hossimo/GMA3Plugins>

## 7.5 `patopesto/GrandMA3-Plugins`

ViewScaler, ScreenSwap, **APIDump** (this is the tool that produces `grandMA3_lua_functions.txt`), TimecodeExporter, PluginTemplate. The PluginTemplate is the cleanest skeleton for a multi-component plugin.

- <https://github.com/patopesto/GrandMA3-Plugins>

## 7.6 `imhofroger/GMA3_LUA`

Sequence/Group/Color-pool-driven workflow plugins. Less directly relevant to MIDI/OSC but their wiki has good Object API examples.

- <https://github.com/imhofroger/GMA3_LUA>

## 7.7 Bambinito's reference site

Not a repo but a hand-curated reference: `https://grandma3.bambinito.net/reference/v23/api/`. Tracks the v2.3 API. Good for cross-checking signatures against the official help.

- <https://grandma3.bambinito.net/>

---

[Back to README](README.md)
