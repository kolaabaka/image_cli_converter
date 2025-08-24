package main

import (
	"bytes"
	"image/png"
	"image_cli_converter/configuration"
	"image_cli_converter/service"
	"os"
)

func main() {

	args := os.Args
	config := configuration.Config(args)

	byte_arr, err := os.ReadFile(config.FileName)
	if err != nil {
		panic("Coudn`t open file")
	}

	reader := bytes.NewReader(byte_arr)

	img, err := png.Decode(reader)

	if err != nil {
		panic(err)
	}

	if config.RotateHorizontal {
		service.RotateHorizontal(config.FileName, img)
	}

	if config.RotateVertical {
		service.RotateVertical(config.FileName, img)
	}

	if config.GraySpectre {
		service.GraySpectre(config.FileName, img)
	}
}
