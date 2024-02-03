package kata

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/54b42f9314d9229fd6000d9c

func DuplicateEncode(word string) string {
	word = strings.ToLower(word)
	charCount := make(map[rune]int)
	var result strings.Builder

	for _, char := range word {
		charCount[char]++
	}

	for _, char := range word {
		if charCount[char] > 1 {
			result.WriteRune(')')
		} else {
			result.WriteRune('(')
		}
	}

	return result.String()
}

func TestDuplicateEncode(t *testing.T) {
	assert.Equal(t, "(((", DuplicateEncode("din"))
	assert.Equal(t, "()()()", DuplicateEncode("recede"))
	assert.Equal(t, "))((", DuplicateEncode("(( @"))
	assert.Equal(t, ")())())", DuplicateEncode("Success"))
}
