package kata

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/51ba717bb08c1cd60f00002f

func Solution(list []int) string {
	if len(list) == 0 {
		return ""
	}

	var result strings.Builder
	startRange := list[0]

	appendRange := func(start, end int) {
		if start == end {
			result.WriteString(strconv.Itoa(start))
		} else if end-start >= 2 {
			result.WriteString(fmt.Sprintf("%d-%d", start, end))
		} else {
			result.WriteString(fmt.Sprintf("%d,%d", start, end))
		}
		result.WriteByte(',')
	}

	for i := 1; i < len(list); i++ {
		if list[i] != list[i-1]+1 {
			appendRange(startRange, list[i-1])
			startRange = list[i]
		}
	}

	appendRange(startRange, list[len(list)-1])

	return strings.TrimRight(result.String(), ",")
}

func TestSolution(t *testing.T) {
	assert.Equal(t, "-6,-3-1,3-5,7-11,14,15,17-20", Solution([]int{-6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20}))
}
