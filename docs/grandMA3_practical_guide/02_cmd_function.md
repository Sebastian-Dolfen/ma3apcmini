# 2. The `Cmd()` function in depth

> Targets MA3 2.3.2.0

## 2.1 Basic `Cmd`

```
Cmd(string[, handle]) -> string
```

- Executes a command-line string **synchronously**. Blocks the Lua coroutine until the command completes.
- Do **not** include `please`/`enter` — `Cmd` adds the terminator itself.
- Returns one of: `"OK"`, `"Syntax Error"`, `"Illegal Command"`.
- The optional handle is an **undo-list handle** from `CreateUndo` (see 2.3).
- Source: <https://help.malighting.com/grandMA3/2.3/HTML/lua_objectfree_cmd.html>

> **"It is executed synchronously, and it blocks the Lua task while executing. This means that a bad command has the potential to block the system. For non-blocking, use CmdIndirect / CmdIndirectWait."**

## 2.2 `CmdIndirect` and `CmdIndirectWait`

```
CmdIndirect(string[, handle:undo[, handle:target]])
CmdIndirectWait(string[, handle:undo[, handle:target]])
```

- `CmdIndirect` queues the command and returns immediately — the command runs in the next idle slot.
- `CmdIndirectWait` queues *and* yields the coroutine until the command finishes.
- The optional `target` handle scopes the command — useful so the executing user/page is not the currently-active one.
- Source: <https://help.malighting.com/grandMA3/2.0/HTML/lua_objectfree_cmdindirectwait.html>

**Rule of thumb for this project:** use `Cmd` for short, deterministic actions (`Off Sequence 5`, `Go+ Sequence 5`). Use `CmdIndirect` if you're firing many commands in a tight loop and don't need to inspect the return.

## 2.3 The undo handle — `CreateUndo` / `CloseUndo`

```lua
local undo = CreateUndo("Bridge: cue swap")
Cmd("Off Sequence 5", undo)
Cmd("Go+ Sequence 5 Cue 3", undo)
CloseUndo(undo)
```

- All commands receiving the same undo handle are grouped into a single Oops/Undo entry, labelled with the string you passed to `CreateUndo`.
- Always `CloseUndo` — leaking handles is the standard footgun.
- Source: <https://help.malighting.com/grandMA3/2.1/HTML/lua_objectfree_createundo.html>

## 2.4 What command-line syntax works

The full grandMA3 command line is accepted. For our project:

| Command | Effect |
|---|---|
| `Go+ Sequence 5` | Advances to next cue on sequence 5 |
| `Goto Sequence 5 Cue 3` | Jumps to a specific cue |
| `Off Sequence 5` | Stops sequence 5 (releases) |
| `On Sequence 5` | Starts sequence 5 from current cue |
| `Load Sequence 5 Cue 3` | Loads a cue (next Go fires it) |
| `Pause Sequence 5` | Toggles pause on a sequence |
| `Top Sequence 5` | Jumps to first cue |
| `FaderMaster Page 1.201 At 50` | Sets fader 201 on page 1 to 50% |
| `Page +` / `Page 5` | Switches executor page |
| `SendOSC 1 "/Fader201,i,42"` | Outbound OSC (see §4) |

- Source: <https://help.malighting.com/grandMA3/2.3/HTML/command_syntax_keywords.html>
- Source: <https://help.malighting.com/grandMA3/2.2/HTML/extended_command_line.html>

## 2.5 Quoting and escaping pitfalls

- Strings with spaces must be `"double-quoted"` inside the command. From Lua you'll typically need `Cmd('Label Sequence 5 "My Show"')`.
- Locale: numeric arguments use `.` as decimal separator regardless of OS locale — MA3 sets its own.
- `$variable` substitution happens **before** the command runs. To pass a literal `$`, escape with `$$` (per the macro/keyword docs).
- Unverified — escaping rules for embedded quotes inside `Cmd` strings differ from macros; safer to build the string with `string.format` and avoid nesting quotes.

---

[Back to README](README.md)
