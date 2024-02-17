package kata

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/57eaeb9578748ff92a000009

func SumMix(arr []any) int {
	sum := 0

	for _, value := range arr {
		switch v := value.(type) {
		case int:
			sum += v
		case string:
			num, err := strconv.Atoi(v)
			if err == nil {
				sum += num
			}
		}
	}

	return sum
}

func TestSumMix(t *testing.T) {
	assert.Equal(t, 22, SumMix([]any{9, 3, "7", "3"}))
	assert.Equal(t, 42, SumMix([]any{"5", "0", 9, 3, 2, 1, "9", 6, 7}))
	assert.Equal(t, 41, SumMix([]any{"3", 6, 6, 0, "5", 8, 5, "6", 2, "0"}))
	assert.Equal(t, 45, SumMix([]any{"1", "5", "8", 8, 9, 9, 2, "3"}))
	assert.Equal(t, 61, SumMix([]any{8, 0, 0, 8, 5, 7, 2, 3, 7, 8, 6, 7}))
}
