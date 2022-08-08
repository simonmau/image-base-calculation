package regiongrowing

import (
	"testing"

	"github.com/simonmau/spacial-base-calculation/point"
	"github.com/stretchr/testify/assert"
)

func TestStackRegionGrowingFullHd(t *testing.T) {
	selectedArea := make([]bool, 1920*1080)
	for i := 0; i < 1920*1080; i++ {
		selectedArea[i] = true
	}

	service := GenStackRegionGrowing(1920, 1080, &[]point.T{{1, 1}})

	result, err := service.FindAllWithStartPoints(&selectedArea)

	assert.Nil(t, err)

	assert.Equal(t, 1, len(*result))

	foundArea := (*result)[0]
	assert.Equal(t, 1920*1080, len(foundArea))
}

func BenchmarkStackRegionGrowingFullHd(b *testing.B) {
	selectedArea := make([]bool, 1920*1080)
	for i := 0; i < 1920*1080; i++ {
		selectedArea[i] = true
	}

	service := GenStackRegionGrowing(1920, 1080, &[]point.T{{1920 / 2, 1080 / 2}})

	for i := 0; i < b.N; i++ {
		result, err := service.FindAllWithStartPoints(&selectedArea)
		assert.Nil(b, err)
		assert.Equal(b, 1, len(*result))
	}
}
