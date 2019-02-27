package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"path/filepath"
)

const (
	width  = 1980
	height = 1080
	n      = 5

	dir = "image"
)

var (
	path = filepath.Join(dir, "%02d.png")
)

func init() {
	rand.Seed(0)
}

func main() {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, 0766)
	}

	encoder := png.Encoder{CompressionLevel: png.NoCompression}
	for i := 0; i < n; i++ {
		dstFile, err := os.Create(fmt.Sprintf(path, i))
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		img := image.NewNRGBA(image.Rect(0, 0, width, height))

		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				c := color.RGBA{
					uint8(rand.Intn(256)),
					uint8(rand.Intn(256)),
					uint8(rand.Intn(256)),
					255,
				}
				img.Set(x, y, c)
			}
		}

		encoder.Encode(dstFile, img)

		dstFile.Close()
	}
}
