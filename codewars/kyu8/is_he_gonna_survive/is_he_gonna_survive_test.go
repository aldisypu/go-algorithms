package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/59ca8246d751df55cc00014c

func Hero(bullets, dragons int) bool {
	return bullets >= dragons*2
}

func TestHero(t *testing.T) {

	assert.Equal(t, true, Hero(10, 5))
	assert.Equal(t, false, Hero(7, 4))
	assert.Equal(t, false, Hero(4, 5))
	assert.Equal(t, true, Hero(100, 40))
	assert.Equal(t, false, Hero(1500, 751))
	assert.Equal(t, false, Hero(0, 1))
}
