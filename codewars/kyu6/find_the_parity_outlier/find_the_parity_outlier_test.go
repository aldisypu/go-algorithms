package kata

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/5526fc09a1bbd946250002dc

func FindOutlier(integers []int) int {
	evenCount, oddCount := 0, 0

	for _, num := range integers[:3] {
		if num%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}

	isMajorityEven := evenCount > oddCount

	for _, num := range integers {
		if num%2 == 0 && !isMajorityEven {
			return num
		} else if num%2 != 0 && isMajorityEven {
			return num
		}
	}

	return 0
}

func TestFindOutlier(t *testing.T) {
	assert.Equal(t, 3, FindOutlier([]int{2, 6, 8, -10, 3}))
	assert.Equal(t, 206847684, FindOutlier([]int{206847684, 1056521, 7, 17, 1901, 21104421, 7, 1, 35521, 1, 7781}))
	assert.Equal(t, 0, FindOutlier([]int{math.MaxInt32, 0, 1}))
}
