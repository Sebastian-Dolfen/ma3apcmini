package midi

import (
	midi "gitlab.com/gomidi/midi/v2"
)

// MIDI channels for LED behavior
const (
	ChannelSolid     = uint8(6) // steady
	ChannelSlowBlink = uint8(7) // active
	ChannelFastBlink = uint8(8) // preloaded
)

// APC mini mk2 note layout helpers
// Grid notes are 0..63 (row-major), side buttons 112..119

// OutPort is the minimal interface we need from a MIDI Out port
type OutPort interface{ Send([]byte) error }

type APCMK2 struct {
	out OutPort
}

func NewAPCMK2(out OutPort) *APCMK2 { return &APCMK2{out: out} }

// Clear all LEDs
func (a *APCMK2) Clear() error {
	for n := uint8(0); n < 128; n++ {
		if err := a.noteOn(n, 0, 0); err != nil {
			return err
		}
	}
	return nil
}

// LightGrid lights a grid pad by grid index 0..63 with given velocity color and channel behavior
func (a *APCMK2) LightGrid(index int, velocity int, channel uint8) error {
	if index < 0 || index > 63 {
		return nil
	}
	note := uint8(index)
	return a.noteOn(note, uint8(velocity), channel)
}

// LightPage lights current page side-button (1..8)
func (a *APCMK2) LightPage(p int) error {
	// turn all off then set one on
	for i := 1; i <= 8; i++ {
		if err := a.noteOn(uint8(111+i), 0, 0); err != nil {
			return err
		}
	}
	// Use a bright green velocity to ensure visibility on MK2
	return a.noteOn(uint8(111+p), 21, 0)
}

func (a *APCMK2) noteOn(note uint8, vel uint8, ch uint8) error {
	// raw status/data bytes
	status := 0x90 | (ch & 0x0F)
	return a.out.Send([]byte{byte(status), byte(note), byte(vel)})
}

// Message classifiers
func IsGridNoteOn(m midi.Message) bool {
	var ch, note, vel uint8
	if m.GetNoteOn(&ch, &note, &vel) {
		return note <= 63 && vel > 0
	}
	return false
}

func IsGridNoteOff(m midi.Message) bool {
	var ch, note, vel uint8
	if m.GetNoteOff(&ch, &note, &vel) {
		return note <= 63
	}
	if m.GetNoteOn(&ch, &note, &vel) {
		return note <= 63 && vel == 0
	}
	return false
}

func GridIndexFromNote(m midi.Message) int {
	var ch, n, v uint8
	if m.GetNoteOn(&ch, &n, &v) {
		return int(n)
	}
	if m.GetNoteOff(&ch, &n, &v) {
		return int(n)
	}
	return -1
}

func IsSideButton(m midi.Message) bool {
	var ch, n, v uint8
	if m.GetNoteOn(&ch, &n, &v) {
		return n >= 112 && n <= 119
	}
	if m.GetNoteOff(&ch, &n, &v) {
		return n >= 112 && n <= 119
	}
	return false
}

func PageFromSide(m midi.Message) int {
	var ch, n, v uint8
	if m.GetNoteOn(&ch, &n, &v) {
		return int(n - 111)
	}
	if m.GetNoteOff(&ch, &n, &v) {
		return int(n - 111)
	}
	return 1
}
