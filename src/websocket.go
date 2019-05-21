package src

import (
	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
)

func (config *ServerConfig) WebWorker() {
	http.Handle(config.Websocket.Pattern, websocket.Handler(Echo))
	if err := http.ListenAndServe(":"+config.Websocket.Port, nil); err != nil {
		fmt.Println(err)
	}
}

func Echo(ws *websocket.Conn) {
	for {
		var message string
		if err := websocket.Message.Receive(ws, &message); err != nil {
			fmt.Println("can't receive")
			break
		}

	}
}
