package main

import (
	"log"
	"net"
)

func main() {
	s := newServer()
	listener, err := net.Listen("tcp", "8888")
	if err != nil {
		log.Fatal("unable to start server")
	}

	defer listener.Close()
	log.Println("started server on :8888")
}
