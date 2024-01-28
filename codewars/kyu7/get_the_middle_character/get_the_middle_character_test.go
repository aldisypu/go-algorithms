package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/56747fd5cb988479af000028

func GetMiddle(s string) string {
	length := len(s)

	if length <= 2 {
		return s
	}

	return GetMiddle(s[1 : length-1])
}

func TestGetMiddle(t *testing.T) {
	assert.Equal(t, "es", GetMiddle("test"))
	assert.Equal(t, "t", GetMiddle("testing"))
	assert.Equal(t, "dd", GetMiddle("middle"))
	assert.Equal(t, "A", GetMiddle("A"))
}
