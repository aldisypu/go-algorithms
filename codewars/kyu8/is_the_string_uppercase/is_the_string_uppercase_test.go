package kata

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/56cd44e1aa4ac7879200010b

type MyString string

func (s MyString) IsUpperCase() bool {
	return s == MyString(strings.ToUpper(string(s)))
}

func TestIsUpperCase(t *testing.T) {
	assert.True(t, MyString("WHAT DOES THE FOX SAY").IsUpperCase())
	assert.True(t, MyString("HTML CSS JAVASCRIPT PYTHON C PERL LISP JAVA GO RUBY NODEJS RUST SCALA").IsUpperCase())
	assert.False(t, MyString("a").IsUpperCase())
	assert.False(t, MyString("False").IsUpperCase())
	assert.False(t, MyString("false").IsUpperCase())
}
