package main

import (
	"net"
	"net/rpc"
)

func main() {
	// 把我们的对象注册成一个rpc的 receiver
	// 其中rpc.Register函数调用会将对象类型中所有满足RPC规则的对象方法注册为RPC函数，
	// 所有注册的方法会放在“HelloService”服务空间之下  用作隔离 类似于namespace
	if err := rpc.RegisterName("HelloService", new(HelloServer)); err != nil {
		return
	}
	// 建立TCP链接
	listen, _ := net.Listen("tcp", ":60001")
	// 通过rpc.ServeConn函数在该TCP链接上为对方提供RPC服务。
	// 没Accept一个请求，就创建一个goRoutine进行处理
	for {
		accept, _ := listen.Accept()
		// 每个客户端单独一个goRoutine来处理
		go rpc.ServeConn(accept)
	}
}

type HelloServer struct{}

// Hello的逻辑 就是 将对方发送的消息前面添加一个Hello 然后返还给对方
// 由于我们是一个rpc服务, 因此参数上面还是有约束：
// 		第一个参数是请求
// 		第二个参数是响应
// 可以类比Http handler

func (h *HelloServer) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}
