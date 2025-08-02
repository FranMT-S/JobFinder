package main

import (
	"log"

	"github.com/FranMT-S/JobFinder/config"
	"github.com/FranMT-S/JobFinder/server"
)

func main() {
	config.SetEnviromentConfig()
	server := server.NewServer(config.API_PORT)

	err := server.Start()
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
