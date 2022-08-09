package filter

import (
	"fmt"
	"testing"

	"github.com/simonmau/spacial-base-calculation/point"
	"github.com/stretchr/testify/assert"
)

func TestResizeByteImageRgb(t *testing.T) {
	sourceWidth := int32(100)
	sourceHeight := int32(100)
	sourceImage := make([]byte, sourceWidth*sourceHeight*3)

	for y := int32(0); y < sourceHeight; y++ {
		for x := int32(0); x < sourceWidth; x++ {
			pt := point.T{x, y}
			index := pt.ToIndex(&sourceWidth, 3)

			sourceImage[index] = 255
			sourceImage[index+1] = 255
			sourceImage[index+2] = 255
		}
	}

	outputWidth := int32(10)
	outputHeight := int32(10)
	outputImage := make([]byte, outputWidth*outputHeight*3)

	ResizeByteImageRgb(&sourceImage, sourceWidth, sourceHeight, 10, 10, outputWidth, outputHeight, &outputImage)

	for y := int32(0); y < outputHeight; y++ {
		for x := int32(0); x < outputWidth; x++ {
			pt := point.T{x, y}
			index := pt.ToIndex(&outputWidth, 3)

			assert.Equal(t, byte(255), outputImage[index], fmt.Sprintf("error with x %v, y %v", x, y))
			assert.Equal(t, byte(255), outputImage[index+1], fmt.Sprintf("error + 1 with x %v, y %v", x, y))
			assert.Equal(t, byte(255), outputImage[index+2], fmt.Sprintf("error + 2 with x %v, y %v", x, y))
		}
	}
}
