package main

import (
	"./src"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

func main() {
	config, err := getConf()
	if err != nil {
		fmt.Println(err)
		return
	}

	tcpPool := *new(src.TcpPool)
	websocketPool := *new(src.WebsocketPool)
	go src.TcpWorker(config, &tcpPool, &websocketPool)
	go src.WebWorker(config, &tcpPool, &websocketPool)
}

func getConf() (conf *src.ServerConfig, err error) {
	content, err := ioutil.ReadFile("ini.json")
	if err != nil {
		return _, errors.New("can't open the config file")
	}

	var config = new(src.ServerConfig)
	json.Unmarshal(content, &config)
	if err != nil {
		return _, errors.New("explain json error")
	}

	return config, nil
}
