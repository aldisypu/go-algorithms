package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/58261acb22be6e2ed800003a

func GetVolumeOfCuboid(length, width, height float64) float64 {
	return length * width * height
}

func TestGetVolumeOfCuboid(t *testing.T) {
	assert.Equal(t, 4.0, GetVolumeOfCuboid(1.0, 2.0, 2.0))
	assert.Equal(t, 63.0, GetVolumeOfCuboid(6.3, 2.0, 5.0))
}
