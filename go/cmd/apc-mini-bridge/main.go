package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	osc "github.com/hypebeast/go-osc/osc"
	rtmidi "github.com/mattrtaylor/go-rtmidi"
	"gopkg.in/yaml.v3"

	"ma3apcmini/internal/color"
	mididev "ma3apcmini/internal/midi"
	oscsrv "ma3apcmini/internal/osc"
)

// ANSI color codes for log prefixes
const (
	colorReset   = "\x1b[0m"
	colorRed     = "\x1b[31m"
	colorGreen   = "\x1b[32m"
	colorYellow  = "\x1b[33m"
	colorBlue    = "\x1b[34m"
	colorMagenta = "\x1b[35m"
	colorCyan    = "\x1b[36m"
)

var useColor = true

func pfx(tag, col string) string {
	if useColor {
		return fmt.Sprintf("%s[%s]%s ", col, tag, colorReset)
	}
	return fmt.Sprintf("[%s] ", tag)
}

func logInfo(format string, args ...interface{}) { log.Printf(pfx("INFO", colorGreen)+format, args...) }
func logWarn(format string, args ...interface{}) {
	log.Printf(pfx("WARN", colorYellow)+format, args...)
}
func logErr(format string, args ...interface{})  { log.Printf(pfx("ERR", colorRed)+format, args...) }
func logMIDI(format string, args ...interface{}) { log.Printf(pfx("MIDI", colorBlue)+format, args...) }
func logOSCIn(format string, args ...interface{}) {
	log.Printf(pfx("OSC<-", colorMagenta)+format, args...)
}
func logOSCOut(format string, args ...interface{}) {
	log.Printf(pfx("OSC->", colorCyan)+format, args...)
}

// outWrap adapts rtmidi.MIDIOut to internal/midi.OutPort
type outWrap struct{ d rtmidi.MIDIOut }

// Send implements midi.OutPort using RtMidi's SendMessage
func (o outWrap) Send(b []byte) error { return o.d.SendMessage(b) }

// Logo8x8Dev1 defines an 8x8 matrix of colors (hex #RRGGBB) to display
// on APC mini device 1 when no OSC device is actively driving LEDs.
// Edit the values below to customize the logo. Use "#000000" for off.
var Logo8x8Dev1 = [8][8]string{
	// Row 0 (top) .. Row 7 (bottom)
	{"#000000", "#000000", "#000000", "#000000", "#000000", "#000000", "#000000", "#000000"},
	{"#000000", "#4CFF4C", "#4CFF4C", "#000000", "#000000", "#4CFF4C", "#4CFF4C", "#000000"},
	{"#000000", "#4CFF4C", "#000000", "#4CFF4C", "#4CFF4C", "#000000", "#4CFF4C", "#000000"},
	{"#000000", "#4CFF4C", "#4CFF4C", "#000000", "#000000", "#4CFF4C", "#4CFF4C", "#000000"},
	{"#000000", "#000000", "#000000", "#000000", "#000000", "#000000", "#000000", "#000000"},
	{"#000000", "#00A9FF", "#00A9FF", "#00A9FF", "#00A9FF", "#00A9FF", "#00A9FF", "#000000"},
	{"#000000", "#00A9FF", "#000000", "#000000", "#000000", "#000000", "#00A9FF", "#000000"},
	{"#000000", "#00A9FF", "#00A9FF", "#00A9FF", "#00A9FF", "#00A9FF", "#00A9FF", "#000000"},
}

// Logo8x8Dev2 defines an alternative 8x8 matrix for APC mini device 2.
var Logo8x8Dev2 = [8][8]string{
	{"#000000", "#000000", "#000000", "#000000", "#000000", "#000000", "#000000", "#000000"},
	{"#000000", "#FF4C4C", "#FF4C4C", "#000000", "#000000", "#FF4C4C", "#FF4C4C", "#000000"},
	{"#000000", "#FF4C4C", "#000000", "#FF4C4C", "#FF4C4C", "#000000", "#FF4C4C", "#000000"},
	{"#000000", "#FF4C4C", "#FF4C4C", "#000000", "#000000", "#FF4C4C", "#FF4C4C", "#000000"},
	{"#000000", "#000000", "#000000", "#000000", "#000000", "#000000", "#000000", "#000000"},
	{"#000000", "#874CFF", "#874CFF", "#874CFF", "#874CFF", "#874CFF", "#874CFF", "#000000"},
	{"#000000", "#874CFF", "#000000", "#000000", "#000000", "#000000", "#874CFF", "#000000"},
	{"#000000", "#874CFF", "#874CFF", "#874CFF", "#874CFF", "#874CFF", "#874CFF", "#000000"},
}

