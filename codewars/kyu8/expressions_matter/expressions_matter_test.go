package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/5ae62fcf252e66d44d00008e

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func ExpressionMatter(a int, b int, c int) int {
	return max(
		a*(b+c),
		max(
			a*b*c,
			max(
				(a+b)*c,
				a+b+c,
			),
		),
	)
}

func TestExpressionMatter(t *testing.T) {
	assert.Equal(t, 6, ExpressionMatter(2, 1, 2))
	assert.Equal(t, 4, ExpressionMatter(2, 1, 1))
	assert.Equal(t, 3, ExpressionMatter(1, 1, 1))
	assert.Equal(t, 9, ExpressionMatter(1, 2, 3))
	assert.Equal(t, 5, ExpressionMatter(1, 3, 1))
}
