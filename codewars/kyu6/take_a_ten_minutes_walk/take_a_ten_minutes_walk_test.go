package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/54da539698b8a2ad76000228

func IsValidWalk(walk []rune) bool {
	if len(walk) != 10 {
		return false
	}

	var north, south, east, west int

	for _, direction := range walk {
		switch direction {
		case 'n':
			north++
		case 's':
			south++
		case 'e':
			east++
		case 'w':
			west++
		}
	}
	return north == south && east == west
}

func TestIsValidWalk(t *testing.T) {
	assert.True(t, IsValidWalk([]rune{'n', 's', 'n', 's', 'n', 's', 'n', 's', 'n', 's'}))
	assert.False(t, IsValidWalk([]rune{'w', 'e', 'w', 'e', 'w', 'e', 'w', 'e', 'w', 'e', 'w', 'e'}))
	assert.False(t, IsValidWalk([]rune{'w'}))
	assert.False(t, IsValidWalk([]rune{'n', 'n', 'n', 's', 'n', 's', 'n', 's', 'n', 's'}))
	assert.False(t, IsValidWalk([]rune{'e', 'e', 'e', 'e', 'w', 'w', 's', 's', 's', 's'}))
}
