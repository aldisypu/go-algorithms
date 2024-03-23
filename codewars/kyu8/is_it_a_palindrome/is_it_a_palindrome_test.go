package kata

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/57a1fd2ce298a731b20006a4

func IsPalindrome(str string) bool {
	str = strings.ToLower(str)
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}

func TestIsPalindrome(t *testing.T) {
	assert.True(t, IsPalindrome("a"))
	assert.True(t, IsPalindrome("aba"))
	assert.True(t, IsPalindrome("Abba"))
	assert.False(t, IsPalindrome("hello"))
}
