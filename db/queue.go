package db

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

type VideoMetadata struct {
	TokenId string `json:"tokenid"`
	TrackUrl string `json:"trackurl"`
	BackgroundUrl string `json:"backgroundurl"`
	ImageUrl string `json:"imageurl"`
	FinalVideoName string `json:"finalvideoname"`
	IsGenerated bool `json:"isgenerated"`
	IsGenerating bool `json:"isgenerating"`
	HasErrored bool `json:"haserrored"`
}


func CheckVideoIsQueuedForGeneration(dbName string, queueCollection string, tokenId string) bool {
	collection, err := GetMongoDbCollection(dbName, queueCollection)
	if err != nil {
		log.Error(err)
	}

	defer disconnectDB()

	res := collection.FindOne(context.Background(), bson.M{"tokenid": tokenId})

	return res.Err() == nil
}

func AddVideoForGeneration(model VideoMetadata) {
	apiUrl := os.Getenv("QUEUE_BASE_URL")
	endpoint := fmt.Sprintf("%s/add-video", apiUrl)

	jsonModel, err := json.Marshal(model)
	if err != nil {
		log.Error(err.Error())
	}
	
	_, err = http.Post(endpoint, "application/json", bytes.NewBuffer(jsonModel))

	if err != nil {
		log.Error(err.Error())
	}
	log.Println("Queued video for generation")
}