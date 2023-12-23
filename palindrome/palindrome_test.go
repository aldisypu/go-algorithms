package palindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func isPalindrome(value string) bool {
	rune := []rune(value)
	for i := 0; i < len(rune)/2; i++ {
		if string(rune[i]) != string(rune[len(rune)-i-1]) {
			return false
		}
	}
	return true
}

func isPalindromeRecursive(value string) bool {
	return isPalindromeRecursiveHelper(value, 0)
}

func isPalindromeRecursiveHelper(value string, i int) bool {
	rune := []rune(value)
	if i < len(rune)/2 {
		if string(rune[i]) != string(rune[len(rune)-i-1]) {
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
