package kata

import (
	"strings"
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/56efc695740d30f963000557

func ToAlternatingCase(str string) string {
	var result strings.Builder

	for _, char := range str {
		if unicode.IsLower(char) {
			result.WriteRune(unicode.ToUpper(char))
		} else if unicode.IsUpper(char) {
			result.WriteRune(unicode.ToLower(char))
		} else {
			result.WriteRune(char)
		}
	}

	return result.String()
}

func TestToAlternatingCase(t *testing.T) {
	assert.Equal(t, "HELLO WORLD", ToAlternatingCase("hello world"))
	assert.Equal(t, "hello world", ToAlternatingCase("HELLO WORLD"))
	assert.Equal(t, "HELLO world", ToAlternatingCase("hello WORLD"))
	assert.Equal(t, "hEllO wOrld", ToAlternatingCase("HeLLo WoRLD"))
	assert.Equal(t, "12345", ToAlternatingCase("12345"))
}
