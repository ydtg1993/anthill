package src

import (
	"golang.org/x/net/websocket"
	"net"
)

var TPool *TcpPool
var WPool *WebsocketPool

func Dispatch(config *ServerConfig) {
	TPool = new(TcpPool)
	TPool.Workers = make(map[string]net.Conn)
	WPool = new(WebsocketPool)
	WPool.Workers = make(map[string]*websocket.Conn)
	go TcpWorker(config)
	go WebWorker(config)
}
