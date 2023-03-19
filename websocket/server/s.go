package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net"
	"net/http"
	"strconv"
	"time"
)

type WsServer struct {
	addr    string
	upgrade *websocket.Upgrader
}

type Request struct {
	A, B int
}
type Response struct {
	Sum int
}

func NewWsServer(port int) *WsServer {
	return &WsServer{
		addr: "0.0.0.0:" + strconv.Itoa(port),
		upgrade: &websocket.Upgrader{
			HandshakeTimeout: 5 * time.Second,
			ReadBufferSize:   4096,
			WriteBufferSize:  4096,
		},
	}
}

func (ws *WsServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 将http协议升级到websocket协议
	conn, err := ws.upgrade.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("协议升级失败: ", err)
		return
	}
	fmt.Printf("和客户端%s建立了websocket连接\n", r.RemoteAddr)
	go ws.handleOneConn(conn)
}

// 处理websocket连接请求信息
func (ws *WsServer) handleOneConn(conn *websocket.Conn) {
	defer conn.Close()
	for { // 长连接
		conn.SetReadDeadline(time.Now().Add(20. * time.Second))
		var request Request
		if err := conn.ReadJSON(&request); err != nil {
			if netError, ok := err.(net.Error); ok {
				if netError.Timeout() {
					fmt.Println("发生了读超时。")
					return
				}
			}
			fmt.Println(err)
			return
		}
		if err := conn.WriteJSON(Response{request.A + request.B}); err != nil {
			fmt.Println(err)
			return
		}
	}
}

func (ws *WsServer) Start() error {
	// 这里就是执行 ServeHTTP 函数
	if err := http.ListenAndServe(ws.addr, ws); err != nil {
		return err
	}
	return nil
}

func main() {
	ws := NewWsServer(5657)
	if err := ws.Start(); err != nil {
		return
	}
}
