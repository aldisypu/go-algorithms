package kata

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/57e76bc428d6fbc2d500036d

func StringToArray(str string) []string {
	return strings.Fields(str)
}

func TestStringToArray(t *testing.T) {
	assert.Equal(t, []string{"Robin", "Singh"}, StringToArray("Robin Singh"))
	assert.Equal(t, []string{"CodeWars"}, StringToArray("CodeWars"))
	assert.Equal(t, []string{"I", "love", "arrays", "they", "are", "my", "favorite"}, StringToArray("I love arrays they are my favorite"))
	assert.Equal(t, []string{"1", "2", "3"}, StringToArray("1 2 3"))
}
