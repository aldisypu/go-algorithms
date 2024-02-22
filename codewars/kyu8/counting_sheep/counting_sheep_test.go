package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/54edbc7200b811e956000556

func CountSheeps(numbers []bool) int {
	count := 0
	for _, present := range numbers {
		if present {
			count++
		}
	}
	return count
}

func TestCountSheeps(t *testing.T) {
	assert.Equal(t, 17, CountSheeps([]bool{
		true, true, true, false,
		true, true, true, true,
		true, false, true, false,
		true, false, false, true,
		true, true, true, true,
		false, false, true, true,
	}))
}
