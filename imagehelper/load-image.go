package imagehelper

import (
	"image"
	"os"
)

func LoadImageFromPath(path string) (*image.Image, error) {
	_, err := os.Stat(path)
	if err != nil {
		panic("path could not be found")
	}

	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer f.Close()

	return LoadImage(f)
}

func LoadImage(f *os.File) (*image.Image, error) {
	img, _, imageError := image.Decode(f)

	if imageError != nil {
		return nil, imageError
	}

	return &img, nil
}
