package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type ServerStatus struct {
	ServerStatus int
}

func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal("Client failed to connect: ", err)
	}
	var s ServerStatus
	client.Call("Status.ReturnServerStatus", 0, &s)
	fmt.Println(s.ServerStatus)
}
