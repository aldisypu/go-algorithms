package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/565f5825379664a26b00007c

func GetSize(w, h, d int) [2]int {
	area := 2 * (w*h + h*d + d*w)
	volume := w * h * d
	return [2]int{area, volume}
}

func TestGetSize(t *testing.T) {
	assert.Equal(t, [2]int{88, 48}, GetSize(4, 2, 6))
	assert.Equal(t, [2]int{6, 1}, GetSize(1, 1, 1))
	assert.Equal(t, [2]int{10, 2}, GetSize(1, 2, 1))
	assert.Equal(t, [2]int{16, 4}, GetSize(1, 2, 2))
	assert.Equal(t, [2]int{600, 1000}, GetSize(10, 10, 10))
}
