package kata

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/52fba66badcd10859f00097e

func Disemvowel(comment string) string {
	var result string

	for _, char := range comment {
		if !strings.ContainsRune("aeiouAEIOU", char) {
			result += string(char)
		}
	}
	return result
}

func TestDisemvowel(t *testing.T) {
	assert.Equal(t, "Ths wbst s fr lsrs LL!", Disemvowel("This website is for losers LOL!"))
}
