package kata

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/517abf86da9663f1d2000003

func ToCamelCase(s string) string {
	words := strings.FieldsFunc(s, func(r rune) bool {
		return r == '-' || r == '_'
	})

	for i := 1; i < len(words); i++ {
		words[i] = strings.ToUpper(string(words[i][0])) + words[i][1:]
	}

	return strings.Join(words, "")
}

func TestToCamelCase(t *testing.T) {
	assert.Equal(t, "", ToCamelCase(""))
	assert.Equal(t, "TheStealthWarrior", ToCamelCase("The_Stealth_Warrior"))
	assert.Equal(t, "theStealthWarrior", ToCamelCase("the-Stealth-Warrior"))
}
