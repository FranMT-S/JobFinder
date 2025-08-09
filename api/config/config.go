package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	// API_PORT is the port the API will listen on
	API_PORT string

	// API_DEFAULT_PORT is the default port the API will listen on
	API_DEFAULT_PORT string = "8080"

	// API_HOST is the host the API will listen on
	API_HOST string

	// API_VERSION is the version of the API
	API_VERSION string

	// CLIENT_HOST is the host of the client
	CLIENT_HOST string

	ENVIRONMENT string
)

func setGlobalConfig() {
	API_PORT = os.Getenv("API_PORT")
	API_HOST = os.Getenv("API_HOST")
	API_VERSION = os.Getenv("API_VERSION")
	CLIENT_HOST = os.Getenv("CLIENT_HOST")

	// Set default values
	if API_PORT == "" {
		API_PORT = API_DEFAULT_PORT
	}

	if CLIENT_HOST == "" {
		// vue client
		CLIENT_HOST = "http://localhost:5173"
	}

	if API_HOST == "" {
		API_HOST = "localhost"
	}

	if API_VERSION == "" {
		API_VERSION = "v1"
	}

	ENVIRONMENT = os.Getenv("ENVIRONMENT")
}

func SetEnviromentConfig() error {
	if os.Getenv("ENVIRONMENT") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
			return err
		}
	}

	setGlobalConfig()
	return nil
}
