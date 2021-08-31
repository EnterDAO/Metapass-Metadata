package main

import (
	"log"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"github.com/joho/godotenv"
	"github.com/lobster-metadata/config"
	"github.com/lobster-metadata/dlt"
	"github.com/lobster-metadata/handlers"
	"github.com/lobster-metadata/logger"
)

func main() {
	args := os.Args[1:]
	if len(args) > 0 {
		godotenv.Load(args[0])
	} else {
		godotenv.Load()
	}

	logger.SetupLogger()

	ethClient := dlt.ConnectToEthereum()

	port := os.Getenv("API_PORT")
	if envport := os.Getenv("PORT"); envport == "" {
		port = os.Getenv("API_PORT")
	}

	contractAddress := os.Getenv("CONTRACT_ADDRESS")
	if contractAddress == "" {
		log.Fatalln("Missing address. Populate in .env")
	}
	//Debug
	configService := config.NewConfigService("../config.json")

	//Local + Release
	// configService := config.NewConfigService("./config.json")

	funcframework.RegisterHTTPFunction("/token", handlers.HandleMetadataRequest(ethClient, contractAddress, configService))

	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}

}
