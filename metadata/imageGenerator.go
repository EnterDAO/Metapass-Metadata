package metadata

import (
	"context"
	"fmt"
	"image"
	"image/color"
	"net/http"
	"strings"

	"cloud.google.com/go/storage"
	"github.com/disintegration/imaging"
	log "github.com/sirupsen/logrus"
)

const IMG_SIZE = 2000
const GCLOUD_UPLOAD_BUCKET_NAME = "metapass-images"
const GCLOUD_SOURCE_BUCKET_NAME = "metapass-source-images"
const BUCKET_BASE_PATH = "./dev-new"

func resourceExists(imageURL string) bool {
	resp, err := http.Get(imageURL)
	if err != nil {
		log.Error(err)
	}

	return resp.StatusCode != 404
}

func combineRemoteImages(bucket *storage.BucketHandle, basePath string, overlayPaths ...string) *image.NRGBA {

	ctx := context.Background()

	baseReader, err := bucket.Object(basePath).NewReader(ctx)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}
	defer baseReader.Close()
	base, err := imaging.Decode(baseReader)
	if err != nil {
		log.Fatalf("failed to decode image: %v", err)
	}
	dst := imaging.New(IMG_SIZE, IMG_SIZE, color.NRGBA{0, 0, 0, 0})
	dst = imaging.Paste(dst, base, image.Pt(0, 0))

	for _, op := range overlayPaths {
		r, err := bucket.Object(op).NewReader(ctx)
		if err != nil {
			log.Fatalf("failed to open image: %v", err)
		}
		defer r.Close()
		o, err := imaging.Decode(r)
		if err != nil {
			log.Fatalf("failed to open image: %v", err)
		}
		dst = imaging.Overlay(dst, o, image.Pt(0, 0), 1)
	}
	return dst
}

func reverseGenesOrder(genes []string) []string {
	res := make([]string, 0, len(genes))
	for i := len(genes) - 1; i >= 0; i-- {
		res = append(res, genes[i])
	}
	return res
}

func saveToGCloud(i *image.NRGBA, name string, imageFormat imaging.Format) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)

	if err != nil {
		log.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	bucket := client.Bucket(GCLOUD_UPLOAD_BUCKET_NAME).Object(name).NewWriter(ctx)

	err = imaging.Encode(bucket, i, imageFormat)

	if err != nil {
		log.Errorf("Upload: %v", err)
	}

	if err = bucket.Close(); err != nil {
		log.Errorf("Writer.Close: %v", err)
	}

}

func GenerateAndSaveImage(genes []string) {
	traitPaths := make([]string, len(genes) - 1)

	for i, gene := range genes {
		attrBucketName := ""
		switch i {
		case 0:
			attrBucketName = "backgrounds-image"
		case 1:
			attrBucketName = "skins"
		case 2:
			attrBucketName = "necklaces"
		case 3:
			attrBucketName = "mouth"
		case 4:
			attrBucketName = "eyes"
		case 5:
			attrBucketName = "vortex"
		case 6:
			attrBucketName = "tracks"
			continue
		}

		traitPaths[i] = fmt.Sprintf("%s/%v/%s.png", BUCKET_BASE_PATH, attrBucketName, gene)
	}

	ctx := context.Background()
	client, err := storage.NewClient(ctx)

	if err != nil {
		log.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	bucket := client.Bucket(GCLOUD_SOURCE_BUCKET_NAME)

	image := combineRemoteImages(bucket, traitPaths[0], traitPaths[1:]...)

	b := strings.Builder{}

	for _, gene := range genes {
		b.WriteString(gene)
	}

	b.WriteString(".jpg") // Finish with jpg extension

	saveToGCloud(image, b.String(), imaging.JPEG)

	log.Println("Uploaded trippy image to bucket!")
}

func GenerateAndSaveImageForVideo(genes []string) {
	traitPaths := make([]string, len(genes) - 2)

	for i, gene := range genes {
		attrBucketName := ""
		switch i {
		case 0:
			attrBucketName = "backgrounds-image"
			continue
		case 1:
			attrBucketName = "skins"
		case 2:
			attrBucketName = "necklaces"
		case 3:
			attrBucketName = "mouth"
		case 4:
			attrBucketName = "eyes"
		case 5:
			attrBucketName = "vortex"
		case 6:
			attrBucketName = "tracks"
			continue
		}
		
		traitPaths[i - 1] = fmt.Sprintf("%s/%v/%s.png", BUCKET_BASE_PATH, attrBucketName, gene)
	}

	ctx := context.Background()
	client, err := storage.NewClient(ctx)

	if err != nil {
		log.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	bucket := client.Bucket(GCLOUD_SOURCE_BUCKET_NAME)

	image := combineRemoteImages(bucket, traitPaths[0], traitPaths[1:]...)

	b := strings.Builder{}

	for _, gene := range genes {
		b.WriteString(gene)
	}

	b.WriteString("-for-video.png") // Finish with jpg extension

	saveToGCloud(image, b.String(), imaging.PNG)

	log.Println("Uploaded trippy image for video to bucket!")
}
