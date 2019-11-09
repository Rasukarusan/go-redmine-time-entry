package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	UserId   string `json:"user_id"`
	Password string `json:"password"`
}

func Load() *Config {
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
