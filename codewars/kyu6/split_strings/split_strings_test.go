package kata

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/515de9ae9dcfc28eb6000001

func Solution(str string) []string {
	return regexp.MustCompile(".{2}").FindAllString(str+"_", -1)
}

func TestSolution(t *testing.T) {
	assert.Equal(t, []string{"ab", "c_"}, Solution("abc"))
	assert.Equal(t, []string{"da", "ws", "d_"}, Solution("dawsd"))
	assert.Equal(t, []string{"aw", "sa", "ws"}, Solution("awsaws"))
}
