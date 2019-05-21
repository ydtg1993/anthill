package src

import (
	"encoding/json"
	"fmt"
	"net"
)

func TcpWorker(config *ServerConfig) {
	tcpServer, _ := net.ResolveTCPAddr("tcp4", config.Tcp.Pattern+":"+config.Tcp.Port)
	listener, _ := net.ListenTCP("tcp4", tcpServer)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go tcpHandle(conn)
	}
}

func tcpHandle(conn net.Conn) {
	for {
		information := *new(Information)
		buffer := make([]byte, 10240)
		readLen, err := conn.Read(buffer)
		if err != nil {
			fmt.Println(err)
			conn.Close()
			break
		}
		fmt.Println(string(buffer))

		json.Unmarshal(buffer[:readLen], &information)
		token := information.Token
		message := information.Message
		switch information.Event {
		case NOTICE_EVENT:
			websocket, ok := WPool.Workers[token]
			if ok {
				websocket.Write([]byte(message))
			} else {
				information := Information{
					Event:   "close",
					Token:   token,
					Message: "websocket closed",
				}
				message, err := json.Marshal(information)
				if err == nil {
					conn.Write(message)
				}
			}
		case BROADCAST_EVENT:
			for _, websocket := range WPool.Workers {
				websocket.Write([]byte(message))
			}
		case REGISTER_EVENT:
			TPool.Workers[token] = conn
			websocket, ok := WPool.Workers[token]
			if ok {
				websocket.Write([]byte(message))
			}
		case LOGOUT_EVENT:
			delete(TPool.Workers, token)
			conn.Close()
			//websocket down
			websocket, ok := WPool.Workers[token]
			if ok {
				websocket.Close()
				delete(WPool.Workers, token)
			}
			return
		}
	}
}
