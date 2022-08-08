package imagehelper

import (
	"image"
	"image/color"
)

func RgbImageToFloat(imgPtr *image.Image) *[]float32 {
	img := *imgPtr

	size := img.Bounds().Size()

	var data = make([]float32, 0)

	for y := 0; y < size.Y; y++ {

		for x := 0; x < size.X; x++ {
			col := img.At(x, y)

			r, g, b, _ := col.RGBA()

			data = append(data, float32(b)/65535.0, float32(g)/65535.0, float32(r)/65535.0)
		}
	}

	return &data
}

func RgbImageToByte(imgPtr *image.Image) *[]uint8 {
	img := *imgPtr

	size := img.Bounds().Size()

	var data = make([]uint8, 0)

	for y := 0; y < size.Y; y++ {

		for x := 0; x < size.X; x++ {
			col := img.At(x, y)

			r, g, b, _ := col.RGBA()

			data = append(data, uint8(b/257), uint8(g/257), uint8(r/257))
		}
	}

	return &data
}

func ByteToRgbImage(dataPtr *[]uint8, width, height int) *image.RGBA {
	data := *dataPtr

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	size := img.Bounds().Size()

	for y := 0; y < size.Y; y++ {

		for x := 0; x < size.X; x++ {
			index := 3 * (y*width + x)

			col := color.RGBA{
				B: data[index],
				G: data[index+1],
				R: data[index+2],
				A: 255}

			img.Set(x, y, col)
		}
	}

	return img
}

func FloatToRgbImage(dataPtr *[]float32, width, height int) *image.RGBA {
	data := *dataPtr

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	size := img.Bounds().Size()

	for y := 0; y < size.Y; y++ {

		for x := 0; x < size.X; x++ {
			index := 3 * (y*width + x)

			img.Set(x, y, color.RGBA{byte(data[index+2] * 255),
				byte(data[index+1] * 255),
				byte(data[index] * 255), 255})
		}
	}

	return img
}

func FloatToGwImage(dataPtr *[]float32, width, height int) *image.Gray {
	data := *dataPtr

	img := image.NewGray(image.Rect(0, 0, width, height))

	size := img.Bounds().Size()

	for y := 0; y < size.Y; y++ {

		for x := 0; x < size.X; x++ {
			index := y*width + x

			img.Set(x, y, color.Gray{byte(data[index] * 255.0)})
		}
	}

	return img
}
