package kata

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/5265326f5fda8eb1160004c8

func NumberToString(n int) string {
	return strconv.Itoa(n)
}

func TestNumberToString(t *testing.T) {
	assert.Equal(t, "67", NumberToString(67))
	assert.Equal(t, "79585", NumberToString(79585))
	assert.Equal(t, "3", NumberToString(3))
	assert.Equal(t, "-1", NumberToString(-1))
}
