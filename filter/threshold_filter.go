package filter

import (
	"errors"

	"github.com/simonmau/spacial-base-calculation/point"
	"github.com/ungerik/go3d/vec3"
)

type ThresholdFilter struct {
	width  int32
	height int32
	size   int32

	threshold float32
}

func GenThresholdFilter(threshold float32, width, height int32) ThresholdFilter {
	return ThresholdFilter{
		width:     width,
		height:    height,
		size:      width * height,
		threshold: threshold,
	}
}

func (t *ThresholdFilter) ApplyFilter(image, baseImage *[]float32, resultImage *[]bool) error {
	img := *image
	res := *resultImage
	base := *baseImage

	if len(img) != int(t.size*3) || len(base) != int(t.size*3) || len(res) != int(t.size) {
		return errors.New("wrong size for threshold")
	}

	for y := int32(0); y < t.height; y++ {
		for x := int32(0); x < t.width; x++ {
			imgPt := point.T{x, y}

			index := imgPt.ToIndex(&t.width, 1)

			index3 := imgPt.ToIndex(&t.width, 3)

			vec := vec3.T{
				img[index3] - base[index3],
				img[index3+1] - base[index3+1],
				img[index3+2] - base[index3+2],
			}

			res[index] = vec.LengthSqr() > t.threshold
		}
	}

	return nil
}
