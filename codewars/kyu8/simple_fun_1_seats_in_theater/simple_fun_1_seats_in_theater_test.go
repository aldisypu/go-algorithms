package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/588417e576933b0ec9000045

func seatsInTheater(nCols int, nRows int, col int, row int) int {
	return (nCols - col + 1) * (nRows - row)
}

func TestSeatsInTheater(t *testing.T) {
	assert.Equal(t, 96, seatsInTheater(16, 11, 5, 3))
	assert.Equal(t, 0, seatsInTheater(1, 1, 1, 1))
	assert.Equal(t, 18, seatsInTheater(13, 6, 8, 3))
	assert.Equal(t, 99, seatsInTheater(60, 100, 60, 1))
	assert.Equal(t, 0, seatsInTheater(1000, 1000, 1000, 1000))
}
