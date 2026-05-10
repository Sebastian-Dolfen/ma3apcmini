# 1. Plugin runtime model

> Targets MA3 2.3.2.0

## 1.1 What a plugin is

A grandMA3 plugin is a Plugin Pool object whose body is one or more Lua **components** (`.lua` files). Each component has its own metadata and can be marked `InStream Yes` (saved into the showfile XML) or `InStream No` (file lives on disk under the user library and is reloaded with `ReloadAllPlugins`).

- Source: <https://help.malighting.com/grandMA3/2.3/HTML/plugins.html>

## 1.2 Invocation: the `Plugin` keyword

The full command syntax (from the `Plugin` keyword help page):

```
([Function]) Plugin ["Plugin_Name" or Plugin_Number](.["LuaComponent_Name" or LuaComponent_Number]) ("Argument_Value")
```

Examples documented officially:

```
Plugin 1 "Richard Roe"        -- runs the first component, passes the arg string
Plugin 1.2                    -- runs component 2 of plugin 1
Edit Plugin 2                 -- opens the plugin editor
Label Plugin 1 "Weakon"       -- renames it
```

The default function (when no `Function` keyword precedes) is `Go+`.

- Source: <https://help.malighting.com/grandMA3/2.3/HTML/keyword_plugin.html>

## 1.3 Multiple arguments

There is **no** native multi-argument form — everything past the plugin number/name is one quoted string. Two community-confirmed patterns:

```
Plugin 1 "helloworld,42,Yippee ki-yay"   -- caller comma-joins, plugin splits
Lua "MyPlugin.MyFun('p1', 'p2')"         -- skip the plugin keyword, call Lua directly
```

The second works because plugins typically register a global table when loaded. Andreas (MA staff) confirms the comma-string pattern as the intended approach.

- Source: <https://forum.malighting.com/forum/thread/8291-call-a-plugin-and-specify-more-than-one-argument/>

## 1.4 The three returned functions: `Main`, `Cleanup`, `Execute`

A plugin component is expected to `return` up to three functions, in this order:

```lua
local function Main(display_handle, args)   -- runs when plugin is invoked normally
  -- args is the raw string passed in quotes
end

local function Cleanup(...)                 -- runs after Main, optional
end

local function Execute(type, ...)           -- runs when invoked with an Action keyword
  -- e.g. `Go+ Plugin 1`, `On Plugin 1`, `Off Plugin 1`
end

return Main, Cleanup, Execute
```

- Only `Main` receives the `args` string.
- `Cleanup` runs *after* `Main` returns. It is the right place to release resources (close sockets, unhook callbacks).
- `Execute` is only called when a Function keyword is prefixed (e.g. `On Plugin 1`).
- Source: <https://grandma3.bambinito.net/guides/plugin-intro/>

## 1.5 Coroutines, yield, and the single-thread reality

This is the single most important fact for our project:

> **"A grandMA3 plugin is handled as a Lua coroutine. … MA3 fundamentally differs from GMA2 — there's no true threading; proper coroutine management is essential for responsive plugins."** — Andreas, MA staff

- All plugin code runs on the **same Lua thread as the UI**.
- `coroutine.yield(seconds)` halts only the calling plugin and lets the UI tick. The argument is in **seconds**, fractional values work (`coroutine.yield(0.1)`).
- Long loops without `yield` will freeze the console. Always yield inside polling loops.
- Source (non-blocking plugins thread): <https://forum.malighting.com/forum/thread/7973-non-blocking-plugins/>

Practical loop pattern (this matches `MA3_OSC_FEEDBACK`'s polling design):

```lua
while keep_running do
  poll_state_and_send_osc()
  coroutine.yield(0.5)   -- 500 ms tick
end
```

**Important caveat:** blocking syscalls — e.g. LuaSocket `server:accept()` or `client:receive()` with no timeout — will freeze the UI even inside a coroutine, because the OS-level call doesn't yield. Always `setTimeout(0)` or a tiny number on sockets, then yield.

- Source: <https://forum.malighting.com/forum/thread/7973-non-blocking-plugins/>

## 1.6 `Timer()` — scheduled callback

```
Timer(function:name, number:delaytime, number:max_count, [function:cleanup], [light_userdata:context])
```

- `delaytime` is in **seconds**.
- `max_count` is the number of times the callback fires (use a very large number for "indefinitely").
- `cleanup` runs once after `max_count` firings.
- The context handle ties the timer to an object's lifetime (when that object is destroyed, the timer dies).
- Source: <https://help2.malighting.com/Page/grandMA3/lua_Timer/en/1.6>

For our use case, polling executors every 200–500 ms, **`coroutine.yield` inside a `while` loop is more flexible than `Timer`** because you can change the interval dynamically and you keep linear control flow. `Timer` is better when you want fire-and-forget without holding a coroutine open.

## 1.7 LuaSocket availability

- grandMA3 ships **Lua 5.4.4** (since v1.9) with the standard libraries.
- `local socket = require("socket")` works on **onPC (Windows, macOS, Linux)** — confirmed by multiple plugins on GitHub including `MA3_OSC_FEEDBACK`.
- **Console availability is unverified — needs testing on the console.** Multiple forum posters caveat their LuaSocket plugins with "tested on onPC, untested on hardware". MA Lighting has not officially documented socket support.
- Source: <https://forum.malighting.com/forum/thread/4501-lua-plugin-executing/>
- Source (LuaSocket usage in the wild): <https://github.com/ArtGateOne/MA3_OSC_FEEDBACK>

For our APC mini bridge, **MA3's built-in OSC** (Menu > In & Out > OSC) is the recommended transport — it sidesteps the LuaSocket question entirely and works identically on console and onPC.

## 1.8 `require` and module path

- Default `package.path` since **MA3 v2.0** includes the user plugins folder, so `require "myhelpers"` finds `gma3_library/datapools/plugins/myhelpers.lua`.
- Pre-2.0 you had to extend it manually:
  ```lua
  local userplugins = GetPath('plugins')
  if not package.path:find(userplugins) then
    package.path = package.path .. ';' .. userplugins .. '/?.lua'
  end
  ```
- `require` caches modules — to force a reload during dev:
  ```lua
  package.loaded["mymod"] = nil
  local mymod = require "mymod"
  ```
- Source: <https://forum.malighting.com/forum/thread/7999-using-require-to-include-lua-modules/>

## 1.9 Persistent state, `_G`, `PluginVars`, `UserVars`/`GlobalVars`

- Lua globals (`_G`) **do not persist** across plugin reloads or showfile transfers. Don't rely on them for config.
- `UserVars()` and `GlobalVars()` return handles to MA3's variable scopes. Use:
  ```lua
  local uv = UserVars()
  SetVar(uv, "MyKey", "value")
  local v = GetVar(uv, "MyKey")
  DelVar(uv, "MyKey")
  ```
- These **are saved with the showfile** and survive restarts. To persist a Lua table, serialise it to a string and store with `SetVar`.
- `PluginVars()` exists in the API dump but is sparsely documented; treat as plugin-scoped variables. **Unverified — needs testing for actual scope semantics.**
- Source: <https://forum.malighting.com/forum/thread/3917-getting-user-global-vars-into-or-out-of-plugins/>
- Source: <https://help.malighting.com/grandMA3/2.0/HTML/lua_objectfree_setvar.html>

---

[Back to README](README.md)
