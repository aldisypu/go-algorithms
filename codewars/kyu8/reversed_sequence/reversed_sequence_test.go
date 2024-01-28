package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/5a00e05cc374cb34d100000d

func ReverseSeq(n int) []int {
	if n == 0 {
		return []int{}
	}

	return append([]int{n}, ReverseSeq(n-1)...)
}

func TestReverseSeq(t *testing.T) {
	assert.Equal(t, []int{5, 4, 3, 2, 1}, ReverseSeq(5))
}
