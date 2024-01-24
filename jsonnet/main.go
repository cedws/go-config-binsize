package main

import (
	"encoding/json"
	"log"

	"github.com/google/go-jsonnet"
)

func main() {
	vm := jsonnet.MakeVM()
	result, _ := vm.EvaluateFile("config.jsonnet")
	var config struct {
		Message string `toml:"message"`
	}
	json.Unmarshal([]byte(result), &config)
	log.Println(config.Message)
}
