package kata

import (
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/57cfdf34902f6ba3d300001e

func TwoSort(arr []string) string {
	sort.Strings(arr)

	return strings.Join(strings.Split(arr[0], ""), "***")
}

func TestTwoSort(t *testing.T) {
	assert.Equal(t, "b***i***t***c***o***i***n", TwoSort([]string{"bitcoin", "take", "over", "the", "world", "maybe", "who", "knows", "perhaps"}))
	assert.Equal(t, "a***r***e", TwoSort([]string{"turns", "out", "random", "test", "cases", "are", "easier", "than", "writing", "out", "basic", "ones"}))
	assert.Equal(t, "a***b***o***u***t", TwoSort([]string{"lets", "talk", "about", "javascript", "the", "best", "language"}))
	assert.Equal(t, "c***o***d***e", TwoSort([]string{"i", "want", "to", "travel", "the", "world", "writing", "code", "one", "day"}))
	assert.Equal(t, "L***e***t***s", TwoSort([]string{"Lets", "all", "go", "on", "holiday", "somewhere", "very", "cold"}))
}
