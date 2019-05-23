package src

import (
	"golang.org/x/net/websocket"
	"sync"
)

var TPool *TcpPool
var WPool *WebsocketPool
var WG sync.WaitGroup

func Dispatch(config *ServerConfig) {
	TPool = new(TcpPool)
	TPool.Workers = make(map[string]bool)
	WPool = new(WebsocketPool)
	WPool.Workers = make(map[string]*websocket.Conn)
	WG.Add(1)
	go TcpWorker(config)
	go WebWorker(config)

	WG.Wait()
}
