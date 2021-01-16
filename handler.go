package main

import (
	"fmt"
	"github.com/reiver/go-oi"
	"github.com/reiver/go-telnet"
	"os"
)

// EchoHandler is a simple TELNET server which "echos" back to the client any (non-command)
// data back to the TELNET client, it received from the TELNET client.
var (
	DHLHandler telnet.Handler = internalDHLHandler{}
	id                        = 1
)

type internalDHLHandler struct{}

func (handler internalDHLHandler) ServeTELNET(ctx telnet.Context, w telnet.Writer, r telnet.Reader) {
	fmt.Fprintln(os.Stdout, id)
	output, err := ss.append(id)
	if err != nil {
		panic(err)
	}
	id += 1
	ss.once.Do(func() {
		go handler.SendHala()
	})

	var buffer [1]byte // Seems like the length of the buffer needs to be small, otherwise will have to wait for buffer to fill up.
	p := buffer[:]
	go func() {
		for {
			oi.LongWrite(w, []byte{<-output})
		}

	}()
	for {
		n, err := r.Read(p)

		if n > 0 {
			fmt.Fprintln(os.Stdout, "DHL ", p[:n], n)
			ss.input <- p[:n][0]

		}

		if nil != err {
			break
		}
	}
}
