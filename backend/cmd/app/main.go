package main

import (
	"chat/config"
	"chat/internal/app"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic("failed to run app")
	}

	app.Run(cfg)
}
