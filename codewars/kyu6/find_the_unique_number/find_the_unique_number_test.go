package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/585d7d5adb20cf33cb000235

func FindUniq(arr []float32) float32 {
	countMap := make(map[float32]int)

	for _, num := range arr {
		countMap[num]++
	}

	for num, count := range countMap {
		if count == 1 {
			return num
		}
	}

	return 0
}

func TestFindUniq(t *testing.T) {
	assert.Equal(t, float32(2), FindUniq([]float32{1.0, 1.0, 1.0, 2.0, 1.0, 1.0}))
	assert.Equal(t, float32(0.55), FindUniq([]float32{0, 0, 0.55, 0, 0}))
}
