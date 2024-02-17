package kata

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/5bb904724c47249b10000131

func Points(games []string) int {
	totalPoints := 0

	for _, result := range games {
		x, _ := strconv.Atoi(string(result[0]))
		y, _ := strconv.Atoi(string(result[2]))

		switch {
		case x > y:
			totalPoints += 3
		case x == y:
			totalPoints += 1
		}
	}

	return totalPoints
}

func TestPoints(t *testing.T) {
	assert.Equal(t, 30, Points([]string{"1:0", "2:0", "3:0", "4:0", "2:1", "3:1", "4:1", "3:2", "4:2", "4:3"}))
	assert.Equal(t, 10, Points([]string{"1:1", "2:2", "3:3", "4:4", "2:2", "3:3", "4:4", "3:3", "4:4", "4:4"}))
	assert.Equal(t, 0, Points([]string{"0:1", "0:2", "0:3", "0:4", "1:2", "1:3", "1:4", "2:3", "2:4", "3:4"}))
	assert.Equal(t, 15, Points([]string{"1:0", "2:0", "3:0", "4:0", "2:1", "1:3", "1:4", "2:3", "2:4", "3:4"}))
	assert.Equal(t, 12, Points([]string{"1:0", "2:0", "3:0", "4:4", "2:2", "3:3", "1:4", "2:3", "2:4", "3:4"}))
}
