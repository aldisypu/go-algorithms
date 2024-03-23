package palindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func isPalindrome(value string) bool {
	for i := 0; i < len(value)/2; i++ {
		if value[i] != value[len(value)-i-1] {
			return false
		}
	}
	return true
}

func isPalindromeRecursive(value string) bool {
	return isPalindromeRecursiveHelper(value, 0)
}

func isPalindromeRecursiveHelper(value string, i int) bool {
	if i < len(value)/2 {
		if value[i] != value[len(value)-i-1] {
			return false
		} else {
			return isPalindromeRecursiveHelper(value, i+1)
		}
	} else {
		return true
	}
}

func TestPalindrome(t *testing.T) {
	assert.True(t, isPalindrome("aba"))
	assert.True(t, isPalindrome("aba"))
	assert.True(t, isPalindrome("malam"))
	assert.True(t, isPalindrome(""))

	assert.False(t, isPalindrome("ab"))
	assert.False(t, isPalindrome("abab"))
	assert.False(t, isPalindrome("siang"))
	assert.False(t, isPalindrome("adi"))

	assert.True(t, isPalindromeRecursive("aba"))
	assert.True(t, isPalindromeRecursive("aba"))
	assert.True(t, isPalindromeRecursive("malam"))
	assert.True(t, isPalindromeRecursive(""))

	assert.False(t, isPalindromeRecursive("ab"))
	assert.False(t, isPalindromeRecursive("abab"))
	assert.False(t, isPalindromeRecursive("siang"))
	assert.False(t, isPalindromeRecursive("adi"))
}
