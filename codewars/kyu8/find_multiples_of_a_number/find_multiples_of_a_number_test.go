package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/58ca658cc0d6401f2700045f

func FindMultiples(integer, limit int) []int {
	multiples := make([]int, 0)
	for i := integer; i <= limit; i += integer {
		multiples = append(multiples, i)
	}
	return multiples
}

func TestFindMultiples(t *testing.T) {
	assert.Equal(t, []int{5, 10, 15, 20, 25}, FindMultiples(5, 25))
	assert.Equal(t, []int{1, 2}, FindMultiples(1, 2))
}
