package filter

import (
	"errors"

	"github.com/simonmau/spacial-base-calculation/point"
)

func ResizeByteImageRgb(input *[]byte, oldWidth, oldHeight, startX, startY, newWidth, newHeight int32, output *[]byte) error {
	if newWidth+startX >= oldWidth {
		return errors.New("cannot resize to bigger image (WIDTH)")
	}

	if newHeight+startY >= oldHeight {
		return errors.New("cannot resize to bigger image (HEIGHT)")
	}

	iRef := *input
	oRef := *output

	if len(iRef) != int(oldWidth*oldHeight*3) || len(oRef) != int(newWidth*newHeight*3) {
		return errors.New("input or output array do not have the right dimensions (ResizeByte)")
	}

	for y := int32(0); y < newHeight; y++ {
		for x := int32(0); x < newWidth; x++ {
			pt := point.T{x, y}
			newImageIndex := pt.ToIndex(&newWidth, 3)

			pt = point.T{x + startX, y + startY}
			oldImageIndex := pt.ToIndex(&oldWidth, 3)

			oRef[newImageIndex] = iRef[oldImageIndex]
			oRef[newImageIndex+1] = iRef[oldImageIndex+1]
			oRef[newImageIndex+2] = iRef[oldImageIndex+2]
		}
	}

	return nil
}
