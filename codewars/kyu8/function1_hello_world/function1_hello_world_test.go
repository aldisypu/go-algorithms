package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/523b4ff7adca849afe000035

func greet() string {
	return "hello world!"
}

func TestGreet(t *testing.T) {
	assert.Equal(t, "hello world!", greet())
}
