package kata

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/596c6eb85b0f515834000049

func ReplaceDots(str string) string {
	return regexp.MustCompile(`\.`).ReplaceAllString(str, "-")
}

func TestReplaceDots(t *testing.T) {
	assert.Equal(t, "", ReplaceDots(""))
	assert.Equal(t, "no dots", ReplaceDots("no dots"))
	assert.Equal(t, "one-two-three", ReplaceDots("one.two.three"))
	assert.Equal(t, "---", ReplaceDots("..."))
}
