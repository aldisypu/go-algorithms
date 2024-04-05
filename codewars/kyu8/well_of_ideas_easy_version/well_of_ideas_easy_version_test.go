package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/57f222ce69e09c3630000212

func Well(x []string) string {
	goodCount := 0

	for _, idea := range x {
		if idea == "good" {
			goodCount++
		}
	}

	switch {
	case goodCount == 0:
		return "Fail!"
	case goodCount <= 2:
		return "Publish!"
	default:
		return "I smell a series!"
	}
}

func TestWell(t *testing.T) {
	assert.Equal(t, "I smell a series!", Well([]string{"good", "bad", "good", "good", "bad", "good", "bad", "bad", "good", "bad", "bad"}))
	assert.Equal(t, "Publish!", Well([]string{"bad", "bad", "bad", "bad", "good", "good", "bad", "bad", "bad"}))
	assert.Equal(t, "Fail!", Well([]string{"bad", "bad"}))
	assert.Equal(t, "Fail!", Well([]string{"bad"}))
	assert.Equal(t, "Publish!", Well([]string{"bad", "bad", "bad", "bad", "good", "bad"}))
}
