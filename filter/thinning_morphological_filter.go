package filter

import (
	"errors"

	"github.com/simonmau/spacial-base-calculation/point"
)

type ThinningMorphological struct {
	width  int32
	height int32
	size   int32
}

func GenThinningMorphological(width, height int32) ThinningMorphological {
	return ThinningMorphological{
		width:  width,
		height: height,
		size:   width * height,
	}
}

func (t *ThinningMorphological) ApplyFilter(image *[]bool, resultImage *[]bool) error {
	img := *image
	resImg := *resultImage

	if len(img) != int(t.size) || len(resImg) != int(t.size) {
		return errors.New("wrong size for thinning-morphological")
	}

	for y := int32(1); y < t.height-1; y++ {
		for x := int32(1); x < t.width-1; x++ {
			pts := []point.T{
				{x - 1, y - 1},
				{x, y - 1},
				{x + 1, y - 1},

				{x - 1, y},
				{x, y},
				{x + 1, y},

				{x - 1, y + 1},
				{x, y + 1},
				{x + 1, y + 1},
			}

			hits := 0

			for _, pt := range pts {
				index := pt.ToIndex(&t.width, 1)

				if img[index] {
					hits++
				}
			}

			currentPosition := point.T{x, y}
			index := currentPosition.ToIndex(&t.width, 1)

			resImg[index] = hits >= 4
		}
	}

	return nil
}
