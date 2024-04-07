package kata

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/5a34af40e1ce0eb1f5000036

func ToCsvText(array [][]int) string {
	var result strings.Builder

	for i, row := range array {
		for j, num := range row {
			result.WriteString(strconv.Itoa(num))
			if j < len(row)-1 {
				result.WriteByte(',')
			}
		}
		if i < len(array)-1 {
			result.WriteString("\n")
		}
	}

	return result.String()
}

func TestToCsvText(t *testing.T) {
	assert.Equal(t, "0,1,2,3,45\n10,11,12,13,14\n20,21,22,23,24\n30,31,32,33,34", ToCsvText([][]int{
		{0, 1, 2, 3, 45},
		{10, 11, 12, 13, 14},
		{20, 21, 22, 23, 24},
		{30, 31, 32, 33, 34}}))
	assert.Equal(t, "-25,21,2,-33,48\n30,31,-32,33,-34", ToCsvText([][]int{
		{-25, 21, 2, -33, 48},
		{30, 31, -32, 33, -34}}))
	assert.Equal(t, "5,55,5,5,55\n6,6,66,23,24\n666,31,66,33,7", ToCsvText([][]int{
		{5, 55, 5, 5, 55},
		{6, 6, 66, 23, 24},
		{666, 31, 66, 33, 7}}))
}
