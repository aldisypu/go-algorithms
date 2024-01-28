package kata

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/54ff3102c1bad923760001f3

func GetCount(str string) (count int) {
	for _, char := range str {
		if strings.ContainsRune("aeiou", char) {
			count++
		}
	}
	return count
}

func TestGetCount(t *testing.T) {
	assert.Equal(t, 5, GetCount("abracadabra"))
}
