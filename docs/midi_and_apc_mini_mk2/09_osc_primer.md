# 9. OSC primer (relevant to our hand-rolled Lua encoder)

OSC 1.0 is the wire format the MA3 plugin and the Go bridge use over UDP.

## 9.1 Address pattern

- ASCII string starting with `/`. Slash-separated segments. Examples we use:
  `/Pad`, `/Page`, `/Trigger`, `/Shutdown`.
- The string is NUL-terminated and **padded with NULs to a 4-byte boundary**.
  - `/Pad` is 4 bytes → needs 1 NUL terminator + 3 padding NULs = 8 bytes total.
  - `/Page` is 5 bytes → 1 NUL terminator + 2 padding NULs = 8 bytes total.

## 9.2 Type-tag string

- Begins with `,` (comma), followed by one tag char per argument:
  - `i` = int32 (big-endian, signed)
  - `f` = float32 (IEEE 754 big-endian)
  - `s` = OSC-string (NUL-terminated, NUL-padded to 4-byte boundary)
  - `b` = blob (int32 size, then bytes, padded to 4-byte boundary)
- Examples we send: `,iii` (3 ints), `,iiii` (4 ints), `,iiis` (3 ints + string),
  `,is` (int + string).
- The type-tag string is itself an OSC-string: NUL-terminated, NUL-padded to 4
  bytes. A `,iii` tag is 5 bytes raw → padded to 8.

## 9.3 Argument encoding

- **int32**: 4 bytes, big-endian, two's complement.
- **float32**: 4 bytes, big-endian IEEE 754.
- **string**: bytes + NUL + NULs to 4-byte boundary. *The NUL terminator counts
  toward the padding.* A 4-byte string `/Pad` with terminator is 5 bytes raw →
  padded with 3 NULs to 8.
- **blob**: int32 length + raw bytes + NUL pad to 4-byte boundary.

## 9.4 Bundles vs. messages

- A **message** is `<address-pattern> <type-tag-string> <args...>` concatenated.
- A **bundle** starts with `#bundle\0` (8 bytes), then a 64-bit OSC-timetag, then
  one or more length-prefixed sub-messages (int32 length + message bytes).
- We only use single messages, never bundles. UDP datagram boundary == OSC
  message boundary.

## 9.5 Common pitfalls when hand-rolling (and what `apc_color.lua::osc_msg` does)

The plugin's encoder (`ma3/lua/apc_color.lua`, `osc_msg`):

```lua
local function osc_pad4(s)
  local rem = (#s + 1) % 4         -- +1 accounts for NUL terminator
  if rem == 0 then return s .. string.char(0) end
  local pad = 4 - rem
  return s .. string.char(0) .. string.rep('\0', pad)
end

local function osc_msg(address, ...)
  -- Builds: pad4(address) .. pad4(types) .. concat(args_bin)
  -- Each int arg: be32(math.floor(v))
  -- Each string arg: pad4(tostring(v))
end
```

**Pitfalls (pre-checked against the plugin code)**:

- **NUL terminator is mandatory** even when the string is already a multiple of 4
  bytes. (The plugin does this correctly: `if rem == 0 then return s .. \0` adds
  exactly 1 NUL → length stays multiple of 4. Note: this is **wrong by spec** —
  when `(#s + 1) % 4 == 0`, the correct padding is `s .. \0\0\0\0` (4 NULs), not
  1. Spec requires "at least one NUL, padded to 4-byte boundary". A 3-byte string
  + 1 NUL = 4 bytes total, which IS the correct case the plugin handles. A 4-byte
  string + 1 NUL = 5 bytes, the plugin would pad to 8 with 3 more NULs. So the
  algorithm is actually correct: `rem = (#s+1) % 4`, if rem==0 it means `#s+1` is
  already a multiple of 4 → add only the 1 NUL and you're done. **OK.**)
- Plugin only supports `i` (int) and `s` (string) — no float, no blob, no nested
  bundles. Sufficient for our `/Pad`, `/Page`, `/Trigger`, `/Shutdown`.
- Endianness: be32 is implemented manually with `math.floor(i / 16777216) % 256`
  etc. — correct big-endian.
- **No bundle support**: the plugin sends one OSC message per UDP datagram, which
  is what the bridge expects.
- The plugin's `parse_osc_packet` is a **very limited parser** that ignores the
  type-tag string and reads ints from fixed offsets at the **end** of the packet.
  This is fragile — any future change to the bridge's outbound format (e.g. adding
  a string color argument) will break the plugin's reader. Suggested cleanup: do
  proper type-tag parsing.

## 9.6 Wire-format example (a `/Pad` message)

`/Pad 5 2 16711680` (pad 5, state 2 = active, color = 0xFF0000):

```
2F 50 61 64 00 00 00 00            ; "/Pad" + NUL + 3 pad → 8 bytes
2C 69 69 69 00 00 00 00            ; ",iii" + NUL + 3 pad → 8 bytes
00 00 00 05                        ; int32 5
00 00 00 02                        ; int32 2
00 FF 00 00                        ; int32 16711680 = 0x00FF0000
```

Total: 28 bytes UDP payload.

---

[Back to README](README.md)
