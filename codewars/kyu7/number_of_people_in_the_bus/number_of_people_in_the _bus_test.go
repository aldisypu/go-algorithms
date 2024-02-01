package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Number(busStops [][2]int) int {
	passengers := 0

	for _, stop := range busStops {
		passengers += stop[0] - stop[1]
	}

	return passengers
}

func TestNumber(t *testing.T) {
	assert.Equal(t, 5, Number([][2]int{{10, 0}, {3, 5}, {5, 8}}))
	assert.Equal(t, 17, Number([][2]int{{3, 0}, {9, 1}, {4, 10}, {12, 2}, {6, 1}, {7, 10}}))
	assert.Equal(t, 21, Number([][2]int{{3, 0}, {9, 1}, {4, 8}, {12, 2}, {6, 1}, {7, 8}}))
	assert.Equal(t, 0, Number([][2]int{{0, 0}}))
}
