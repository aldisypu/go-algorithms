package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/56f695399400f5d9ef000af5

func CorrectTail(body string, tail rune) bool {
	return rune(body[len(body)-1]) == tail
}

func TestCorrectTail(t *testing.T) {
	assert.True(t, CorrectTail("Fox", 'x'))
	assert.True(t, CorrectTail("Rhino", 'o'))
	assert.True(t, CorrectTail("Meerkat", 't'))
}
