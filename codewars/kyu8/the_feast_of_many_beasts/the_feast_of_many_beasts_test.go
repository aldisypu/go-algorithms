package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/5aa736a455f906981800360d

func Feast(beast string, dish string) bool {
	return beast[0] == dish[0] && beast[len(beast)-1] == dish[len(dish)-1]
}

func TestFeast(t *testing.T) {
	assert.True(t, Feast("great blue heron", "garlic naan"))
	assert.True(t, Feast("chickadee", "chocolate cake"))
	assert.False(t, Feast("brown bear", "bear claw"))
}
