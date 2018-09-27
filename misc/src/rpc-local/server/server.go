package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"rpc-local/service"
)

// port 1234 of the localhost
const (
	addr     = ":1234"
	protocol = "tcp"
)

func main() {
	var isHTTP bool
	if len(os.Args) > 1 && os.Args[1] == "http" {
		isHTTP = true
	}

	// register Service object's methods in the DefaultServer
	rpc.Register(new(service.Service))

	if isHTTP { // http mode
		// rpc.HandleHTTP() is required for the connection http, otherwise,
		// DialHTTP() from client side will result in unexpected HTTP response: 404 Not Found
		rpc.HandleHTTP()
		// http.ListenAndServe(addr, nil) is equivalent to
		// listener, _ := net.Listen(protocol, addr)
		// http.Serve(listener, nil)
		http.ListenAndServe(addr, nil)
	} else { // raw mode
		listener, err := net.Listen(protocol, addr)
		if err != nil {
			panic(err)
		}

		for {
			conn, err := listener.Accept()
			if err != nil {
				continue
			}
			go func() {
				fmt.Println("start serving a connection...")
				rpc.ServeConn(conn)
			}()
		}
	}
}
