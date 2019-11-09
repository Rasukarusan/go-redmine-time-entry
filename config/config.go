package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Url      string `json:"url"`
}

func Load() *Config {
	// @TODO 環境変数から取得する形式、もしくはrakyll/statikを使う形式にする
	jsonString, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		log.Println(err)
	}
	c := new(Config)
	err = json.Unmarshal(jsonString, c)
	if err != nil {
		log.Println(err)
	}
	return c
}
