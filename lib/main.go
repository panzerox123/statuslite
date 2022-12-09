package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

func runRPCListener() {
	srv := rpc.NewServer()
	s := NewStatus()
	srv.Register(&s)
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		srv.Accept(listener)
	}
}

func main() {
	fmt.Println("Initializing Status-Lite module.")
	runRPCListener()
}
