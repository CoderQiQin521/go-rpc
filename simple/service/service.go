package service

import (
	"fmt"
	"rpc/simple/data"
)

type SomeServer int

func (s *SomeServer) SomeMethod(request *data.RequestParam, response *data.ResponseData) error {
	fmt.Printf("请求参数是 %v", request)
	response.Result = fmt.Sprintf("客户端请求参数是 %v", request)
	return nil
}
