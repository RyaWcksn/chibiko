package main

import (
	"github.com/RyaWcksn/chibiko/configs"
	"github.com/RyaWcksn/chibiko/server"
)

func main() {
	cfg, err := configs.ReadFromFile("./configs/config.json")
	if err != nil {
		panic(err)
	}
	sv := server.New(cfg)
	sv.Start()
}
