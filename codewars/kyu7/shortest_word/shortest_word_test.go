package kata

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/57cebe1dc6fdc20c57000ac9

func FindShort(s string) int {
	words := strings.Fields(s)

	shortestLength := len(s)

	for _, word := range words {
		wordLength := len(word)
		if wordLength < shortestLength {
			shortestLength = wordLength
		}
	}

	return shortestLength

}

func TestFindShort(t *testing.T) {
	assert.Equal(t, 3, FindShort("bitcoin take over the world maybe who knows perhaps"))
	assert.Equal(t, 3, FindShort("turns out random test cases are easier than writing out basic ones"))
	assert.Equal(t, 2, FindShort("Let's travel abroad shall we"))
}
