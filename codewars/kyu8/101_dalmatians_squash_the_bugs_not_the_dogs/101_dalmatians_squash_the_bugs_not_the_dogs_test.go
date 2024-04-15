package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/56f6919a6b88de18ff000b36

func HowManyDalmatians(number int) string {
	switch {
	case number <= 10:
		return "Hardly any"
	case number <= 50:
		return "More than a handful!"
	case number == 101:
		return "101 DALMATIONS!!!"
	default:
		return "Woah that's a lot of dogs!"
	}
}

func TestHowManyDalmatians(t *testing.T) {
	assert.Equal(t, "More than a handful!", HowManyDalmatians(26))
	assert.Equal(t, "Hardly any", HowManyDalmatians(8))
	assert.Equal(t, "Woah that's a lot of dogs!", HowManyDalmatians(80))
	assert.Equal(t, "Hardly any", HowManyDalmatians(10))
	assert.Equal(t, "101 DALMATIONS!!!", HowManyDalmatians(101))
}
