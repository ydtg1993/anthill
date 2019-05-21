package src

var TPool *TcpPool
var WPool *WebsocketPool

func Dispatch(config *ServerConfig) {
	TPool = new(TcpPool)
	WPool = new(WebsocketPool)
	go TcpWorker(config)
	go WebWorker(config)
}
