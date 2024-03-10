package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/55225023e1be1ec8bc000390

func Greet(name string) string {
	if name == "Johnny" {
		return "Hello, my love!"
	}
	return "Hello, " + name + "!"
}

func TestGreet(t *testing.T) {
	assert.Equal(t, "Hello, Alfred!", Greet("Alfred"))
	assert.Equal(t, "Hello, my love!", Greet("Johnny"))
}
