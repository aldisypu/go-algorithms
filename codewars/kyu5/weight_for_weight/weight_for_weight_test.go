package kata

import (
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/55c6126177c9441a570000cc

func sumDigits(num string) int {
	sum := 0
	for _, digit := range num {
		digitValue, _ := strconv.Atoi(string(digit))
		sum += digitValue
	}
	return sum
}

func OrderWeight(strng string) string {
	weights := strings.Fields(strng)

	sort.SliceStable(weights, func(i, j int) bool {
		sumA, sumB := sumDigits(weights[i]), sumDigits(weights[j])

		if sumA == sumB {
			return weights[i] < weights[j]
		}

		return sumA < sumB
	})

	result := strings.Join(weights, " ")

	return result
}

func TestOrderWeight(t *testing.T) {
	assert.Equal(t, "2000 103 123 4444 99", OrderWeight("103 123 4444 99 2000"))
	assert.Equal(t, "11 11 2000 10003 22 123 1234000 44444444 9999", OrderWeight("2000 10003 1234000 44444444 9999 11 11 22 123"))
	assert.Equal(t, "", OrderWeight(""))
}
