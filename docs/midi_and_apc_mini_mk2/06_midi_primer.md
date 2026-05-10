# 6. MIDI message format primer

A Standard MIDI byte-level reminder, scoped to what our bridge handles.

## 6.1 Status bytes

The high bit (`0x80`) marks a status byte; data bytes have the high bit clear
(0..127 = `0x00..0x7F`).

| Mnemonic         | Status nibble | Status byte (ch=0) | Data bytes                |
|------------------|---------------|--------------------|---------------------------|
| Note Off         | `0x8`         | `0x80`             | note (0..127), velocity   |
| Note On          | `0x9`         | `0x90`             | note (0..127), velocity   |
| Polyphonic AT    | `0xA`         | `0xA0`             | note, pressure            |
| Control Change   | `0xB`         | `0xB0`             | controller#, value        |
| Program Change   | `0xC`         | `0xC0`             | program (1 data byte)     |
| Channel Pressure | `0xD`         | `0xD0`             | pressure (1 data byte)    |
| Pitch Bend       | `0xE`         | `0xE0`             | LSB, MSB                  |
| System Exclusive | `0xF0`        | `0xF0`             | manufacturer..F7          |

Channel masking: `status_byte & 0x0F` gives the 0-based channel (0..15);
`status_byte & 0xF0` gives the message type. Our bridge uses
`0x90 | (ch & 0x0F)` to construct LED Note Ons.

## 6.2 Note On with velocity 0 = Note Off (the "running-status convention")

A Note On (`0x9X`) with velocity 0 is, by spec, equivalent to a Note Off. This
allowed early sequencers to use running status to send a stream of `9X nn 00 nn vv`
bytes without re-emitting `0x8X` for releases. Our `IsGridNoteOff` correctly
handles both cases (NoteOn vel=0, and NoteOff).

## 6.3 14-bit CCs

The MIDI spec reserves CC 0..31 (MSB) paired with CC 32..63 (LSB) for 14-bit
controllers. **The APC mini mk2 does NOT use 14-bit CCs**: faders are 7-bit only on
CC `0x30..0x38`. The fact that those CCs are in the 0x20..0x3F range (the LSB-bank
range) is coincidental — the firmware does not emit a paired MSB.

## 6.4 Running status

Running status omits the status byte for repeated same-status messages — e.g. a
burst of Note Ons on channel 0 can be `90 nn vv nn vv nn vv …`. RtMidi /
gomidi/midi handle running status transparently; our bridge always emits a fresh
status byte per LED command, which is fine.

## 6.5 USB-MIDI framing

USB-MIDI (USB Class Spec for MIDI Devices) wraps each MIDI event in a 4-byte USB
packet with a 1-byte CIN (Code Index Number) header. The host/driver unwraps this;
RtMidi exposes the unwrapped MIDI bytes directly. SysEx >3 bytes is split into
multiple USB packets with CIN=0x4 (start/cont) and 0x5/6/7 (end with 1/2/3 bytes).
Relevant only if you're poking at USB packets directly — we are not.

---

[Back to README](README.md)
