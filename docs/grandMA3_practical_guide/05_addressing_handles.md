# 5. Object addressing and handles

> Targets MA3 2.3.2.0

## 5.1 Handle representation

A "handle" in MA3 Lua is a `light_userdata` — opaque, but printable as `H#<id>` in dumps. Handles are not stable across showfile loads.

- Source: <https://help.malighting.com/grandMA3/2.0/HTML/lua_handle.html>

## 5.2 `Addr`, `AddrNative`, `ToAddr`

All three convert a handle to an address string, but differ in *what kind* of address:

| Function | Returns | Example |
|---|---|---|
| `Addr(h)` | Numbered, root-relative | `13.5.7` |
| `AddrNative(h)` | Named, root-relative | `Sequences."My Show".Cue 3` |
| `ToAddr(h)` | Address string suitable for command-line use | `Sequence 5 Cue 3` |

- All three accept an optional second handle to make the result *relative to that base*, which must lie on the path from root to the target.
- A boolean third arg toggles between `Addr` and `ToAddr` numbering when they differ.
- Source: <https://help.malighting.com/grandMA3/2.1/HTML/lua_object_toaddr.html>
- Source: <https://help2.malighting.com/Page/grandMA3/lua_addr/en/1.9>

**Rule of thumb:** use `ToAddr` when you're going to feed the result back into `Cmd`. Use `AddrNative` for human-readable logs.

## 5.3 `FromAddr`

```
FromAddr(string[, handle]) -> handle
```

Resolves a string like `"Sequence 5"` or `"Sequence 5 Cue 3"` to a handle. Optional second handle scopes the lookup.

- Source: <https://help.malighting.com/grandMA3/2.0/HTML/lua_objectfree.html#FromAddr>

---

[Back to README](README.md)
