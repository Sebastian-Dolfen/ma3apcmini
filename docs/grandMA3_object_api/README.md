# grandMA3 Lua Object API Reference

This file is a per-function reference for grandMA3's **Lua Object API** - the methods callable on object handles (sequences, cues, executors, data-pool entries, etc.).

- **Source:** crawled from <https://help.malighting.com/grandMA3/2.3/HTML/lua_object.html> and the 20 linked subpages.
- **Version:** grandMA3 2.3 (the in-page subtitle reads "Version 2.2" but the URL prefix is `/grandMA3/2.3/HTML/`).
- **Generated for:** the APCmini-mk2 <-> grandMA3 OSC bridge plugin in this repo.

## Calling convention - read this first

Every signature listed below shows `light_userdata:handle` as the first argument, e.g.:

```
Children(light_userdata:handle): {light_userdata:child_handles}
```

In real plugin code the handle is virtually always supplied implicitly via Lua's **colon notation**:

```lua
local cues = mySequence:Children()       -- colon: handle is implicit
local cues = mySequence.Children(mySequence)  -- equivalent dot form
```

So when you read a signature `Get(handle, name[, role])` think of the call site as `obj:Get(name[, role])`.

**Note:** This reference documents only the 20 functions that have a dedicated detail page on the Object API index. The full surface (UI, grid helpers, scroll, list-builder, command helpers, undo helpers, etc.) is much larger - see `../grandMA3_lua_functions.txt` for the complete signature dump. Methods listed in `docs/grandMA3_lua_functions.txt` but missing here have **no dedicated subpage on the official docs index** — that's expected, not a gap to fill.

---

## Categories

| File | Covers |
| --- | --- |
| [addressing.md](addressing.md) | Converting handles to address strings for command-line use (`Addr`, `AddrNative`, `ToAddr`). |
| [class_introspection.md](class_introspection.md) | Class-name lookups for objects and their children (`GetClass`, `GetChildClass`). |
| [faders.md](faders.md) | Reading and writing fader values - the central area for the APCmini bridge (`GetFader`, `GetFaderText`, `SetFader`). |
| [import_export.md](import_export.md) | Reading and writing objects to XML files (`Export`, `Import`). |
| [playback_state.md](playback_state.md) | Runtime playback state and the dependency / reference graph (`HasActivePlayback`, `GetReferences`, `GetDependencies`). |
| [properties.md](properties.md) | Reading property values and inspecting full object state (`Get`, `Dump`). |
| [traversal.md](traversal.md) | Walking the object tree: enumerating, counting, indexing children (`Children`, `Count`, `Ptr`). |
| [ui_editor.md](ui_editor.md) | Looking up the UI editor / settings windows for an object (`GetUIEditor`, `GetUISettings`). |

---

## Alphabetical function index

Functions marked with **★** are used directly by the APC plugin in this repo (see each function's "Used by APC plugin for:" line for context).

| Function | Page | APC |
| --- | --- | --- |
| [Addr](addressing.md#addr) | addressing.md | ★ |
| [AddrNative](addressing.md#addrnative) | addressing.md | ★ |
| [Children](traversal.md#children) | traversal.md | ★ |
| [Count](traversal.md#count) | traversal.md | ★ |
| [Dump](properties.md#dump) | properties.md | ★ |
| [Export](import_export.md#export) | import_export.md | ★ |
| [Get](properties.md#get) | properties.md | ★ |
| [GetChildClass](class_introspection.md#getchildclass) | class_introspection.md | ★ |
| [GetClass](class_introspection.md#getclass) | class_introspection.md | ★ |
| [GetDependencies](playback_state.md#getdependencies) | playback_state.md | |
| [GetFader](faders.md#getfader) | faders.md | ★ |
| [GetFaderText](faders.md#getfadertext) | faders.md | ★ |
| [GetReferences](playback_state.md#getreferences) | playback_state.md | |
| [GetUIEditor](ui_editor.md#getuieditor) | ui_editor.md | |
| [GetUISettings](ui_editor.md#getuisettings) | ui_editor.md | |
| [HasActivePlayback](playback_state.md#hasactiveplayback) | playback_state.md | ★ |
| [Import](import_export.md#import) | import_export.md | ★ |
| [Ptr](traversal.md#ptr) | traversal.md | ★ |
| [SetFader](faders.md#setfader) | faders.md | ★ |
| [ToAddr](addressing.md#toaddr) | addressing.md | ★ |

**Functions documented:** 20
