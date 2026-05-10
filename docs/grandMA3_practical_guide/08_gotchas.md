# 8. Common gotchas

> Targets MA3 2.3.2.0

## 8.1 `Echo` vs `Printf` vs `ErrPrintf`

| Function | Goes to System Monitor | Goes to Cmd Line History |
|---|---|---|
| `Echo(s)` | yes | no |
| `Printf(fmt, ...)` | yes | yes |
| `ErrEcho(s)` | yes (red) | no |
| `ErrPrintf(fmt, ...)` | yes (red) | yes |

- All four take C-style format strings (`%s`, `%d`, `%f`). **Mismatched format specifiers crash the Lua engine** — a famous footgun.
- For dev, attach the external terminal app (`app_terminal.exe`) over CMDLINE/SYSMON so output survives UI freezes.
- Source: <https://forum.malighting.com/forum/thread/68211-grandma3-vscode-workflow/>

## 8.2 `string.starts`, `string.split`, `Filter_inplace`, `CleanNils`

These are **not** part of MA3 or standard Lua. They're community/per-author conventions:

- `string.starts(s, prefix)` — common pattern, monkey-patched by individual authors. If `apc_color2.lua` uses it, look for the `function string.starts(...)` definition near the top.
- `string.split` — likewise not standard. `string.gmatch` is the standard substitute.
- `Filter_inplace`, `CleanNils` — author-specific helpers, not idiomatic to MA3.

If you're porting code from another plugin, expect to drag these helpers along (or replace them).

## 8.3 Path handling

- `GetPath("plugins")` returns the user plugins folder. `GetPath("plugins", true)` creates it if missing.
- Windows: `C:\ProgramData\MALightingTechnology\gma3_library\datapools\plugins\`
- macOS: `/Users/<user>/MALightingTechnology/gma3_<version>/datapools/plugins/`
- Linux/console: under `/var/MALightingTechnology/...` — exact path **unverified, needs testing on console**.
- Forward slashes work everywhere; the engine normalises.
- The originating drive of an imported plugin/file is **not preserved**. If you load a file by name, search all known drives or use `FileExists`.
- Source: <https://forum.malighting.com/forum/thread/9033-how-to-get-correct-plugin-path/>
- Source: <https://forum.malighting.com/forum/thread/68211-grandma3-vscode-workflow/>

## 8.4 Locale and decimal separators

MA3 uses `.` as the decimal separator regardless of OS locale. `string.format("%f", 0.5)` is safe. `tostring(0.5)` is safe. Don't call `os.setlocale`.

## 8.5 `HelpLua` for the live function list

Type `HelpLua` in the command line and MA3 dumps the full available Lua function set to `gma3_library/grandMA3_lua_functions.txt`. **Always cross-check this file** against documentation when something seems off — the dump is ground truth for *your* installed version.

(Our project already has this dumped at `docs/grandMA3_lua_functions.txt`.)

## 8.6 Plugin reload during development

- Save your `.lua` file on disk.
- Run `ReloadAllPlugins` from the cmdline (or `ReloadPlugins <number>` for a single one).
- Note: this does **not** clear `_G` state your previous run left behind. To reset Lua state cleanly, restart the show or restart MA3.

## 8.7 `Cmd` is synchronous and yields nothing

Calling `Cmd("Off Sequence 5")` *blocks the calling coroutine until the command completes*. If you have 50 of them in a row, the UI feels sticky. Either:

- Insert `coroutine.yield(0)` periodically, or
- Use `CmdIndirect` for fire-and-forget batches.

---

[Back to README](README.md)
