package main

import (
	"fmt"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)
func main() {
	// ffmpeg -i ./in/woman.png -i ./in/border.png  -filter_complex "[1][0:v] overlay" ./in/womanWithBorder.png
	womanPng := ffmpeg.Input("./in/woman.png").Get("0")
	borderPng := ffmpeg.Input("./in/border.png", ffmpeg.KwArgs{"filter_complex": fmt.Sprintf("[1][0:v] overlay")}).
	Output("./out/testWithCode.png").OverWriteOutput().ErrorToStdOut().Run()


	fmt.Println(borderPng)
}

// womanPng := ffmpeg.Input("./in/woman.png").Get("0");
// borderPng := ffmpeg.Input("./in/border.png").Get("1");
// result := ffmpeg_go.Filter([]*ffmpeg.Stream{
// 	womanPng,
// 	borderPng,
// }, "", ffmpeg_go.Args{} ,ffmpeg.KwArgs{"filter_complex": "[1][0:v] overlay"}).
// Output("./out/testWithCode.png").OverWriteOutput().ErrorToStdOut().Run()


// func main() {
// 	args := os.Args[1:]
// 	if len(args) > 0 {
// 		godotenv.Load(args[0])
// 	} else {
// 		godotenv.Load()
// 	}

// 	logger.SetupLogger()

// 	ethClient := dlt.ConnectToEthereum()

// 	port := os.Getenv("API_PORT")
// 	if envport := os.Getenv("PORT"); envport == "" {
// 		port = os.Getenv("API_PORT")
// 	}

// 	contractAddress := os.Getenv("CONTRACT_ADDRESS")
// 	if contractAddress == "" {
// 		log.Fatalln("Missing address. Populate in .env")
// 	}
// 	//Debug
// 	configService := config.NewConfigService("../config.json")

// 	//Local + Release
// 	// configService := config.NewConfigService("./config.json")

// 	funcframework.RegisterHTTPFunction("/token", handlers.HandleMetadataRequest(ethClient, contractAddress, configService))

// 	if err := funcframework.Start(port); err != nil {
// 		log.Fatalf("funcframework.Start: %v\n", err)
// 	}

// }
