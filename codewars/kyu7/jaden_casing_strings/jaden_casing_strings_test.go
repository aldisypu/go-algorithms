package kata

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/5390bac347d09b7da40006f6

func ToJadenCase(str string) string {
	return strings.Title(str)
}

func TestToJadenCase(t *testing.T) {
	assert.Equal(t, "Most Trees Are Blue", ToJadenCase("most trees are blue"))
	assert.Equal(t, "All The Rules In This World Were Made By Someone No Smarter Than You. So Make Your Own.", ToJadenCase("All the rules in this world were made by someone no smarter than you. So make your own."))
	assert.Equal(t, "When I Die. Then You Will Realize", ToJadenCase("When I die. then you will realize"))
	assert.Equal(t, "Jonah Hill Is A Genius", ToJadenCase("Jonah Hill is a genius"))
	assert.Equal(t, "Dying Is Mainstream", ToJadenCase("Dying is mainstream"))
}
