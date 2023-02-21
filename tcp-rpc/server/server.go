package main

import (
	"net"
	"net/rpc"
	"rpc/service"
)

func main() {
	server := new(service.SomeServer)
	// 注册rpc服务
	rpc.Register(server)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":8989")
	if err != nil {
		panic(err)
	}

	// tcp监听器
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		panic(err)
	}

	for {
		// 接收请求
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(conn)
	}
}
