package filter

//takes about 2ms for FHD, 3CHAN
func ByteToFloatFilter(input *[]byte, output *[]float32) {
	iRef := *input
	oRef := *output

	fin := len(iRef)

	//to make this array constant doesnt improve performance
	var lookup [256]float32

	for i := 0; i <= 255; i++ {
		lookup[i] = float32(i) * _BYTE_INC_FAC
	}

	//this is the bottleneck
	for i := 0; i < fin; i++ {
		oRef[i] = lookup[iRef[i]]
	}
}
