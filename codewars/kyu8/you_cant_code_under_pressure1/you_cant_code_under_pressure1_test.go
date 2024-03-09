package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/53ee5429ba190077850011d4

func DoubleInteger(i int) int {
	return i * 2
}

func TestDoubleInteger(t *testing.T) {
	assert.Equal(t, 2, DoubleInteger(1))
}
