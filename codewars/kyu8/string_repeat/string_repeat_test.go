package kata

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/57a0e5c372292dd76d000d7e

func RepeatStr(repetitions int, value string) string {
	return strings.Repeat(value, repetitions)
}

func TestRepeatStr(t *testing.T) {
	assert.Equal(t, "aaaa", RepeatStr(4, "a"))
	assert.Equal(t, "hello hello hello ", RepeatStr(3, "hello "))
	assert.Equal(t, "abcabc", RepeatStr(2, "abc"))
}
