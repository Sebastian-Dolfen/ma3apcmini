package osc

import (
	"fmt"
	"log"
	"net"

	osc "github.com/hypebeast/go-osc/osc"
)

type Server struct {
	addr       string
	server     *osc.Server
	onPad      func(dev int32, i int32, state int32, color string)
	onPage     func(dev int32, p int32)
	onShutdown func(dev int32)
}

func NewServer(ip string, port int) *Server {
	addr := net.JoinHostPort(ip, fmt.Sprintf("%d", port))
	s := &Server{addr: addr}
	m := osc.NewStandardDispatcher()
	// Handlers
	_ = m.AddMsgHandler("/Pad", func(msg *osc.Message) {
		// Support 3-arg format: i, state, color (defaults to dev=1)
		// and 4-arg format: dev, i, state, color
		if len(msg.Arguments) < 3 {
			return
		}
		var dev int32 = 1
		var i, st int32
		var col string
		if len(msg.Arguments) >= 4 {
			dev, _ = argI32(msg.Arguments[0])
			i, _ = argI32(msg.Arguments[1])
			st, _ = argI32(msg.Arguments[2])
			col = fmt.Sprint(msg.Arguments[3])
		} else {
			i, _ = argI32(msg.Arguments[0])
			st, _ = argI32(msg.Arguments[1])
			col = fmt.Sprint(msg.Arguments[2])
		}
		log.Printf("[OSC<-] /Pad dev=%d i=%d st=%d col=%s", dev, i, st, col)
		if s.onPad != nil {
			s.onPad(dev, i, st, col)
		}
	})
	_ = m.AddMsgHandler("/Page", func(msg *osc.Message) {
		// Support 1-arg format: p (defaults to dev=1)
		// and 2-arg format: dev, p
		if len(msg.Arguments) < 1 {
			return
		}
		var dev int32 = 1
		var p int32
		if len(msg.Arguments) >= 2 {
			dev, _ = argI32(msg.Arguments[0])
			p, _ = argI32(msg.Arguments[1])
		} else {
			p, _ = argI32(msg.Arguments[0])
		}
		log.Printf("[OSC<-] /Page dev=%d p=%d", dev, p)
		if s.onPage != nil {
			s.onPage(dev, p)
		}
	})
	_ = m.AddMsgHandler("/Shutdown", func(msg *osc.Message) {
		var dev int32 = 1
		if len(msg.Arguments) >= 1 {
			if d, ok := argI32(msg.Arguments[0]); ok {
				dev = d
			}
		}
		log.Printf("[OSC<-] /Shutdown dev=%d", dev)
		if s.onShutdown != nil {
			s.onShutdown(dev)
		}
	})
	s.server = &osc.Server{Addr: addr, Dispatcher: m}
	return s
}

func (s *Server) HandlePad(h func(dev int32, i int32, state int32, color string)) { s.onPad = h }
func (s *Server) HandlePage(h func(dev int32, p int32))                           { s.onPage = h }
func (s *Server) HandleShutdown(h func(dev int32))                                { s.onShutdown = h }

func (s *Server) ListenAndServe() error {
	log.Printf("[OSC] listening on %s", s.addr)
	return s.server.ListenAndServe()
}

func argI32(v interface{}) (int32, bool) {
	switch t := v.(type) {
	case int32:
		return t, true
	case int:
		return int32(t), true
	case float32:
		return int32(t), true
	case float64:
		return int32(t), true
	default:
		return 0, false
	}
}
