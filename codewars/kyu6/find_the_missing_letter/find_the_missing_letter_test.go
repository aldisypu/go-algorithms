package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/5839edaa6754d6fec10000a2

func FindMissingLetter(chars []rune) rune {
	if chars[1]-chars[0] > 1 {
		return chars[0] + 1
	}

	return FindMissingLetter(chars[1:])
}

func TestFindMissingLetter(t *testing.T) {
	assert.Equal(t, 'e', FindMissingLetter([]rune{'a', 'b', 'c', 'd', 'f'}))
	assert.Equal(t, 'P', FindMissingLetter([]rune{'O', 'Q', 'R', 'S'}))
}
