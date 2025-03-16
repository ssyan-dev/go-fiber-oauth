package main

import (
	"log"

	"github.com/ssyan-dev/go-fiber-oauth/internal/application/config"
	"github.com/ssyan-dev/go-fiber-oauth/internal/application/server"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("%v", err)
	}

	server := server.NewServer(cfg)
	server.Run()
}
