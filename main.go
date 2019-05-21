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

	src.Dispatch(config)
}

func getConf() (conf *src.ServerConfig, err error) {
	content, err := ioutil.ReadFile("ini.json")
	if err != nil {
		return nil, errors.New("can't open the config file")
	}

	var config = new(src.ServerConfig)
	json.Unmarshal(content, &config)
	if err != nil {
		return nil, errors.New("explain json error")
	}

	return config, nil
}
