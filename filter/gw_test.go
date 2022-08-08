package filter

import (
	"math/rand"
	"testing"
)

//2.9ms
func BenchmarkGwFilter(b *testing.B) {
	input := make([]float32, 1920*1080*3)

	for i := 0; i < len(input); i++ {
		input[i] = rand.Float32()
	}

	output := make([]float32, 1920*1080*3)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ConvertToGw(&input, &output, 1920, 1080)
	}
}
