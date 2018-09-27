package main

import (
	"fmt"
	"net/rpc"
	"os"
	"rpc-local/service"
)

const (
	addr     = ":1234"
	protocol = "tcp"
)

func main() {
	var isHTTP bool
	if len(os.Args) > 1 && os.Args[1] == "http" {
		isHTTP = true
	}

	var (
		client *rpc.Client
		err    error
	)

	if isHTTP {
		client, err = rpc.DialHTTP(protocol, addr)
	} else {
		client, err = rpc.Dial(protocol, addr)
	}

	if err != nil {
		panic(err)
	}

	args := &service.Args{4, 6}
	var reply int
	if err = client.Call("Service.Add", args, &reply); err != nil {
		panic(err)
	}
	fmt.Printf("Service: %d + %d = %d\n", args.Left, args.Right, reply)
}
