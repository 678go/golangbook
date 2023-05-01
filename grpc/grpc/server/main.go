package main

import (
	"context"
	"github.com/678g0/golangbook/grpc/grpc/service"
	"google.golang.org/grpc"
	"net"
)

type HelloServiceServer struct {
	service.UnimplementedHelloServiceServer
}

// Hello 重写了UnimplementedHelloServiceServer的hello方法
func (h *HelloServiceServer) Hello(ctx context.Context, req *service.Request) (*service.Reply, error) {
	return &service.Reply{Value: "hello" + req.Value}, nil
}

// 在服务端扩展HelloServiceServer结构体实现Channel方法

func (h *HelloServiceServer) Channel(stream service.HelloService_ChannelServer) error {
	for {
		// 接收请求
		recv, _ := stream.Recv()
		// 返回响应
		if err := stream.Send(&service.Reply{Value: "hell0 " + recv.GetValue()}); err != nil {
			return err
		}
	}
}

func main() {
	// 首先是通过grpc.NewServer()构造一个gRPC服务对象
	// 在grpc启动的时候添加
	grpcServer := grpc.NewServer(
		// 添加认证中间件, 如果有多个中间件需要添加 使用ChainUnaryInterceptor
		grpc.UnaryInterceptor(service.NewGrpcAuthUnaryServerInterceptor()),
	)
	// 然后通过gRPC插件生成的RegisterHelloServiceServer函数注册我们实现的HelloServiceImpl服务
	service.RegisterHelloServiceServer(grpcServer, &HelloServiceServer{})
	listen, _ := net.Listen("tcp", ":60002")
	_ = grpcServer.Serve(listen)
}
