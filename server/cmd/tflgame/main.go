package main

import (
	"os"

	"tflgame/server"
	"tflgame/server/lib/config"
)

func main() {
	cfg := server.DefaultConfig()

	err := config.FromEnvironment(os.Getenv, &cfg)
	if err != nil {
		panic(err)
	}

	err = server.Run(cfg)
	if err != nil {
		panic(err)
	}
}
