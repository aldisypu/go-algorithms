package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/5672682212c8ecf83e000050

func DblLinear(n int) int {
	u := make([]int, n+1)
	u[0] = 1

	i, j := 0, 0

	for k := 1; k <= n; k++ {
		y, z := 2*u[i]+1, 3*u[j]+1
		if y < z {
			u[k], i = y, i+1
		} else if y > z {
			u[k], j = z, j+1
		} else {
			u[k], i, j = y, i+1, j+1
		}
	}

	return u[n]
}

func TestDblLinear(t *testing.T) {
	assert.Equal(t, 175, DblLinear(50))
	assert.Equal(t, 447, DblLinear(100))
	assert.Equal(t, 3355, DblLinear(500))
	assert.Equal(t, 8488, DblLinear(1000))
}
