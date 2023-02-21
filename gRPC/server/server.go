package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	pt "rpc/gRPC/protocol"
)

const (
	srvAddr = "localhost:8989"
)

type server struct {
	pt.UnimplementedSomeServerServer
}

func (s *server) SomeMethod(ctx context.Context, in *pt.RequestParam) (*pt.ResponseData, error) {
	fmt.Println("请求参数是", in.Name)
	return &pt.ResponseData{Result: "hello " + in.Name}, nil
}

func main() {
	listen, err := net.Listen("tcp", srvAddr)
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	pt.RegisterSomeServerServer(srv, &server{})

	err = srv.Serve(listen)
	if err != nil {
		panic(err)
	}
}
