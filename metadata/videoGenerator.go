package metadata

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"cloud.google.com/go/storage"
	log "github.com/sirupsen/logrus"
)

//RELEASE
const inBackgroundVideoPath = "/tmp/background.mp4"
const inSoundAudioPath = "/tmp/track.wav";

const outOverlayedImgPath = "/tmp/overlayed.png"
const outOverlayedNoBackgroundImgPath = "/tmp/overlayed-no-background.png"
const outTempVideoPath = "/tmp/video-no-sound.mp4"
const outFinalVideoPath = "/tmp/video-with-sound.mp4"

// LOCAL RUN
// const inBackgroundVideoPath = "./cmd/in/background.mp4"
// const inSoundAudioPath = "./cmd/in/track.wav";

// const outOverlayedImgPath = "./cmd/out/overlayed.png"
// const outOverlayedNoBackgroundImgPath = "./cmd/out/overlayed-no-background.png"
// const outTempVideoPath = "./cmd/out/video-no-sound.mp4"
// const outFinalVideoPath = "./cmd/out/video-with-sound.mp4"

//DEBUG
// const inBackgroundVideoPath = "./in/background.mp4"
// const inSoundAudioPath = "./in/track.wav";

// const outOverlayedImgPath = "./out/overlayed.png"
// const outOverlayedNoBackgroundImgPath = "./out/overlayed-no-background.png"
// const outTempVideoPath = "./out/video-no-sound.mp4"
// const outFinalVideoPath = "./out/video-with-sound.mp4"

func GenerateAndSaveVideo(genes []string) {
	backgroundVideoUrl := fmt.Sprintf("%s/backgrounds-video/%s.mp4", BUCKET_BASE_PATH, genes[0])
	trackUrl := fmt.Sprintf("%s/tracks/%s.wav", BUCKET_BASE_PATH, genes[len(genes) - 1])
	readAndSaveMedia(backgroundVideoUrl, inBackgroundVideoPath)
	log.Println("Loaded trippy background video in memory!")
	readAndSaveMedia(trackUrl, inSoundAudioPath)
	log.Println("Loaded dope track in memory!")
	addBackground()
	log.Println("Added trippy background")
	addAudio()
	log.Println("Added banger track")
	videoNameBuilder := strings.Builder{}
	for _, gene := range genes {
		videoNameBuilder.WriteString(gene)
	}

	videoNameBuilder.WriteString(".mp4")

	saveVideoToGCloud(outFinalVideoPath, videoNameBuilder.String())
	log.Println("Saved final video to bucket!")
}

func readAndSaveMedia(mediaUrl string, savePath string) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)

	if err != nil {
		log.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	bucket := client.Bucket(GCLOUD_SOURCE_BUCKET_NAME)

	mediaReader, err := bucket.Object(mediaUrl).NewReader(ctx)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	defer mediaReader.Close()

	dst, err := os.Create(savePath)
	if err != nil {
		log.Fatal(err)
	}

	io.Copy(dst, mediaReader)
	if err != nil {
		log.Fatal(err)
	}
}

