encoder := png.Encoder{CompressionLevel: png.BestSpeed} 👈 独自エンコーダを定義
for _, path := range GetPathAll() {
	...
	img, err := png.Decode(srcFile)
	...
	pix := img.(*image.NRGBA).Pix
	for i := 0; i < len(pix); i += 4 {
		pix[i] = 255 - pix[i]
		pix[i+1] = 255 - pix[i+1]
		pix[i+2] = 255 - pix[i+2]
	}
	...
	encoder.Encode(dstFile, img) 👈 独自エンコーダを使用
	...
}
