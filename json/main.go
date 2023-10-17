package main

import (
	"encoding/json"
	"log"
	"os"
)

func main() {
	var config struct {
		Message string `json:"message"`
	}
	f, _ := os.Open("config.json")
	json.NewDecoder(f).Decode(&config)
	log.Println(config.Message)
}
