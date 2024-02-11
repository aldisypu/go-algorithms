package kata

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/52b7ed099cdc285c300001cd

func SumOfIntervals(intervals [][2]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	mergedIntervals := make([][2]int, 0)
	mergedIntervals = append(mergedIntervals, intervals[0])

	for i := 1; i < len(intervals); i++ {
		lastMerged := &mergedIntervals[len(mergedIntervals)-1]
		if intervals[i][0] <= lastMerged[1] {
			// Intervals overlap, merge them
			if intervals[i][1] > lastMerged[1] {
				lastMerged[1] = intervals[i][1]
			}
		} else {
			mergedIntervals = append(mergedIntervals, intervals[i])
		}
	}

	sum := 0
	for _, interval := range mergedIntervals {
		sum += interval[1] - interval[0]
	}

	return sum
}

func TestSumOfIntervals(t *testing.T) {
	assert.Equal(t, 4, SumOfIntervals([][2]int{{1, 5}}))
	assert.Equal(t, 8, SumOfIntervals([][2]int{{1, 5}, {6, 10}}))
	assert.Equal(t, 4, SumOfIntervals([][2]int{{1, 5}, {1, 5}}))
	assert.Equal(t, 2_000_000_000, SumOfIntervals([][2]int{{-1_000_000_000, 1_000_000_000}}))
	assert.Equal(t, 100_000_030, SumOfIntervals([][2]int{{0, 20}, {-100_000_000, 10}, {30, 40}}))
}
