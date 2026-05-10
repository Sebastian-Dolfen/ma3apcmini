# 3. Sequence / Cue / Executor object model

> Targets MA3 2.3.2.0

## 3.1 Iterating cues in a sequence — and the OffCue / CueZero gotcha

A grandMA3 sequence handle's `Children()` returns **all children including the two synthetic cues at the start**:

| Index | Content |
|---|---|
| 1 | **OffCue** — fires when the sequence is turned Off |
| 2 | **CueZero** — implicit cue 0, the "everything off" state |
| 3..N | The user-defined cues, in order |

So a real cue iterator looks like:

```lua
local seq = ...                 -- handle to sequence
local children = seq:Children()
for i = 3, #children do         -- skip OffCue and CueZero
  local cue = children[i]
  Echo(cue.Name .. " (No=" .. cue.No .. ")")
end
```

- Source: <https://forum.malighting.com/forum/thread/8744-lua-insert-cue-in-sequence/>
- Source: <https://help.malighting.com/grandMA3/2.3/HTML/cue_sequence.html>

## 3.2 Cue numbering

- The `No` property on a cue is an integer storing **cue-number × 1000**, because cues support up to three decimal places (1.234, 5.5, etc.).
- Display value: `cue.No / 1000`.
- Source: <https://forum.malighting.com/forum/thread/8069-get-the-current-cue-number-of-selected-sequence/>

## 3.3 Cue Appearance — the property names that matter

The `Appearance` property of a cue (or any pool object) is a handle to an Appearance object. The actual color components are **separate float fields**, not a hex string:

```lua
local app = cue.Appearance
Echo("BG R="  .. app.BackR  .. " G=" .. app.BackG  .. " B=" .. app.BackB)
Echo("IMG R=" .. app.ImageR .. " G=" .. app.ImageG .. " B=" .. app.ImageB)
Echo("alpha=" .. app.ImageAlpha)
```

- `BackR`, `BackG`, `BackB`, `BackAlpha` — background (the fader/exec colour bar)
- `ImageR`, `ImageG`, `ImageB`, `ImageAlpha` — image fill (the pool tile colour)
- Values are **0..255 floats** (not 0..1) per the forum example. Unverified — confirm range on console; the help page is silent.
- Source: <https://forum.malighting.com/forum/thread/67878-reading-appearance-color-data-in-lua/>

If `cue.Appearance` is `nil`, the cue has no explicit appearance — fall back to the sequence's `Appearance`.

## 3.4 Detecting the active cue at runtime

Two equivalent paths:

```lua
-- Object-free, on the user's selected sequence
local cue = GetCurrentCue()

-- Generic, for any sequence handle
local seq = ...
local cue = seq:CurrentChild()    -- returns the currently active cue
```

- `GetCurrentCue()` returns a handle to "the last activated cue in the selected sequence". Returns `nil` if no sequence is selected. Unverified — behaviour when the sequence is Off (only OffCue conceptually active).
- `seq:CurrentChild()` is more reliable because it doesn't depend on user selection.
- Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_getcurrentcue.html>
- Source: <https://forum.malighting.com/forum/thread/7593-find-currentcue/>

## 3.5 Sequence playback state

```lua
local seq = ...                        -- a sequence handle
if seq:HasActivePlayback() then
  Echo("running")
else
  Echo("off")
end
```

- `HasActivePlayback()` returns a boolean. It is **true** while the sequence has an active cue, regardless of whether it's playing or paused.
- Distinguishing **paused vs running** from Lua is **unverified — needs testing on the console.** No documented method returns a tri-state. Likely path: check the executor's underlying object's `Paused` property if present; otherwise track state via `HookObjectChange` on changes.
- Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_object_hasactiveplayback.html>
- Source: <https://forum.malighting.com/forum/thread/8567-get-state-active-inactive-and-type-of-executor-toggle-flash-temp-via-lua-script/>

## 3.6 Executors, pages, and the `5.101` notation

- An executor is identified by a number like `201` (executor 201 on the *current* page) or `5.101` (executor 101 on page 5).
- `GetExecutor(integer) -> executorHandle, pageHandle` returns *both* handles. The integer is just the executor number — the resolved page is whatever the system maps it to (typically the current page).
- For cross-page resolution, address the **sequence** in the DataPool directly:
  ```lua
  local seq = DataPool().Sequences["Sunstrip All"]
  local fader = seq:GetFader{ token = "FaderMaster" }
  ```
  This works whether or not the sequence is currently assigned to a visible executor. Andreas's recommended pattern.
- Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_getexecutor.html>
- Source: <https://forum.malighting.com/forum/thread/9151-find-executor-location-in-lua/>

## 3.7 Reading the assigned object of an executor

```lua
local exec = GetExecutor(201) or error("no exec 201")
local obj  = exec.Object         -- the Sequence (or Group, Master, etc.) handle
if obj then
  Printf(obj:HasActivePlayback() and "ON" or "OFF")
end
```

- An executor without an assignment has `exec.Object == nil`. Always guard.
- Source: <https://forum.malighting.com/forum/thread/8567-get-state-active-inactive-and-type-of-executor-toggle-flash-temp-via-lua-script/>

## 3.8 Firing a cue programmatically

For our bridge, prefer the command line over object-method gymnastics:

```lua
Cmd(string.format("Goto Sequence %d Cue %s", seqNo, cueNoString))
-- or
Cmd(string.format("Load Sequence %d Cue %d", seqNo, cueNo))
Cmd(string.format("Go+ Sequence %d", seqNo))
```

There is an Object API method `seq:Goto{...}` but its surface differs across versions; the command line is more stable across MA3 point releases.

---

[Back to README](README.md)
