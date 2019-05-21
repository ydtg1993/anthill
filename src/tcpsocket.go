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
		go tcpPool.handle(conn)
	}
}

func (pool *TcpPool) handle(conn net.Conn) {
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
		switch information.Event {
		case NOTICE_EVENT:

		case BROADCAST_EVENT:

		case REGISTER_EVENT:
			pool.Workers[token] = conn
		case LOGOUT_EVENT:
			delete(pool.Workers, token)
			conn.Close()
		}
	}
}
