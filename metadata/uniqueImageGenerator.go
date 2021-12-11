package metadata

import (
	"context"
	"fmt"
	"strconv"

	"cloud.google.com/go/storage"
	"github.com/disintegration/imaging"
	log "github.com/sirupsen/logrus"
)

func GenerateAndSaveUniqueImage(geneUrl string, uniqueIndex int64) {
	traitPaths := make([]string, 2)

	traitPaths[0] = fmt.Sprintf("%s/%s/%d.png", BUCKET_BASE_PATH, uniqueBackgroundImagesFolder, int(uniqueIndex))
	traitPaths[1] = fmt.Sprintf("%s/%s/%d.png", BUCKET_BASE_PATH, uniqueImagesFolder, int(uniqueIndex))

	ctx := context.Background()
	client, err := storage.NewClient(ctx)

	if err != nil {
		log.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	bucket := client.Bucket(GCLOUD_SOURCE_BUCKET_NAME)

	image := combineRemoteImages(bucket, traitPaths[0], traitPaths[1:]...)


	saveToGCloud(image, fmt.Sprintf("%s%s", geneUrl, ".jpg"), imaging.JPEG)

	log.Println("Uploaded UNIQUE image to bucket!")
}

// This is not completely necessary. As we have preassembled image. Keeping it to keep both flows with consistent steps
func GenerateAndSaveUniqueImageForVideo(geneUrl string,  uniqueIndex int64) {
	uniqueImagePath := fmt.Sprintf("%s/%v/%s.png", BUCKET_BASE_PATH, uniqueImagesFolder, strconv.Itoa(int(uniqueIndex)))
	ctx := context.Background()
	client, err := storage.NewClient(ctx)

	if err != nil {
		log.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	bucket := client.Bucket(GCLOUD_SOURCE_BUCKET_NAME)

	image := combineRemoteImages(bucket, uniqueImagePath)

	saveToGCloud(image, geneUrl + "-for-video.png", imaging.PNG)

	log.Println("Uploaded UNIQUE image for video to bucket!")
}

