package kata

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/55a70521798b14d4750000a4

func Greet(name string) string {
	return fmt.Sprintf("Hello, %s how are you doing today?", name)
}

func TestGreet(t *testing.T) {
	assert.Equal(t, "Hello, Ryan how are you doing today?", Greet("Ryan"))
}
