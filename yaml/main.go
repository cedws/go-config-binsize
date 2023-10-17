package main

import (
	"log"
	"os"

	yaml "gopkg.in/yaml.v3"
)

func main() {
	var config struct {
		Message string `yaml:"message"`
	}
	f, _ := os.Open("config.yaml")
	yaml.NewDecoder(f).Decode(&config)
	log.Println(config.Message)
}
