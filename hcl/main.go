package main

import (
	"log"

	"github.com/hashicorp/hcl/v2/hclsimple"
)

func main() {
	var config struct {
		Message string `hcl:"message"`
	}
	hclsimple.DecodeFile("config.hcl", nil, &config)
	log.Println(config.Message)
}
