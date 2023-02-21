package main

import (
	"net/http"
	"net/rpc"
	"rpc/service"
)

func main() {
	someserver := new(service.SomeServer)

	rpc.Register(someserver)
	rpc.HandleHTTP()

	err := http.ListenAndServe(":8989", nil)
	if err != nil {
		panic(err)
	}
}
