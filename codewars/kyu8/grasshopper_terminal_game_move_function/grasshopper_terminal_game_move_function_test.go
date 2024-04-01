package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/563a631f7cbbc236cf0000c2

func Move(position int, roll int) int {
	return position + roll*2
}

func TestMove(t *testing.T) {
	assert.Equal(t, 8, Move(0, 4))
	assert.Equal(t, 15, Move(3, 6))
	assert.Equal(t, 12, Move(2, 5))
}
