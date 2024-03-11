package kata

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/51c8991dee245d7ddf00000e

func ReverseWords(str string) string {
	words := strings.Fields(str)

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	result := strings.Join(words, " ")

	return result
}

func TestReverseWords(t *testing.T) {
	assert.Equal(t, "world! hello", ReverseWords("hello world!"))
	assert.Equal(t, "this like speak doesn't yoda", ReverseWords("yoda doesn't speak like this"))
	assert.Equal(t, "foobar", ReverseWords("foobar"))
	assert.Equal(t, "editor kata", ReverseWords("kata editor"))
	assert.Equal(t, "boat your row row row", ReverseWords("row row row your boat"))
}
