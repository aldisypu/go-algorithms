package kata

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/526dbd6c8c0eb53254000110

func alphanumeric(str string) bool {
	match, _ := regexp.MatchString("^[a-zA-Z0-9]+$", str)
	return match
}

func TestAlphanumeric(t *testing.T) {
	assert.False(t, alphanumeric(".*?"))
	assert.True(t, alphanumeric("a"))
	assert.True(t, alphanumeric("Mazinkaiser"))
	assert.False(t, alphanumeric("hello world_"))
	assert.True(t, alphanumeric("PassW0rd"))
}
