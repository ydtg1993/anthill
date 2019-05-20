package src

import (
	"golang.org/x/net/websocket"
	"log"
	"net/http"
)

func (config *ServerConfig)WebWorker()  {
	http.Handle(config.Websocket.Pattern, websocket.Handler(Echo))
	if err := http.ListenAndServe(":"+config.Websocket.Port, nil); err != nil {
		log.Println(err)
	}
}

func Echo(ws *websocket.Conn) {
	for {
		var raply string
		if err := websocket.Message.Receive(ws, &raply); err != nil { //get infomation,write in adress
			log.Println("can't receive")
			break
		}
		msg := "Received:" + raply
		log.Println(msg)
		if err := websocket.Message.Send(ws, "come back infomation"); err != nil { //send infomation
			log.Println("can't send")
			break
		}
	}
}
