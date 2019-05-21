package src

import (
	"encoding/json"
	"fmt"
	"net"
)

func TcpWorker(config *ServerConfig, tcpPool *TcpPool, websocketPool *WebsocketPool) {
	tcpServer, _ := net.ResolveTCPAddr("tcp4", config.Tcp.Pattern+":"+config.Tcp.Port)
	listener, _ := net.ListenTCP("tcp4", tcpServer)

	tcpPool.Workers = make(map[string]net.Conn)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go tcpHandle(conn, tcpPool, websocketPool)
	}
}

func tcpHandle(conn net.Conn, tcpPool *TcpPool, websocketPool *WebsocketPool) {
	for {
		information := *new(Information)
		buffer := make([]byte, 10240)
		readLen, err := conn.Read(buffer)
		if err != nil {
			fmt.Println(err)
			conn.Close()
			break
		}

		json.Unmarshal(buffer[:readLen], &information)
		token := information.Token
		message := information.Message
		switch information.Event {
		case NOTICE_EVENT:
			websocket, ok := websocketPool.Workers[token]
			if ok {
				websocket.Write([]byte(message))
			}
		case BROADCAST_EVENT:
			for _, websocket := range websocketPool.Workers {
				websocket.Write([]byte(message))
			}
		case REGISTER_EVENT:
			tcpPool.Workers[token] = conn
			websocket, ok := websocketPool.Workers[token]
			if ok {
				websocket.Write([]byte(message))
			}
		case LOGOUT_EVENT:
			delete(tcpPool.Workers, token)
			conn.Close()
			//websocket down
			websocket, ok := websocketPool.Workers[token]
			if ok {
				websocket.Close()
				delete(websocketPool.Workers, token)
			}
		}
	}
}
