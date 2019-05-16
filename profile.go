package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/pkg/profile"
)

const (
	dir    = "image"
	dstDir = "dst"
)

var (
	path = filepath.Join(dir, "*.png")
)

type Processor struct {
	Encoder Encoder
}

func (processor Processor) ProcessAll() {
	for _, p := range GetPathAll() {
		srcFile, err := os.Open(p)
		if err != nil {
			fmt.Fprintln(os.Stderr, err, p)
			continue
		}

		srcImage, err := png.Decode(srcFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		ProcessSlice(srcImage)

		dstFile, err := os.Create(filepath.Join(dstDir, filepath.Base(p)))
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		processor.Encoder.Encode(dstFile, srcImage)

		dstFile.Close()

		srcFile.Close()
	}

}

func ProcessSlice(img image.Image) {
	pix := img.(*image.NRGBA).Pix
	for i := 0; i < len(pix); i += 4 {
		pix[i] = 255 - pix[i]
		pix[i+1] = 255 - pix[i+1]
		pix[i+2] = 255 - pix[i+2]
	}
}

func ProcessImage(img image.Image) {
	imgNRGBA := img.(*image.NRGBA)
	bounds := imgNRGBA.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := imgNRGBA.At(x, y).(color.NRGBA)
			imgNRGBA.Set(x, y, color.RGBA{
				uint8(255 - c.R),
				uint8(255 - c.G),
				uint8(255 - c.B),
				c.A,
			})
		}
	}
}

type Encoder interface {
	Encode(w io.Writer, m image.Image) error
}

type DefaultEncoder struct{}

func (de *DefaultEncoder) Encode(w io.Writer, m image.Image) error {
	return png.Encode(w, m)
}

func EncodeDefault() {
	processor := Processor{Encoder: &DefaultEncoder{}}
	processor.ProcessAll()
}

func EncodeNoCompression() {
	processor := Processor{Encoder: &png.Encoder{CompressionLevel: png.NoCompression}}
	processor.ProcessAll()
}

func EncodeDefaultCompression() {
	processor := Processor{Encoder: &png.Encoder{CompressionLevel: png.DefaultCompression}}
	processor.ProcessAll()
}

func EncodeBestSpeed() {
	processor := Processor{Encoder: &png.Encoder{CompressionLevel: png.BestSpeed}}
	processor.ProcessAll()
}

func BestCompression() {
	processor := Processor{Encoder: &png.Encoder{CompressionLevel: png.BestCompression}}
	processor.ProcessAll()
}

func GetPathAll() []string {
	paths, _ := filepath.Glob(path)
	return paths
}

func main() {
	defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()

	if err := os.Mkdir(dstDir, 0766); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer func() {
		if err := os.RemoveAll(dstDir); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
	}()

	switch os.Args[1] {
	case "Default":
		EncodeDefault()
	case "NoCompression":
		EncodeNoCompression()
	case "DefaultCompression":
		EncodeDefaultCompression()
	case "BestSpeed":
		EncodeBestSpeed()
	case "BestCompression":
		BestCompression()
	default:
		log.Println("no operation")
		return
	}

}
