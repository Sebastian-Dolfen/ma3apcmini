# 8. Multi-device patterns

When two physical APC mini mk2s are plugged into one host:

## 8.1 USB enumeration

- USB devices get unique bus/port addresses; class-compliant MIDI devices are
  enumerated by the OS in plug order.
- On Linux, both devices show up under one ALSA "card" each, named identically
  (`APC mini mk2`). Kernel suffixes them with index numbers: e.g. ALSA cards 2 and
  3. The CoreMIDI / WinMM situation is similar — same name, repeated.

## 8.2 Disambiguation strategies (best → worst)

1. **Pin by USB serial / iSerial**. On Linux:
   `udevadm info -q property -n /dev/snd/midiC2D0 | grep ID_SERIAL`. Use it to
   build a stable "device 1 = serial X, device 2 = serial Y" config. (Best.)
2. **Pin by USB bus/port**: `BUS-PORT.PORT` paths are stable across replug **as
   long as you don't change which USB jack you use**. Less robust than serial but
   doesn't require a SysEx round-trip.
3. **Use SysEx Device Inquiry** (§7.2) at startup. Send to each port, read serial
   from response, build a serial→port map. Robust on all OSes; one MIDI
   round-trip per device.
4. **Port-name n-th occurrence** (current bridge approach). Works if the OS
   returns ports in stable order. **Fragile** on macOS (CoreMIDI may reorder on
   replug) and on Windows (driver order is not guaranteed).
5. **Hard-code RtMidi index**. Brittle; breaks on any other USB-MIDI device being
   plugged in or out.

## 8.3 RtMidi / portmidi quirks

- RtMidi (Linux/ALSA) prefixes ALSA port names with `client:port`, e.g. `APC mini
  mk2:APC mini mk2 MIDI 1 28:0`. `strings.Contains(name, "APC mini")` matches
  both ports, so the occurrence counter must be paired with port-direction
  (in vs. out) and possibly with the explicit `MIDI 1` / `MIDI 2` suffix.
- RtMidi (CoreMIDI) returns just the endpoint name `APC mini mk2 Port 1`.
- RtMidi (WinMM) returns `MIDIIN2 (APC mini mk2)` for the second-port endpoints.
- **portmidi** (older library) is known to crash on hot-replug of USB MIDI
  devices; RtMidi handles this gracefully if you re-poll `PortCount()` (which our
  bridge does, every 3 s).

---

[Back to README](README.md)
