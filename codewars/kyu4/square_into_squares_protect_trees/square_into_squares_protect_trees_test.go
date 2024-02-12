package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/54eb33e5bc1a25440d000891

func Decompose(n int64) []int64 {
	result, success := decomposeHelper(n*n, n-1)
	if success {
		return result
	}
	return []int64{}
}

func decomposeHelper(remaining, current int64) ([]int64, bool) {
	if remaining == 0 {
		return []int64{}, true
	}

	for i := current; i > 0; i-- {
		if i*i <= remaining {
			if result, success := decomposeHelper(remaining-i*i, i-1); success {
				return append(result, i), true
			}
		}
	}

	return nil, false
}

func TestDecompose(t *testing.T) {
	assert.Equal(t, []int64{1, 3, 5, 8, 49}, Decompose(50))
	assert.Equal(t, []int64{3, 4}, Decompose(5))
	assert.Equal(t, []int64{}, Decompose(2))
	assert.Equal(t, []int64{2, 3, 6}, Decompose(7))
}
