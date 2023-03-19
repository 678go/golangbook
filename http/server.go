package main

import (
	"fmt"
	"log"
	"net"
)

// socket编程
func main() {
	tcpAddr, _ := net.ResolveTCPAddr("tcp4", "127.0.0.1:8081")
	listenTCP, _ := net.ListenTCP("tcp4", tcpAddr)
	conn, _ := listenTCP.Accept() // 可能阻塞等待客户端发起链接

	request := make([]byte, 256)
	n, _ := conn.Read(request)
	fmt.Printf("request is : %s\n", string(request[:n]))

	response := "hello " + string(request[:n])
	_, err := conn.Write([]byte(response))
	if err != nil {
		log.Println("会写失败")
	}
}
