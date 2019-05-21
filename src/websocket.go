package src

import (
	"encoding/json"
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
		var data []byte
		if err := websocket.Message.Receive(ws, &data); err != nil {
			fmt.Println(err)
			ws.Close()
			break
		}
		fmt.Println(string(data))

		information := *new(Information)
		json.Unmarshal(data, &information)
		token := information.Token
		message := information.Message

		switch information.Event {
		case NOTICE_EVENT:
			socket, ok := TPool.Workers[token]
			if ok {
				socket.Write([]byte(message))
			}
		case REGISTER_EVENT:
			WPool.Workers[token] = ws
			socket, ok := TPool.Workers[token]
			if ok {
				socket.Write([]byte(message))
			}
		case LOGOUT_EVENT:
			ws.Close()
			delete(WPool.Workers, token)
			return
		}
	}
}
