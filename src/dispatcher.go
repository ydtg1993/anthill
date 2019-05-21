package src

import (
	"golang.org/x/net/websocket"
	"net"
	"sync"
)

var TPool *TcpPool
var WPool *WebsocketPool
var WG sync.WaitGroup

func Dispatch(config *ServerConfig) {
	TPool = new(TcpPool)
	TPool.Workers = make(map[string]net.Conn)
	WPool = new(WebsocketPool)
	WPool.Workers = make(map[string]*websocket.Conn)
	WG.Add(1)
	go TcpWorker(config)
	go WebWorker(config)

	WG.Wait()
}
