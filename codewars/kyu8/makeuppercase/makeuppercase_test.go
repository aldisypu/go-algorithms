package kata

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/57a0556c7cb1f31ab3000ad7

func MakeUpperCase(str string) string {
	return strings.ToUpper(str)
}

func TestMakeUpperCase(t *testing.T) {
	assert.Equal(t, "HELLO", MakeUpperCase("hello"))
	assert.Equal(t, "HELLO WORLD", MakeUpperCase("hello world"))
	assert.Equal(t, "HELLO WORLD !", MakeUpperCase("hello world !"))
	assert.Equal(t, "HELLO WORLD !", MakeUpperCase("heLlO wORLd !"))
	assert.Equal(t, "1,2,3 HELLO WORLD!", MakeUpperCase("1,2,3 hello world!"))
}
