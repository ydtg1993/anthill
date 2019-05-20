package main

import (
	"./src"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

func main() {
	config,err := getConf()
	if(err != nil){
		fmt.Println(err)
		return
	}

	go config.TcpWorker()
	go config.WebWorker()
}

func getConf() (conf *src.ServerConfig,err error) {
	content, err := ioutil.ReadFile("ini.json")
	if err != nil {
		return _,errors.New("can't open the config file")
	}

	var config = new(src.ServerConfig)
	json.Unmarshal(content, &config)
	if err != nil {
		return _,errors.New("explain json error")
	}

	return config,nil
}