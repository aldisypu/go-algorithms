package kata

// https://www.codewars.com/kata/53d40c1e2f13e331fc000c26

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

var memo map[int64]*big.Int

func init() {
	memo = make(map[int64]*big.Int)
}

func fib(n int64) *big.Int {
	if n < 0 {
		sign := big.NewInt(1)
		if n%2 == 0 {
			sign.SetInt64(-1)
		}
		return new(big.Int).Mul(sign, fib(-n))
	}

	if val, found := memo[n]; found {
		return val
	}

	if n <= 1 {
		return big.NewInt(n)
	}

	matrix := [][]*big.Int{{big.NewInt(1), big.NewInt(1)}, {big.NewInt(1), big.NewInt(0)}}
	result := power(matrix, n-1)
	memo[n] = result[0][0]

	return result[0][0]
}

func multiply(a, b [][]*big.Int) [][]*big.Int {
	c := make([][]*big.Int, 2)
	for i := 0; i < 2; i++ {
		c[i] = make([]*big.Int, 2)
		for j := 0; j < 2; j++ {
			c[i][j] = new(big.Int)
			for k := 0; k < 2; k++ {
				c[i][j].Add(c[i][j], new(big.Int).Mul(a[i][k], b[k][j]))
			}
		}
	}
	return c
}

func power(matrix [][]*big.Int, n int64) [][]*big.Int {
	result := [][]*big.Int{{big.NewInt(1), big.NewInt(0)}, {big.NewInt(0), big.NewInt(1)}}
	for n > 0 {
		if n%2 == 1 {
			result = multiply(result, matrix)
		}
		matrix = multiply(matrix, matrix)
		n /= 2
	}
	return result
}

func TestFib(t *testing.T) {
	assert.Equal(t, big.NewInt(0), fib(0))
	assert.Equal(t, big.NewInt(1), fib(1))
	assert.Equal(t, big.NewInt(1), fib(2))
	assert.Equal(t, big.NewInt(2), fib(3))
	assert.Equal(t, big.NewInt(5), fib(5))
}
