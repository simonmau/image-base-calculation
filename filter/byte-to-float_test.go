package filter

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestByteToFloatFilter(t *testing.T) {
	input := make([]byte, 4)

	input[0] = 0
	input[1] = 255
	input[2] = 127
	input[3] = 50

	output := make([]float32, 4)

	ByteToFloatFilter(&input, &output)

	assert.InDelta(t, float32(0.0), output[0], 0.00194)
	assert.InDelta(t, float32(1.0), output[1], 0.00194)
	assert.InDelta(t, float32(0.498), output[2], 0.00194)
	assert.InDelta(t, float32(0.196), output[3], 0.00194)
}

func BenchmarkByteToFloatFilter(b *testing.B) {
	input := make([]byte, 1920*1080*3)

	for i := 0; i < len(input); i++ {
		input[i] = byte(rand.Intn(255))
	}

	output := make([]float32, 1920*1080*3)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ByteToFloatFilter(&input, &output)
	}
}
