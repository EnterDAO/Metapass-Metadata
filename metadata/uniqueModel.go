package metadata

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/metapass-metadata/config"
	"github.com/metapass-metadata/db"
)
const uniqueImagesFolder = "unique-images"
const uniqueBackgroundsFolder = "unique-backgrounds"
const uniqueBackgroundImagesFolder = "unique-background-images"
const uniqueTracksFolder = "unique-tracks"

func getLegendaryGeneAttribute(g string, configService *config.UniqueConfigService, uniqueIndex int64) StringAttribute {
	return StringAttribute{
		TraitType: "Legendary",
		Value:     configService.Name[uniqueIndex],
	}
}
// TODO: Check overlay order (https://enterdao.slack.com/archives/C02CD2N5TBK/p1638837850047000)
func (g *Genome) uniqueAttributes(uniqueConfigService *config.UniqueConfigService, uniqueIndex int64) []interface{} {
	gStr := string(*g)

	res := []interface{}{}
	res = append(res, getLegendaryGeneAttribute(gStr, uniqueConfigService, uniqueIndex))
	return res
}


func (g *Genome) UniqueMetadata(tokenId string, uniqueIndex int64, configService *config.ConfigService, uniqueConfigService *config.UniqueConfigService) Metadata {
	var m Metadata
	m.Attributes = g.uniqueAttributes(uniqueConfigService, uniqueIndex)
	m.Name = g.name(nil, tokenId)
	m.Description = g.description(nil, tokenId)
	m.ExternalUrl = fmt.Sprintf("%s%s", EXTERNAL_URL, tokenId)

	fullUrl := strings.Builder{}
	fileName := strings.Builder{}
	fullUrl.WriteString(METAPASS_IMAGE_URL) // Start with base url

	for i := 0; i < 10; i++ {
		fullUrl.WriteString(strconv.Itoa(int(uniqueIndex)))
		fileName.WriteString(strconv.Itoa(int(uniqueIndex)))
	}

	geneUrl := fullUrl.String()

	m.Image = geneUrl + ".jpg"
	m.AnimationUrl = geneUrl + ".mp4"

	forVideoImageUrl := fmt.Sprintf("%s-for-video.png", geneUrl)
	imageUrl := fmt.Sprintf("%s.jpg", geneUrl)

	dbName := os.Getenv("QUEUE_DB")
	queueCollection := os.Getenv("QUEUE_COLLECTION")
	
	videoImageExists  := resourceExists(forVideoImageUrl)
	imageExists := resourceExists(imageUrl)

	videoExists := db.CheckVideoIsQueuedForGeneration(dbName, queueCollection, tokenId)
	if !imageExists {
		GenerateAndSaveUniqueImage(fileName.String(), uniqueIndex)
	}
	if !videoImageExists {
		GenerateAndSaveUniqueImageForVideo(fileName.String(), uniqueIndex)		
	}

	if !videoExists {
		videoUrl := fmt.Sprintf("%s/%s/%d.mp4", BUCKET_BASE_PATH, uniqueBackgroundsFolder, uniqueIndex)
		trackUrl := fmt.Sprintf("%s/%s/%d.wav", BUCKET_BASE_PATH, uniqueTracksFolder, uniqueIndex)
		forVideoUrl := fmt.Sprintf("%s-for-video.png", fileName.String())
		videoName := fmt.Sprintf("%s.mp4", fileName.String())
		db.AddVideoForGeneration(db.VideoMetadata{
			TokenId:        tokenId,
			TrackUrl:       trackUrl,
			BackgroundUrl:  videoUrl,
			ImageUrl:       forVideoUrl,
			FinalVideoName: videoName,
			IsGenerated:    false,
			IsGenerating:   false,
			HasErrored:     false,
		})
	}

	return m
}