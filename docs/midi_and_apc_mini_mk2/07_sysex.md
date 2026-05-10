# 7. SysEx — Device Inquiry, Introduction, RGB Lighting

The mk2 supports three meaningful SysEx flows.

## 7.1 General SysEx envelope (AKAI-specific)

```
F0 47 7F 4F <msg-id> <len-MSB> <len-LSB> <n data bytes...> F7
   ^^ ^^ ^^
   |  |  +-- Product Model ID = 0x4F (APC mini mk2)
   |  +----- SysEx Device ID = 0x7F (broadcast/single)
   +-------- AKAI Manufacturer ID = 0x47
```

Source: protocol PDF, "MIDI System Exclusive" table (lines 333–342).

## 7.2 Device Inquiry (Universal Non-Realtime, MMC)

Outbound (host → device):

```
F0 7E 00 06 01 F7
```

Inbound response (device → host) — 35 bytes, returns firmware version,
SysEx Device ID, and a 4-digit serial number plus 16-byte manufacturing data:

```
F0 7E <ch> 06 02 47 4F 00 19 <V1 V2 V3 V4> <DevID> <S1 S2 S3 S4> <Mfg1..Mfg16> F7
```

Useful for the bridge to **disambiguate two physical devices by serial** instead of
relying on port-name 2nd-occurrence (see §2). Source: protocol PDF, lines 346–445.

## 7.3 Introduction message (mode-set, called by Ableton on startup)

Outbound:

```
F0 47 7F 4F 60 00 04 00 <vH> <vL> <bugfix> F7
```

The device responds (msg-id `0x61`) with a snapshot of all 9 fader positions:

```
F0 47 7F 4F 61 00 04 <F1 F2 F3 F4 F5 F6 F7 F8 F9> F7
```

This is what Ableton's APC control surface script sends on connection. Our bridge
does **not** currently send an Introduction; the mk2 still works without it, but
sending one would let us **prime our cached fader state** without waiting for the
user to wiggle each fader. Worth adding for snappy startup.

Source: protocol PDF, lines 451–530.

## 7.4 Direct RGB lighting (24-bit per pad, range syntax)

For colors that aren't in the 128-entry palette:

```
F0 47 7F 4F 24 <total-MSB> <total-LSB>
   <startPad> <endPad>
   <Rmsb> <Rlsb> <Gmsb> <Glsb> <Bmsb> <Blsb>
   [<startPad> <endPad> <Rmsb> ... <Blsb>] ...
F7
```

- `startPad` / `endPad` ∈ `0x00..0x3F`. A range lights an inclusive run of pads
  with the same RGB.
- Each color channel is split into 7-bit MSB and 7-bit LSB (the firmware
  reassembles 14 bits per channel — though the LED hardware is presumably 8-bit;
  the 14-bit framing matches Akai's other gear).
- `total-MSB:total-LSB` is the count of *data* bytes between byte 7 and the
  trailing `F7` (i.e. `8 * num_groups` for `num_groups` pad-range groups).

This is the path to use for arbitrary cue-color matches when the palette's
nearest-neighbor lookup is too coarse. We don't currently use it.

Source: protocol PDF, lines 273–319.

## 7.5 No mode-set SysEx

There is **no** documented `Mode 0/1/2` SysEx for switching between Note Mode and
Drum Mode — those modes are entered only via the **Shift + Scene 6/7 hardware
combo**. Likewise, there is no firmware-update or factory-reset SysEx in the public
v1.0 protocol document.

---

[Back to README](README.md)