func overlayTraits(traitPaths []string) {
	// ffmpeg -y -i ./in/background.mp4 -vf "select=eq(n\,1)" -vframes 1 background.png
	// ffmpeg -i ./in/blue-fur.png -i ./in/shark-teeth-chain.png -i ./in/black-gas-mask.png -i ./in/cat-eyes.png -i ./in/daimunds.png -filter_complex "[0][1]overlay[bg0];[bg0][2]overlay[bg1];[bg1][3]overlay[bg2];[bg2][4]overlay[v]" -map "[v]" ./out/combined.png
	// ffmpeg -i ./in/woman.png -i ./in/shark-teeth-chain.png -i ./in/gas-mask.png -i ./in/glasses.png -i ./in/eth.png -filter_complex "[0][1]overlay[bg0];[bg0][2]overlay[bg1];[bg1][3]overlay[bg2];[bg2][4]overlay[v]" -map "[v]" ./out/combined2.png
	
	// // First time we include the background frame to have the full image
	overlayParams := []string{}
	overlayParams = append(overlayParams, "-y")
	// overlayParams = append(overlayParams, "-i", outBackgroundImagePath)
	for i := 0; i < len(traitPaths); i++ {
		overlayParams = append(overlayParams, "-i")
		overlayParams = append(overlayParams, traitPaths[i])
	}
	overlayParams = append(overlayParams, "-filter_complex", "[0][1]overlay[bg0];[bg0][2]overlay[bg1];[bg1][3]overlay[bg2];[bg2][4]overlay[bg3];[bg3][5]overlay[v]", "-map", "[v]", outOverlayedImgPath)
	overlayCommand := exec.Command("ffmpeg", overlayParams...)

	var out bytes.Buffer
	var stderr bytes.Buffer
	overlayCommand.Stdout = &out
	overlayCommand.Stderr = &stderr
	err := overlayCommand.Run()
	if err != nil {
    fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
    return
	}

	// Second time we exclude the background in order to overlay the real video
	overlayParams = []string{}
	overlayParams = append(overlayParams, "-y")
	for i := 1; i < len(traitPaths); i++ {
		overlayParams = append(overlayParams, "-i")
		overlayParams = append(overlayParams, traitPaths[i])
	}
	overlayParams = append(overlayParams, "-filter_complex", "[0][1]overlay[bg0];[bg0][2]overlay[bg1];[bg1][3]overlay[bg2];[bg2][4]overlay[v]", "-map", "[v]", outOverlayedNoBackgroundImgPath)
	overlayCommand = exec.Command("ffmpeg", overlayParams...)

	overlayCommand.Stdout = &out
	overlayCommand.Stderr = &stderr
	err = overlayCommand.Run()
	if err != nil {
    fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
    return
	}
}

func addBackground() {
	// ffmpeg -y -i ./in/background.mp4 -framerate 30 -i ./out/combined.png -filter_complex [0]overlay=x=0:y=0[out] -map [out] -map 0:a? -tag:v hvc1 -vcodec libx265 -pix_fmt yuv420p ./out/video-no-sound-compressed.mp4
	toVideo := exec.Command("ffmpeg", "-y",
	"-i", inBackgroundVideoPath,
	"-framerate", "30",
	"-i", outOverlayedNoBackgroundImgPath, 
	"-filter_complex", "[0]overlay=x=0:y=0[out]",
	"-map", "[out]",
	"-map", "0:a?",
	"-tag:v", "hvc1",
	"-vcodec", "libx265",
	"-pix_fmt", "yuv420p",
	outTempVideoPath)

	var out bytes.Buffer
	var stderr bytes.Buffer
	toVideo.Stdout = &out
	toVideo.Stderr = &stderr
	log.Println("Starting to generate dope video")
	err := toVideo.Run()
	if err != nil {
    fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
    return
	}
}

func addAudio() {
	// ffmpeg -y -i ./out/video-no-sound-compressed.mp4 -i ./in/track.wav -map 0:v -map 1:a -c:v copy -shortest ./out/final.mp4
	addAudio := exec.Command("ffmpeg", "-y",
	"-i", outTempVideoPath,
	"-i", inSoundAudioPath, 
	"-map", "0:v",
	"-map", "1:a",
	"-c:v", "copy",
	"-shortest",
	outFinalVideoPath)

	var out bytes.Buffer
	var stderr bytes.Buffer
	addAudio.Stdout = &out
	addAudio.Stderr = &stderr
	err := addAudio.Run()
	if err != nil {
    fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
    return
	}
}

func saveVideoToGCloud(filePath string, fileName string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer file.Close()

	log.Println("Opened file for upload")
	ctx := context.Background()
	client, err := storage.NewClient(ctx)

	if err != nil {
		log.Fatalf("storage.NewClient: %v", err)
	}
	defer client.Close()

	bucket := client.Bucket(GCLOUD_UPLOAD_BUCKET_NAME).Object(fileName).NewWriter(ctx)

	if _, err := io.Copy(bucket, file); err != nil {
		log.Fatalf("io.Copy: %v", err)
	}

	if err = bucket.Close(); err != nil {
		log.Errorf("Writer.Close: %v", err)
	}
}
