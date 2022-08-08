package filter

const (
	REDFAC   = 0.263
	GREENFAC = 0.655
	BLUEFAC  = 0.081
)

//2.9ms
func ConvertToGw(rgbImage *[]float32, result *[]float32, width, height int32) {
	colorImg := *rgbImage
	resImg := *result

	for y := int32(0); y < height; y++ {
		for x := int32(0); x < width; x++ {
			gwIndex := y*width + x

			colorIndex := gwIndex * 3

			resImg[gwIndex] = colorImg[colorIndex]*BLUEFAC + colorImg[colorIndex+1]*GREENFAC + colorImg[colorIndex+2]*REDFAC
		}
	}
}
