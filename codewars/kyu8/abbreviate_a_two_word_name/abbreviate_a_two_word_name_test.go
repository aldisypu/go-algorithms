package kata

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/57eadb7ecd143f4c9c0000a3

func AbbrevName(name string) string {
	parts := strings.Split(name, " ")
	return strings.ToUpper(string(parts[0][0])) + "." + strings.ToUpper(string(parts[1][0]))
}

func TestAbbrevName(t *testing.T) {
	assert.Equal(t, "S.H", AbbrevName("Sam Harris"))
	assert.Equal(t, "P.F", AbbrevName("Patrick Feenan"))
	assert.Equal(t, "E.C", AbbrevName("Evan Cole"))
	assert.Equal(t, "P.F", AbbrevName("P Favuzzi"))
	assert.Equal(t, "D.M", AbbrevName("David Mendieta"))
}
