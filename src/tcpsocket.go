package src

import (
	"fmt"
	"net"
)

func (config *ServerConfig) TcpWorker() {
	tcpServer, _ := net.ResolveTCPAddr("tcp4", config.Tcp.Pattern+":"+config.Tcp.Port)
	listener, _ := net.ListenTCP("tcp4", tcpServer)

	tcpPool := new(TcpPool)
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
	sign := RandSign(16)
	pool.Workers[sign] = conn

	for {
		buffer := make([]byte, 10240)
		readLen, _ := conn.Read(buffer)
		message := string(buffer[:readLen])

	}
}
