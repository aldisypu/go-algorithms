package kata

import (
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/52bc74d4ac05d0945d00054e

func FirstNonRepeating(str string) string {
	charCount := make(map[rune]int)

	for _, char := range str {
		charCount[unicode.ToLower(char)]++
	}

	for _, char := range str {
		if charCount[unicode.ToLower(char)] == 1 {
			return string(char)
		}
	}

	return ""
}

func TestFirstNonRepeating(t *testing.T) {
	assert.Equal(t, "a", FirstNonRepeating("a"))
	assert.Equal(t, "", FirstNonRepeating(""))
	assert.Equal(t, "", FirstNonRepeating("aa"))
	assert.Equal(t, "w", FirstNonRepeating("hello world, eh?"))
	assert.Equal(t, "T", FirstNonRepeating("sTreSS"))
}
