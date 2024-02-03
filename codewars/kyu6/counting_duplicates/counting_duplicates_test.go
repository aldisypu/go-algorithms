package kata

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/54bf1c2cd5b56cc47f0007a1

func duplicate_count(s1 string) int {
	charCount := make(map[rune]int)
	duplicateCount := 0

	s1 = strings.ToLower(s1)

	for _, char := range s1 {
		charCount[char]++
		if charCount[char] == 2 {
			duplicateCount++
		}
	}

	return duplicateCount
}

func TestDuplicateCount(t *testing.T) {
	assert.Equal(t, 0, duplicate_count("abcde"))
	assert.Equal(t, 1, duplicate_count("abcdea"))
	assert.Equal(t, 3, duplicate_count("abcdeaB11"))
	assert.Equal(t, 1, duplicate_count("indivisibility"))
}
