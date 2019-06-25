for _, path := range GetPathAll() {
	...
	img, err := png.Decode(srcFile)
	...
	pix := img.(*image.NRGBA).Pix 👈 スライスとして抜き出す
	for i := 0; i < len(pix); i += 4 {
		pix[i] = 255 - pix[i]
		pix[i+1] = 255 - pix[i+1]
		pix[i+2] = 255 - pix[i+2]
	}
	...
	png.Encode(dstFile, img)
	...
}