// showLogo renders the given 8x8 matrix to the 8x8 grid using solid channel.
// Indexing: [row][col] with row 0 = top, col 0 = left. Note numbers are row-major 0..63.
func showLogo(apc *mididev.APCMK2, mat [8][8]string) {
	for r := 0; r < 8; r++ {
		for c := 0; c < 8; c++ {
			hex := mat[r][c]
			vel := color.LookupVelocity(hex)
			idx := r*8 + c
			_ = apc.LightGrid(idx, vel, mididev.ChannelSolid)
		}
	}
}

type Config struct {
	MIDIIn        string `yaml:"midi_in"`
	MIDIOut       string `yaml:"midi_out"`
	MIDIInIndex   int    `yaml:"midi_in_index"`  // optional explicit input port index (>=0)
	MIDIOutIndex  int    `yaml:"midi_out_index"` // optional explicit output port index (>=0)
	MIDIIn2       string `yaml:"midi_in2"`
	MIDIOut2      string `yaml:"midi_out2"`
	MIDIIn2Index  int    `yaml:"midi_in2_index"`
	MIDIOut2Index int    `yaml:"midi_out2_index"`
	LocalIP       string `yaml:"localip"`
	Local         int    `yaml:"localport"`
	RemoteIP      string `yaml:"remoteip"`
	Remote        int    `yaml:"remoteport"`
}

func loadConfig(path string) (Config, error) {
	// defaults
	cfg := Config{
		MIDIIn:        "APC mini mk2 1",
		MIDIOut:       "APC mini mk2 2",
		MIDIInIndex:   -1,
		MIDIOutIndex:  -1,
		MIDIIn2:       "APC mini mk2 3", // optional, leave empty to disable
		MIDIOut2:      "APC mini mk2 4", // optional, leave empty to disable
		MIDIIn2Index:  -1,
		MIDIOut2Index: -1,
		LocalIP:       "127.0.0.1",
		Local:         8001,
		RemoteIP:      "127.0.0.1",
		Remote:        8000,
	}
	f, err := os.Open(path)
	if err != nil {
		return cfg, nil // allow missing file, keep defaults
	}
	defer f.Close()
	dec := yaml.NewDecoder(f)
	if err := dec.Decode(&cfg); err != nil {
		return cfg, err
	}
	return cfg, nil
}

func resolveAddr(ip string, port int) string {
	return net.JoinHostPort(ip, fmt.Sprintf("%d", port))
}

