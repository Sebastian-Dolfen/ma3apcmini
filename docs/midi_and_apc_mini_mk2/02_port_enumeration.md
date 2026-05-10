# 2. MIDI port enumeration

The mk2 enumerates as a **multi-port USB MIDI** device. The official protocol refers to
"Port 0" and "Port 1" in MIDI byte tables; the User Guide (v1.7, p.3, step 8) instructs
the user to enable Track input on **"APC mini mk2 (Port 2)"** in Ableton Live — note
the off-by-one: Live's "Port 2" is the protocol's "Port 1" (Live numbers from 1, AKAI
numbers from 0). Source: /tmp/apc_mk2_docs/userguide.txt line 57.

## What each port carries

| Protocol name | Live name | What rides on it |
|---------------|-----------|------------------|
| Port 0        | Port 1    | All "control surface" traffic: pad note-on/off in Session/Drum mode, side/track button notes, fader CCs, all LED feedback (Note On with channel-encoded behavior + velocity-encoded color), all SysEx (Device Inquiry, Introduction, RGB Lighting). |
| Port 1        | Port 2    | **Note Mode only.** When the user enters Note Mode (Shift + Scene 7), the clip-pad note-on/off messages are routed to Port 1, MIDI CH 00, with the chromatic-scale note numbers — so Live's instrument tracks see real notes, not pad indices. |

Source: AKAI Communications Protocol v1.0, "Pads/Buttons" Notes column
(/tmp/apc_mk2_docs/protocol.txt lines 579–584):

> *Port 0, MIDI CH 00–0F used when in Session View.
> *Port 0, MIDI CH 09 used when in Drum Mode.
> *Port 1, MIDI CH 00 used in Note Mode.

## OS-specific port names

Empirically (cross-referenced from the Ardour and Bome forum threads):

- **Linux / ALSA**: the device exposes two ALSA "ports" under one client.
  Common names: `APC mini mk2 MIDI 1` (= protocol Port 0, control surface) and
  `APC mini mk2 MIDI 2` (= protocol Port 1, Note Mode). Older firmware / certain
  distros may expose them as `Contr` and `Notes`.
  Source: https://discourse.ardour.org/t/akai-apc-mini-mk2/109765
- **macOS / CoreMIDI**: a single CoreMIDI device "APC mini mk2" with two endpoints
  named `APC mini mk2 Port 1` and `APC mini mk2 Port 2`.
- **Windows / WinMM (legacy)**: "APC mini mk2" and "MIDIIN2 (APC mini mk2)" /
  "MIDIOUT2 (APC mini mk2)". The duplicated-name phenomenon is a Windows MIDI driver
  artifact; on WinRT/UWP MIDI the names are clean.

## How our bridge addresses ports

`go/cmd/apc-mini-bridge/main.go` opens an RtMidi port whose name **contains** the
configured `midi_in` / `midi_out` substring, with an `inOccur` / `outOccur` counter
controlling which match to take (1 = first match, 2 = second match). For a single
device our bridge just wants protocol Port 0 (control surface), so the **first**
"APC mini" occurrence is correct. For a second physical device on the same host, the
2nd-occurrence trick selects the second device's first match — provided RtMidi
returns ports in stable USB-enumeration order, which it does on Linux/ALSA but is
**not** guaranteed on macOS or Windows.

**DISCREPANCY / RISK**: "2nd occurrence" is **not** a reliable canonical strategy
for multi-device. Better strategies:

- **USB serial / iSerialNumber**: AKAI's Device Inquiry response (SysEx `F0 7E 00 06
  01 F7`) returns a 4-byte serial — see §7. The bridge could use it to disambiguate.
- **PortMIDI / RtMidi index** with a config-file pin to a specific index after a
  one-time `--list-ports` discovery.
- On Linux: read `/sys/bus/usb/devices/.../serial` and bind by serial.

For now: 2nd-occurrence works on Linux when both devices are plugged into stable USB
ports; flag for replacement before shipping multi-device.

---

[Back to README](README.md)
