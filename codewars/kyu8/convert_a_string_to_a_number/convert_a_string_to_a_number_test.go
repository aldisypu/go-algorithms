package kata

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/544675c6f971f7399a000e79

func StringToNumber(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return num
}

func TestStringToNumber(t *testing.T) {
	assert.Equal(t, 1234, StringToNumber("1234"))
	assert.Equal(t, 605, StringToNumber("605"))
	assert.Equal(t, 1405, StringToNumber("1405"))
	assert.Equal(t, -7, StringToNumber("-7"))
}
