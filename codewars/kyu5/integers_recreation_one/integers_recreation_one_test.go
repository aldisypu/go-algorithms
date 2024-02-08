package kata

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/55aa075506463dac6600010d

func sumOfSquaredDivisors(num int) int {
	sum := 0
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			sum += i * i
		}
	}
	return sum
}

func ListSquared(m, n int) [][]int {
	result := [][]int{}

	for num := m; num <= n; num++ {
		sum := sumOfSquaredDivisors(num)
		sqrtSum := int(math.Sqrt(float64(sum)))

		if sqrtSum*sqrtSum == sum {
			result = append(result, []int{num, sum})
		}
	}

	return result
}

func TestListSquared(t *testing.T) {
	assert.Equal(t, [][]int{{1, 1}, {42, 2500}, {246, 84100}}, ListSquared(1, 250))
	assert.Equal(t, [][]int{{287, 84100}}, ListSquared(250, 500))
	assert.Equal(t, [][]int{}, ListSquared(300, 600))
}
