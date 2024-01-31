package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/51f2d1cafc9c0f745c00037d

func solution(str, ending string) bool {
	return len(str) >= len(ending) && str[len(str)-len(ending):] == ending
}

func TestSolution(t *testing.T) {
	assert.True(t, solution("", ""))
	assert.True(t, solution(" ", ""))
	assert.True(t, solution("abc", "c"))
	assert.True(t, solution("banana", "ana"))
	assert.False(t, solution("a", "z"))
	assert.False(t, solution("", "t"))
	assert.False(t, solution("$a = $b + 1", "+1"))
	assert.True(t, solution("    ", "   "))
	assert.False(t, solution("1oo", "100"))
}
