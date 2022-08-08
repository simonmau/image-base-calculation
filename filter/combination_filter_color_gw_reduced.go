package filter

import "errors"

//gets a FHD/rgb image
//colorImage is the resized (640x360) rgb image
//edgeImage is the sobel-gw-image
func CombinationFilterColorGwReduced(input *[]byte, colorImage *[]float32, edgeImage *[]float32) error {
	iRef := *input
	cRef := *colorImage

	fin := len(iRef)

	if fin != 1920*1080*3 {
		return errors.New("wrong image input size")
	}

	//to make this array constant doesnt improve performance
	var lookup [256]float32

	for i := 0; i <= 255; i++ {
		lookup[i] = float32(i) * _BYTE_INC_FAC
	}

	div := float32(1.0 / 9.0)

	ySmall := 0
	xSmall := 0

	for y := 1; y < 1080; y += 3 {
		xSmall = 0

		for x := 1; x < 1920; x += 3 {
			colIndex := (y*1920 + x) * 3

			valB := (lookup[iRef[colIndex-5760-3]] + lookup[iRef[colIndex-5760]] + lookup[iRef[colIndex-5760+3]] +
				lookup[iRef[colIndex-3]] + lookup[iRef[colIndex]] + lookup[iRef[colIndex+3]] +
				lookup[iRef[colIndex+5760-3]] + lookup[iRef[colIndex+5760]] + lookup[iRef[colIndex+5760+3]]) * div

			colIndex++

			valG := (lookup[iRef[colIndex-5760-3]] + lookup[iRef[colIndex-5760]] + lookup[iRef[colIndex-5760+3]] +
				lookup[iRef[colIndex-3]] + lookup[iRef[colIndex]] + lookup[iRef[colIndex+3]] +
				lookup[iRef[colIndex+5760-3]] + lookup[iRef[colIndex+5760]] + lookup[iRef[colIndex+5760+3]]) * div

			colIndex++

			valR := (lookup[iRef[colIndex-5760-3]] + lookup[iRef[colIndex-5760]] + lookup[iRef[colIndex-5760+3]] +
				lookup[iRef[colIndex-3]] + lookup[iRef[colIndex]] + lookup[iRef[colIndex+3]] +
				lookup[iRef[colIndex+5760-3]] + lookup[iRef[colIndex+5760]] + lookup[iRef[colIndex+5760+3]]) * div

			smallImgIndex := (ySmall*640 + xSmall) * 3

			cRef[smallImgIndex] = valB
			cRef[smallImgIndex+1] = valG
			cRef[smallImgIndex+2] = valR

			xSmall++
		}

		ySmall++
	}

	gwImg := make([]float32, 640*360)
	ConvertToGw(colorImage, &gwImg, 640, 360)
	Sobel(&gwImg, edgeImage, 640, 360)

	return nil
}
