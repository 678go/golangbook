package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	//client, _ := rpc.Dial("tcp", ":60001")
	//// 然后通过client.Call调用具体的RPC方法
	//// 在调用client.Call时:
	//// 		第一个参数是用点号链接的RPC服务名字和方法名字，
	//// 		第二个参数是 请求参数
	////      第三个是请求响应, 必须是一个指针, 有底层rpc服务帮你赋值
	//var reply string
	//if err := client.Call("HelloService.Hello", "hello", &reply); err != nil {
	//	return
	//}
	//fmt.Println(reply)
	var reply string
	_ = NewHelloClient("tcp", ":60001").Hello("你好", &reply)
	fmt.Println(reply)
}

type HelloClient struct {
	client *rpc.Client
}

func (h *HelloClient) Hello(request string, reply *string) error {
	return h.client.Call("HelloService.Hello", request, &reply)
}

func NewHelloClient(network string, addr string) *HelloClient {
	client, _ := rpc.Dial("tcp", ":60001")
	return &HelloClient{
		client: client,
	}
}
