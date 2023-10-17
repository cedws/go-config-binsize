package main

import (
	"log"

	"github.com/BurntSushi/toml"
)

func main() {
	var config struct {
		Message string `toml:"message"`
	}
	toml.DecodeFile("config.toml", &config)
	log.Println(config.Message)
}
