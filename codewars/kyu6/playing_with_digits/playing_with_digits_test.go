package kata

import (
	"math"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/5552101f47fc5178b1000050

func DigPow(n, p int) int {
	strN := strconv.Itoa(n)
	sum := 0

	for _, s := range strN {
		digit, _ := strconv.Atoi(string(s))
		sum += int(math.Pow(float64(digit), float64(p)))
		p++
	}

	if sum%n == 0 {
		return sum / n
	}

	return -1
}

func TestDigPow(t *testing.T) {
	assert.Equal(t, 1, DigPow(89, 1))
	assert.Equal(t, -1, DigPow(92, 1))
}
