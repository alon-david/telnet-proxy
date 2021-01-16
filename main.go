package main

import (
	"github.com/reiver/go-telnet"
)

func main() {
	// TEST SERVER - SERIAL just an echo..
	go func() {
		handler := EchoHandler
		err := telnet.ListenAndServe(":4000", handler)
		if nil != err {
			panic(err)
		}
	}()

	err := telnet.ListenAndServe(":4001", DHLHandler)
	if nil != err {
		panic(err)
	}
}
