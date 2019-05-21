package src

import (
	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
)

func WebWorker(config *ServerConfig) {
	http.Handle(config.Websocket.Pattern, websocket.Handler(handle))
	if err := http.ListenAndServe(":"+config.Websocket.Port, nil); err != nil {
		fmt.Println(err)
	}
}

func handle(ws *websocket.Conn) {
	for {
		var message string
		if err := websocket.Message.Receive(ws, &message); err != nil {
			fmt.Println(err)
			break
		}

	}
}
