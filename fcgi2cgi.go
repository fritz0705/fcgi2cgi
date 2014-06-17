package main

import (
	"flag"
	"log"
	"net"
	"net/http/fcgi"
	"net/http/cgi"
)

func main() {
	var (
		flProto = flag.String("protocol", "tcp", "listener protocol")
		flAddr = flag.String("address", ":8902", "listener address")
	)

	flag.Parse()

	listener, err := net.Listen(*flProto, *flAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	handler := &cgi.Handler{
		Path: flag.Arg(0),
	}

	err = fcgi.Serve(listener, handler)
	if err != nil {
		log.Fatal(err)
	}
}
