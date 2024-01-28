package kata

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/554b4ac871d6813a03000035

func HighAndLow(in string) string {
	numbers := strings.Fields(in)

	if len(numbers) == 0 {
		return ""
	}

	highest, _ := strconv.Atoi(numbers[0])
	lowest, _ := strconv.Atoi(numbers[0])

	for _, numStr := range numbers {
		num, _ := strconv.Atoi(numStr)
		if num > highest {
			highest = num
		}
		if num < lowest {
			lowest = num
		}
	}

	return fmt.Sprintf("%d %d", highest, lowest)
}

func TestHighAndLow(t *testing.T) {
	assert.Equal(t, "3 1", HighAndLow("1 2 3"))
}
