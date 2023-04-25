package main

import (
	"context"
	"fmt"
	"github.com/678g0/golangbook/grpc/grpc/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, _ := grpc.DialContext(context.Background(), "127.0.0.1:60002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()

	client := service.NewHelloServiceClient(conn)
	req := &service.Request{Value: "tom"}
	reply, _ := client.Hello(context.Background(), req)
	fmt.Println(reply.GetValue())
}
