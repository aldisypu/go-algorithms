package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/5168bb5dfe9a00b126000018

func Solution(word string) string {
	if len(word) <= 1 {
		return word
	}
	return Solution(word[1:]) + string(word[0])
}

func TestSolution(t *testing.T) {
	assert.Equal(t, "dlrow", Solution("world"))
}
