package service

import (
	"image"
	"image/png"
	"os"
)

func RotateVertical(fileName string, img image.Image) {
	maxY := img.Bounds().Max.Y
	maxX := img.Bounds().Max.X

	bounds := img.Bounds()
	rotateImg := image.NewRGBA(bounds)

	for i := 0; i < maxY; i++ {
		for j := 0; j < maxX; j++ {
			oldColor := img.At(j, i)

			rotateImg.Set(maxX-1-j, maxY-1-i, oldColor)
		}
	}

	outFile, err := os.Create(fileName + "_rotate_verical.png")
	if err != nil {
		panic(err)
	}

	defer outFile.Close()

	err = png.Encode(outFile, rotateImg)
	if err != nil {
		panic(err)
	}
}

func RotateHorizontal(fileName string, img image.Image) {
	maxY := img.Bounds().Max.Y
	maxX := img.Bounds().Max.X

	bounds := img.Bounds()
	rotateImg := image.NewRGBA(bounds)

	for i := 0; i < maxY; i++ {
		for j := 0; j < maxX; j++ {
			oldColor := img.At(j, i)

			rotateImg.Set(maxX-1-j, i, oldColor)
		}
	}

	outFile, err := os.Create(fileName + "_rotate_horizontal.png")
	if err != nil {
		panic(err)
	}

	defer outFile.Close()

	err = png.Encode(outFile, rotateImg)
	if err != nil {
		panic(err)
	}
}

func GraySpectre(fileName string, img image.Image) {
	maxY := img.Bounds().Max.Y
	maxX := img.Bounds().Max.X

	bounds := img.Bounds()
	rotateImg := image.NewGray(bounds)

	for i := 0; i < maxY; i++ {
		for j := 0; j < maxX; j++ {
			oldColor := img.At(j, i)

			rotateImg.Set(j, i, oldColor)
		}
	}

	outFile, err := os.Create(fileName + "_gray.png")
	if err != nil {
		panic(err)
	}

	defer outFile.Close()

	err = png.Encode(outFile, rotateImg)
	if err != nil {
		panic(err)
	}
}
