package main

import (
	"fmt"
	"net/rpc"
	"rpc/data"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:8989")
	if err != nil {
		panic(err)
	}

	param := &data.RequestParam{Param1: 1, Param2: 2}
	result := &data.ResponseData{}

	err = client.Call("SomeServer.SomeMethod", param, result)
	if err != nil {
		panic(err)
	}

	fmt.Printf("请求参数是: %v, 返回结果是: %v", param, result.Result)
}
