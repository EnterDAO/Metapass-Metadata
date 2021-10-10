package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"

	"github.com/fogleman/gg"
)

type Request struct {
	BgImgPath string
	FontPath  string
	FontSize  float64
	Text      string
	OutputPath string
}

func TextOnImg(request Request) (image.Image, error) {
	bgImage, err := gg.LoadImage(request.BgImgPath)
	if err != nil {
		return nil, err
	}
	imgWidth := bgImage.Bounds().Dx()
	imgHeight := bgImage.Bounds().Dy()

	dc := gg.NewContext(imgWidth, imgHeight)
	dc.DrawImage(bgImage, 0, 0)

	if err := dc.LoadFontFace(request.FontPath, request.FontSize); err != nil {
		return nil, err
	}
	//TODO: These shold be reconfigured
	x := float64(imgWidth - 140)
	y := float64(40)
	maxWidth := float64(imgWidth) - 60.0
	dc.SetColor(color.White)
	dc.DrawStringWrapped(request.Text, x, y, 0.5, 0.5, maxWidth, 1.5, gg.AlignCenter)
	dc.SavePNG(request.OutputPath)
	return dc.Image(), nil
}

func SaveImageToDest(outputPath string, targetImage image.Image) {
	f, err := os.Create(outputPath)
	if err != nil {
			panic(err)
	}
	defer f.Close()
	if err = jpeg.Encode(f, targetImage, nil); err != nil {
			log.Printf("failed to encode: %v", err)
	}

}
