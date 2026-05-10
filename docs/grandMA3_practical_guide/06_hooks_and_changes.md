# 6. Hooks & change notification — `HookObjectChange`

> Targets MA3 2.3.2.0

```
HookObjectChange(callback, monitored_handle, plugin_handle [, extra_handle]) -> hook_id
```

- `callback(monitored_handle, ?, extra_handle)` — fires whenever the monitored object changes. The exact arity of the callback is **partially documented**; the handles are passed positionally and the second slot is reserved.
- `plugin_handle` — pass the running plugin's own handle so the hook is automatically destroyed if the plugin is unloaded. Required.
- Returns an integer `hook_id`.
- Threading: hooks fire on the same Lua thread as everything else — they're synchronous against UI updates, not free-running.

Cleanup:

```
Unhook(hook_id)
UnhookMultiple({hook_id1, hook_id2, ...})
DumpAllHooks()
```

- Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_hookobjectchange.html>

## 6.1 Use case for our project

Instead of polling `HasActivePlayback()` every 500 ms (the `MA3_OSC_FEEDBACK` approach), we *could* hook the sequence's playback status:

```lua
local function on_seq_change(h)
  -- read state, update LED via OSC
end

local hook = HookObjectChange(on_seq_change, my_seq_handle, my_plugin_handle)
```

**Caveat:** `HookObjectChange` fires on a wide range of property changes, not just playback. You'll get spurious wake-ups (label edits, appearance edits). Filter inside the callback.

Polling at 200–500 ms remains the simplest correct approach for an APC mini-style bridge — `MA3_OSC_FEEDBACK` chose this for good reason.

---

[Back to README](README.md)
