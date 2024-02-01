package kata

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/5259b20d6021e9e14c0010d4

func ReverseWords(str string) string {
	words := strings.Split(str, " ")

	for i, word := range words {
		words[i] = reverseString(word)
	}

	return strings.Join(words, " ")
}

func reverseString(s string) string {
	if len(s) <= 1 {
		return s
	}
	return string(s[len(s)-1]) + reverseString(s[:len(s)-1])
}

func TestReverseWords(t *testing.T) {
	assert.Equal(t, "ehT kciuq nworb xof spmuj revo eht yzal .god", ReverseWords("The quick brown fox jumps over the lazy dog."))
	assert.Equal(t, "elppa", ReverseWords("apple"))
	assert.Equal(t, "a b c d", ReverseWords("a b c d"))
	assert.Equal(t, "elbuod  decaps  sdrow", ReverseWords("double  spaced  words"))
	assert.Equal(t, "desserts stressed", ReverseWords("stressed desserts"))
	assert.Equal(t, "olleh olleh", ReverseWords("hello hello"))
}
