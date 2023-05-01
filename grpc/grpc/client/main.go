package main

import (
	"context"
	"fmt"
	"github.com/678g0/golangbook/grpc/grpc/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	conn, err := grpc.DialContext(context.Background(), "127.0.0.1:60002",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithPerRPCCredentials(service.NewClientAuthInfo("admin", "admin")))
	if err != nil {
		log.Panicln(err)
	}
	defer conn.Close()

	client := service.NewHelloServiceClient(conn)

	//	普通grpc
	req := &service.Request{Value: "tom"}
	reply, err := client.Hello(context.Background(), req)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(reply.GetValue())

	// grpc的流模式
	// 1. 客户端需要先调用Channel方法获取返回的流对象
	//channelClient, err := client.Channel(context.Background())
	//if err != nil {
	//	panic(err)
	//}
	//// 2. 发送数据 通过goroutine
	//go func() {
	//	for {
	//		err = channelClient.Send(&service.Request{Value: "hil"})
	//		if err != nil {
	//			panic(err)
	//		}
	//		time.Sleep(2 * time.Second)
	//	}
	//}()
	//// 4.接收服务端返回的数据
	//for {
	//	recv, err := channelClient.Recv()
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Println(recv.GetValue())
	//}
}
