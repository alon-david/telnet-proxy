package main

import (
	"fmt"
	"github.com/reiver/go-telnet"
	"os"
	"sync"
)

type session struct {
	conn    *telnet.Conn
	outputs map[int]chan byte
	input   chan byte
	once    *sync.Once
}

var (
	ss *session
)

func init() {
	ss = new(session)
	ss.input = make(chan byte)
	ss.outputs = make(map[int]chan byte, 10)
	ss.once = new(sync.Once)
}

func (s *session) append(id int) (chan byte, error) {
	ch := make(chan byte)
	s.outputs[id] = ch
	return ch, nil
}

func (handler internalDHLHandler) SendHala() {
	fmt.Fprintln(os.Stdout, "HALA")
	conn, err := telnet.DialTo("localhost:4000")
	if err != nil {
		panic(err)
	}
	ss.conn = conn
	go func() {
		for {
			w := <-ss.input
			fmt.Fprintln(os.Stdout, "W", w)
			_, err := ss.conn.Write([]byte{w})
			if err != nil {
				panic(err)
			}
		}
	}()

	for {
		var buffer2 [1]byte // Seems like the length of the buffer needs to be small, otherwise will have to wait for buffer to fill up.
		p2 := buffer2[:]
		n2, err := ss.conn.Read(p2)
		if err != nil {
			panic(err)
		}
		for key := range ss.outputs {
			ss.outputs[key] <- p2[:n2][0]
		}
	}

}
