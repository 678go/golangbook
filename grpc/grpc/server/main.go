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

func main() {
	// 首先是通过grpc.NewServer()构造一个gRPC服务对象
	grpcServer := grpc.NewServer()
	// 然后通过gRPC插件生成的RegisterHelloServiceServer函数注册我们实现的HelloServiceImpl服务
	service.RegisterHelloServiceServer(grpcServer, &HelloServiceServer{})
	listen, _ := net.Listen("tcp", ":60002")
	_ = grpcServer.Serve(listen)
}
