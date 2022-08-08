package imagehelper

import (
	"image"
	"image/png"
	"os"
)

func SaveImageToPath(path string, img *image.Image) error {
	f, err := os.Create(path)

	if err != nil {
		return err
	}
	defer f.Close()

	return SaveImage(f, img)
}

// Encode to `PNG` with `DefaultCompression` level
// then save to file
func SaveImage(f *os.File, img *image.Image) error {
	return png.Encode(f, *img)
}

func SaveRgbaImageToPath(path string, img *image.RGBA) error {
	baseImg := img.SubImage(img.Rect)
	return SaveImageToPath(path, &baseImg)
}

func SaveGwImageToPath(path string, img *image.Gray) error {
	baseImg := img.SubImage(img.Rect)
	return SaveImageToPath(path, &baseImg)
}
