package main

import (
	"log"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"github.com/joho/godotenv"
	"github.com/metapass-metadata/config"
	"github.com/metapass-metadata/dlt"
	"github.com/metapass-metadata/handlers"
	"github.com/metapass-metadata/logger"
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
	// configService := config.NewConfigService("../config.json")
	// uniqueConfigService := config.NewConfigService("../config-unique.json")
	//Local
	configService := config.NewConfigService("./config.json")
	uniqueConfigService := config.NewUniqueConfigService("./config-unique.json")

	funcframework.RegisterHTTPFunction("/token", handlers.HandleMetadataRequest(ethClient, contractAddress, configService, uniqueConfigService))

	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}

}
