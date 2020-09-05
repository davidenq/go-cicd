package main

import (
	"log"
	"os"

	"github.com/davidenq/go-cicd/cmd/gocicd/infra"
	"github.com/davidenq/go-cicd/cmd/gocicd/infra/config"
	"github.com/joho/godotenv"
)

func main() {

	//load environment variables from dot env file
	if os.Getenv("ENVIRONMENT") == "develop" {
		err := godotenv.Load("./cmd/gocicd/.env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	//load configuration
	config.Load()

	//start server
	infra.NewServer()
}