func main() {
	var cfgPath string
	flag.StringVar(&cfgPath, "config", "config.yaml", "path to config.yaml")
	var noColor bool
	flag.BoolVar(&noColor, "no-color", false, "disable ANSI color in logs")
	flag.Parse()

	cfg, err := loadConfig(cfgPath)
	if err != nil {
		log.Fatalf("config error: %v", err)
	}
	useColor = !noColor

	type apcDevice struct {
		id        int
		inName    string
		outName   string
		inOccur   int
		outOccur  int
		in        rtmidi.MIDIIn
		out       rtmidi.MIDIOut
		apc       *mididev.APCMK2
		lastMsg   int64 // unix nanos of last incoming MIDI
		openedIn  bool
		openedOut bool
	}

	mkIn := func() rtmidi.MIDIIn {
		d, err := rtmidi.NewMIDIInDefault()
		if err != nil {
			log.Fatalf("rtmidi in: %v", err)
		}
		return d
	}
	mkOut := func() rtmidi.MIDIOut {
		d, err := rtmidi.NewMIDIOutDefault()
		if err != nil {
			log.Fatalf("rtmidi out: %v", err)
		}
		return d
	}

	// Log all detected MIDI ports up front for diagnostics
	listPorts := func() {
		inTmp := mkIn()
		defer inTmp.Close()
		outTmp := mkOut()
		defer outTmp.Close()
		if n, _ := inTmp.PortCount(); n > 0 {
			logInfo("Detected MIDI IN ports (%d):", n)
			for i := 0; i < n; i++ {
				name, _ := inTmp.PortName(i)
				logMIDI("IN[%d]: %s", i, name)
			}
		} else {
			logWarn("Detected MIDI IN ports: 0")
		}
		if n, _ := outTmp.PortCount(); n > 0 {
			logInfo("Detected MIDI OUT ports (%d):", n)
			for i := 0; i < n; i++ {
				name, _ := outTmp.PortName(i)
				logMIDI("OUT[%d]: %s", i, name)
			}
		} else {
			logWarn("Detected MIDI OUT ports: 0")
		}
	}
	listPorts()

	dev1 := &apcDevice{id: 1, inName: cfg.MIDIIn, outName: cfg.MIDIOut, inOccur: 1, outOccur: 1, in: mkIn(), out: mkOut()}
	// Always create dev2. If midi_in2/out2 not provided, fallback to same name but 2nd occurrence.
	dev2InName := cfg.MIDIIn
	dev2OutName := cfg.MIDIOut
	dev2InOccur := 2
	dev2OutOccur := 2
	if cfg.MIDIIn2 != "" {
		dev2InName = cfg.MIDIIn2
		dev2InOccur = 1
	}
	if cfg.MIDIOut2 != "" {
		dev2OutName = cfg.MIDIOut2
		dev2OutOccur = 1
	}
	dev2 := &apcDevice{id: 2, inName: dev2InName, outName: dev2OutName, inOccur: dev2InOccur, outOccur: dev2OutOccur, in: mkIn(), out: mkOut()}
	defer dev1.in.Close()
	defer dev1.out.Close()
	defer dev2.in.Close()
	defer dev2.out.Close()

	// OSC server for LED updates from plugin
	srv := oscsrv.NewServer(cfg.LocalIP, cfg.Local)
	// Register handlers for LED updates (supports dev routing)
	srv.HandlePad(func(dev int32, i int32, state int32, colorHex string) {
		target := dev1
		if dev == 2 && dev2 != nil {
			target = dev2
		}
		if target == nil || target.apc == nil {
			return
		}
		logOSCIn("/Pad dev=%d i=%d state=%d color=%s -> route to dev%d", dev, i, state, colorHex, target.id)
		ch := mididev.ChannelSolid
		switch state {
		case 1:
			ch = mididev.ChannelFastBlink // preloaded - fast blink
		case 2:
			ch = mididev.ChannelSlowBlink // active - slow blink
		default:
			ch = mididev.ChannelSolid
		}
		vel := color.LookupVelocity(colorHex)
		if err := target.apc.LightGrid(int(i), vel, ch); err != nil {
			logErr("light grid (dev %d): %v", target.id, err)
		}
	})
	srv.HandlePage(func(dev int32, p int32) {
		target := dev1
		if dev == 2 && dev2 != nil {
			target = dev2
		}
		if target == nil || target.apc == nil {
			return
		}
		logOSCIn("/Page dev=%d p=%d -> route to dev%d (clear+set)", dev, p, target.id)
		// Clear and then set to make sure stale logo LEDs are removed
		_ = target.apc.Clear()
		if err := target.apc.LightPage(int(p)); err != nil {
			logErr("light page (dev %d): %v", target.id, err)
		}
	})

	// When plugin shuts down, restore logo on the corresponding device
	srv.HandleShutdown(func(dev int32) {
		target := dev1
		if dev == 2 {
			target = dev2
		}
		if target == nil || target.apc == nil {
			return
		}
		logOSCIn("/Shutdown dev=%d -> redraw logo on dev%d", dev, target.id)
		_ = target.apc.Clear()
		if target.id == 2 {
			showLogo(target.apc, Logo8x8Dev2)
		} else {
			showLogo(target.apc, Logo8x8Dev1)
		}
	})

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("OSC listen: %v", err)
		}
	}()

	// OSC client to send input events to plugin
	client := osc.NewClient(cfg.RemoteIP, cfg.Remote)

	// MIDI input: translate to OSC (per-device callbacks)
	registerCB := func(d *apcDevice) {
		_ = d.in.SetCallback(func(_ rtmidi.MIDIIn, data []byte, _ float64) {
			d.lastMsg = time.Now().UnixNano()
			if len(data) < 3 {
				return
			}
			status := data[0]
			st := status & 0xF0
			note := int(data[1])
			vel := int(data[2])
			// Always include device id in OSC to support two devices consistently
			{
				switch st {
				case 0x90: // Note On (vel==0 -> off)
					if note >= 0 && note <= 63 {
						if vel > 0 {
							logMIDI("dev%d NOTE ON pad=%d vel=%d", d.id, note, vel)
							client.Send(osc.NewMessage("/Pad", int32(d.id), int32(note), int32(1)))
							logOSCOut("/Pad dev=%d i=%d state=1", d.id, note)
							if note == 63 {
								client.Send(osc.NewMessage("/Trigger", int32(1)))
								logOSCOut("/Trigger 1")
							}
						} else {
							logMIDI("dev%d NOTE ON(0->off) pad=%d", d.id, note)
							client.Send(osc.NewMessage("/Pad", int32(d.id), int32(note), int32(0)))
							logOSCOut("/Pad dev=%d i=%d state=0", d.id, note)
						}
					} else if note >= 112 && note <= 119 {
						p := note - 111
						client.Send(osc.NewMessage("/Page", int32(d.id), int32(p)))
						logMIDI("dev%d SIDE page=%d", d.id, p)
						logOSCOut("/Page dev=%d p=%d", d.id, p)
					}
				case 0x80: // Note Off
					if note >= 0 && note <= 63 {
						logMIDI("dev%d NOTE OFF pad=%d", d.id, note)
						client.Send(osc.NewMessage("/Pad", int32(d.id), int32(note), int32(0)))
						logOSCOut("/Pad dev=%d i=%d state=0", d.id, note)
					}
				}
			}
		})
	}
	// do not pre-register callbacks; scanner registers after input open

	// periodic scanner: ensures ports are opened when present and logos are (re)drawn
	scanner := func(d *apcDevice) {
		ticker := time.NewTicker(3 * time.Second)
		doScan := func() {
			// input
			presentIn := false
			var inIdx int
			if n, _ := d.in.PortCount(); n > 0 {
				match := 0
				for i := 0; i < n; i++ {
					name, _ := d.in.PortName(i)
					if strings.Contains(strings.ToLower(name), strings.ToLower(d.inName)) {
						match++
						if match == d.inOccur {
							presentIn = true
							inIdx = i
							break
						}
					}
				}
			}
			// If explicit index configured, prefer it
			if d.id == 1 && cfg.MIDIInIndex >= 0 {
				presentIn, inIdx = true, cfg.MIDIInIndex
			}
			if d.id == 2 && cfg.MIDIIn2Index >= 0 {
				presentIn, inIdx = true, cfg.MIDIIn2Index
			}

			if presentIn && !d.openedIn {
				if err := d.in.OpenPort(inIdx, fmt.Sprintf("apc-mini-in-%d", d.id)); err == nil {
					d.openedIn = true
					registerCB(d)
					logInfo("MIDI IN opened for dev %d at index %d (name contains '%s', occurrence %d)", d.id, inIdx, d.inName, d.inOccur)
				}
			}
			if !presentIn && d.openedIn {
				_ = d.in.Close()
				d.in = mkIn()
				d.openedIn = false
				logWarn("MIDI IN closed for dev %d (device disappeared)", d.id)
			}

			// output
			presentOut := false
			var outIdx int
			if n, _ := d.out.PortCount(); n > 0 {
				match := 0
				for i := 0; i < n; i++ {
					name, _ := d.out.PortName(i)
					if strings.Contains(strings.ToLower(name), strings.ToLower(d.outName)) {
						match++
						if match == d.outOccur {
							presentOut = true
							outIdx = i
							break
						}
					}
				}
			}
			// If explicit index configured, prefer it
			if d.id == 1 && cfg.MIDIOutIndex >= 0 {
				presentOut, outIdx = true, cfg.MIDIOutIndex
			}
			if d.id == 2 && cfg.MIDIOut2Index >= 0 {
				presentOut, outIdx = true, cfg.MIDIOut2Index
			}

			if presentOut && !d.openedOut {
				if err := d.out.OpenPort(outIdx, fmt.Sprintf("apc-mini-out-%d", d.id)); err == nil {
					d.openedOut = true
					d.apc = mididev.NewAPCMK2(outWrap{d: d.out})
					// Draw logo initially; will be overridden by any OSC traffic.
					_ = d.apc.Clear()
					if d.id == 2 {
						showLogo(d.apc, Logo8x8Dev2)
					} else {
						showLogo(d.apc, Logo8x8Dev1)
					}
					logInfo("MIDI OUT opened for dev %d at index %d (name contains '%s', occurrence %d); logo drawn", d.id, outIdx, d.outName, d.outOccur)
				}
			}
			if !presentOut && d.openedOut {
				_ = d.out.Close()
				d.out = mkOut()
				d.openedOut = false
				d.apc = nil
				logWarn("MIDI OUT closed for dev %d (device disappeared)", d.id)
			}
		}
		// run immediately, then on ticker
		doScan()
		for range ticker.C {
			doScan()
		}
	}
	go scanner(dev1)
	go scanner(dev2)

	// Keep running
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, syscall.SIGTERM)
	<-sigc
	if dev1.apc != nil {
		_ = dev1.apc.Clear()
	}
	if dev2 != nil && dev2.apc != nil {
		_ = dev2.apc.Clear()
	}
}
