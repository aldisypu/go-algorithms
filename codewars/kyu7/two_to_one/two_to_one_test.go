package kata

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/5656b6906de340bd1b0000ac

func TwoToOne(s1 string, s2 string) string {
	combined := s1 + s2

	uniqueChars := make(map[rune]struct{})
	for _, char := range combined {
		uniqueChars[char] = struct{}{}
	}

	var uniqueSlice []rune
	for char := range uniqueChars {
		uniqueSlice = append(uniqueSlice, char)
	}

	sort.Slice(uniqueSlice, func(i, j int) bool {
		return uniqueSlice[i] < uniqueSlice[j]
	})

	return string(uniqueSlice)
}

func TestTwoToOne(t *testing.T) {
	assert.Equal(t, "aehrsty", TwoToOne("aretheyhere", "yestheyarehere"))
	assert.Equal(t, "abcdefghilnoprstu", TwoToOne("loopingisfunbutdangerous", "lessdangerousthancoding"))
}
