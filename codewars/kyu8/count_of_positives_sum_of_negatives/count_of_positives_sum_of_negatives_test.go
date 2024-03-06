package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/576bb71bbbcf0951d5000044

func CountPositivesSumNegatives(numbers []int) []int {
	countPositives := 0
	sumNegatives := 0

	for _, num := range numbers {
		if num > 0 {
			countPositives++
		} else {
			sumNegatives += num
		}
	}

	return []int{countPositives, sumNegatives}
}

func TestCountPositivesSumNegatives(t *testing.T) {
	assert.Equal(t, []int{10, -65}, CountPositivesSumNegatives([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15}))
}
