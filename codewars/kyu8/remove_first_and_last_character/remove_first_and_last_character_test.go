package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/56bc28ad5bdaeb48760009b0

func RemoveChar(word string) string {
	return word[1 : len(word)-1]
}

func TestRemoveChar(t *testing.T) {
	assert.Equal(t, "loquen", RemoveChar("eloquent"))
	assert.Equal(t, "ountr", RemoveChar("country"))
	assert.Equal(t, "erso", RemoveChar("person"))
	assert.Equal(t, "lac", RemoveChar("place"))
}
