package main

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"rpc/service"
)

func main() {
	server := new(service.SomeServer)
	rpc.Register(server)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":8989")
	if err != nil {
		panic(err)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)

	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}
