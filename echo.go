package main

import (
	"fmt"
	"github.com/reiver/go-oi"
	"github.com/reiver/go-telnet"
	"os"
)

// EchoHandler is a simple TELNET server which "echos" back to the client any (non-command)
// data back to the TELNET client, it received from the TELNET client.
var EchoHandler telnet.Handler = internalEchoHandler{}

type internalEchoHandler struct{}

func (handler internalEchoHandler) ServeTELNET(ctx telnet.Context, w telnet.Writer, r telnet.Reader) {

	var buffer [1]byte // Seems like the length of the buffer needs to be small, otherwise will have to wait for buffer to fill up.
	p := buffer[:]

	for {

		n, err := r.Read(p)
		if n > 0 {
			fmt.Fprintln(os.Stdout, "ECHO ", p[:n], n)
			_, err := oi.LongWrite(w, p[:n])
			if err != nil {
				panic(err)
			}
		}

		if nil != err {
			break
		}
	}
}
