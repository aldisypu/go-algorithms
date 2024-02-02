package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/514b92a657cdc65150000006

func Multiple3And5(number int) int {
	sum := 0

	for i := 3; i < number; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	return sum
}

func TestMultiple3And5(t *testing.T) {
	assert.Equal(t, 23, Multiple3And5(10))
}
