package kata

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/54de279df565808f8b00126a

var MultipleOf3Regex = "^0*(1(01*0)*10*)*$"

func TestMatchMultipleOf3Regex(t *testing.T) {
	assert.False(t, matchMultipleOf3Regex(" 0"))
	assert.False(t, matchMultipleOf3Regex("abc"))
	assert.True(t, matchMultipleOf3Regex("000"))
	assert.True(t, matchMultipleOf3Regex("110"))
	assert.False(t, matchMultipleOf3Regex("111"))
}

func matchMultipleOf3Regex(s string) bool {
	matched, _ := regexp.MatchString(MultipleOf3Regex, s)
	return matched
}
