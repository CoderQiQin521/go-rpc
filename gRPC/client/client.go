package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pt "rpc/gRPC/protocol"
)

const (
	srvAddr = "localhost:8989"
)

func main() {
	conn, err := grpc.Dial(srvAddr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	c := pt.NewSomeServerClient(conn)
	response, err := c.SomeMethod(context.Background(), &pt.RequestParam{Name: "xiao ming", Age: 18})
	if err != nil {
		panic(err)
	}

	fmt.Println("响应结果是", response.Result)
}
