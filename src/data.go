package src

import "net"

type ServerConfig struct {
	Tcp       Tcp
	Websocket Websocket
}

type Tcp struct {
	Pattern string `json:"pattern"`
	Port    string `json:"port"`
}

type Websocket struct {
	Pattern string `json:"pattern"`
	Port    string `json:"port"`
}

type TcpPool struct {
	Workers map[string]net.Conn
}
