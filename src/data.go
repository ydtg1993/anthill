package src

import (
	"golang.org/x/net/websocket"
)

const (
	NOTICE_EVENT    = "notice"
	BROADCAST_EVENT = "broadcast"
	REGISTER_EVENT  = "register"
	LOGOUT_EVENT    = "logout"
)

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
	Workers map[string]bool
}

type WebsocketPool struct {
	Workers map[string]*websocket.Conn
}

type Information struct {
	Event   string `json:"event"`
	Token   string `json:"token"`
	Message string `json:"message"`
}
