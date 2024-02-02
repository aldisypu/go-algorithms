package kata

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/5264d2b162488dc400000001

func SpinWords(str string) string {
	words := strings.Fields(str)

	for i, word := range words {
		if len(word) >= 5 {
			words[i] = reverseWord(word)
		}
	}

	return strings.Join(words, " ")
}

func reverseWord(word string) string {
	runes := []rune(word)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func TestSpinWords(t *testing.T) {
	assert.Equal(t, "emocleW", SpinWords("Welcome"))
	assert.Equal(t, "to", SpinWords("to"))
	assert.Equal(t, "sraWedoC", SpinWords("CodeWars"))
}
