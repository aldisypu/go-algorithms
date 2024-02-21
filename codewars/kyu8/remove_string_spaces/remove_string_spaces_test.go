package kata

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/57eae20f5500ad98e50002c5

func NoSpace(word string) string {
	return strings.ReplaceAll(word, " ", "")
}

func TestNoSpace(t *testing.T) {
	assert.Equal(t, "8j8mBliB8gimjB8B8jlB", NoSpace("8 j 8   mBliB8g  imjB8B8  jl  B"))
	assert.Equal(t, "88Bifk8hB8BB8BBBB888chl8BhBfd", NoSpace("8 8 Bi fk8h B 8 BB8B B B  B888 c hl8 BhB fd"))
	assert.Equal(t, "8aaaaaddddr", NoSpace("8aaaaa dddd r     "))
	assert.Equal(t, "jfBmgklf8hg88lbe8", NoSpace("jfBm  gk lf8hg  88lbe8 "))
	assert.Equal(t, "8jaam", NoSpace("8j aam"))
}
