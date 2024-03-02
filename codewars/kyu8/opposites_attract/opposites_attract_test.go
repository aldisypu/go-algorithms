package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/555086d53eac039a2a000083

func LoveFunc(flower1, flower2 int) bool {
	return flower1%2 != flower2%2
}

func TestLoveFunc(t *testing.T) {
	assert.True(t, LoveFunc(1, 4))
	assert.True(t, LoveFunc(0, 1))
	assert.False(t, LoveFunc(2, 2))
	assert.False(t, LoveFunc(0, 0))
}
