for _, path := range GetPathAll() {
	...
	img, err := png.Decode(srcFile)
	 ...
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
	...
	png.Encode(dstFile, img)
	...
}
