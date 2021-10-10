package metadata

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"cloud.google.com/go/storage"
	log "github.com/sirupsen/logrus"
)
const inWomanImgPath = "./in/woman.png"
const inBorderImgPath = "./in/border-wide.png"
const inBackgroundVideoPath = "./in/background.mp4"
const inSoundAudioPath = "./in/squidgame.mp3";

const outImagePath = "./out/womanWithBorder.png"
const outTempVideoPath = "./out/video-no-sound.mp4"
const outFinalVideoPath = "./out/final.mp4"

const fontPath = "./font/Poppins/Poppins-Regular.ttf"
const fontSize = 24

func generateAndSaveVideo(tokenId string, genes []string) {
	// TODO: Get border image based on gene
	addBorder()
	addTokenId(tokenId)
	addBackground()
	addAudio()

	videoNameBuilder := strings.Builder{}
	for _, gene := range genes {
		videoNameBuilder.WriteString(gene)
	}

	videoNameBuilder.WriteString(".mp4")

	saveVideoToGCloud(tokenId, videoNameBuilder.String())
}

func addBorder() {
		addBorderCommand := exec.Command("ffmpeg", "-y",
		"-i", inWomanImgPath,
		"-i", inBorderImgPath,
		"-filter_complex", "[1][0:v] overlay=80:0", 
		outImagePath)
 
	 err := addBorderCommand.Run()
	 if err != nil {
		 fmt.Println(err.Error())
	 }
}

func addTokenId(tokenId string) {
	params := DrawTokenIdParams{
		BgImgPath: outImagePath,
		FontPath:  fontPath,
		FontSize:  fontSize,
		Text:      fmt.Sprintf("Meta Pass #%s", tokenId),
		OutputPath: outImagePath,
	}

	_, err := DrawTextAndSavePNG(params)

	if err != nil {
		fmt.Println(err)
	}
}

func addBackground() {
	toVideo := exec.Command("ffmpeg", "-y",
	"-i", inBackgroundVideoPath,
	"-i", outImagePath, 
	"-filter_complex", "[0:v][1:v] overlay=0:0:enable='between(t,0,20)'",
	"-pix_fmt", "yuv420p",
	"-b:v", "2000k",
	outTempVideoPath)

	err := toVideo.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func addAudio() {
	addAudio := exec.Command("ffmpeg", "-y",
	"-i", outTempVideoPath,
	"-i", inSoundAudioPath, 
	"-map", "0:v",
	"-map", "1:a",
	"-b:v", "2000k",
	"-c:v", "copy",
	"-shortest",
	outFinalVideoPath)

	err := addAudio.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func saveVideoToGCloud(tokenId string, fileName string) {
	file, err := os.Open(outFinalVideoPath)
	if err != nil {
		fmt.Println(err)
	}
	ctx := context.Background()
	client, err := storage.NewClient(ctx)

	if err != nil {
		log.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	bucket := client.Bucket(GCLOUD_UPLOAD_BUCKET_NAME).Object(fileName).NewWriter(ctx)

	if _, err := io.Copy(bucket, file); err != nil {
		log.Errorf("io.Copy: %v", err)
	}

	if err != nil {
		log.Errorf("Upload: %v", err)
	}

	if err = bucket.Close(); err != nil {
		log.Errorf("Writer.Close: %v", err)
	}
}
