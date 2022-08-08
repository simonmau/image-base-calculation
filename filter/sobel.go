package filter

//20ms
func Sobel(source, result *[]float32, width, height int32) {
	img := *source
	rImg := *result

	upperBoundWidth := width - 1
	upperBoundHeight := height - 1

	for y := int32(1); y < upperBoundHeight; y++ {
		for x := int32(1); x < upperBoundWidth; x++ {
			gwIndex := y*width + x

			fx := img[gwIndex+width-1] + 2.0*img[gwIndex+width] + img[gwIndex+width+1] -
				img[gwIndex-width-1] - 2.0*img[gwIndex-width] - img[gwIndex-width+1]

			fy := img[gwIndex-width-1] + 2.0*img[gwIndex-1] + img[gwIndex+width-1] -
				img[gwIndex-width+1] - 2.0*img[gwIndex+1] - img[gwIndex+width+1]

			//64-bit sqrt is faster than the 32-bit version
			// value := math.Sqrt(float32(fx*fx + fy*fy))

			value := (fx + fy) * 0.5

			if value > 1.0 {
				value = 1.0
			}

			if value < 0.0 {
				value = 0.0
			}

			rImg[gwIndex] = value
			// rImg[gwIndex] = float32(value)
		}
	}
}

//20ms
func SobelRgb(source, result *[]float32, width, height int32) {
	img := *source
	rImg := *result

	w3 := width * 3

	upperBoundWidth := width - 1
	upperBoundHeight := height - 1

	for y := int32(1); y < upperBoundHeight; y++ {
		for x := int32(1); x < upperBoundWidth; x++ {
			for ch := int32(0); ch < 3; ch++ {
				gwIndex := (y*width+x)*3 + ch

				fx := img[gwIndex+w3-3] + 2.0*img[gwIndex+w3] + img[gwIndex+w3+3] -
					img[gwIndex-w3-3] - 2.0*img[gwIndex-w3] - img[gwIndex-w3+3]

				fy := img[gwIndex-w3-3] + 2.0*img[gwIndex-3] + img[gwIndex+w3-3] -
					img[gwIndex-w3+3] - 2.0*img[gwIndex+3] - img[gwIndex+w3+3]

				value := (fx + fy) * 0.5

				if value > 1.0 {
					value = 1.0
				}

				if value < 0.0 {
					value = 0.0
				}

				rImg[gwIndex] = value
			}
		}
	}
}
