package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/5541f58a944b85ce6d00006a

func ProductFib(prod uint64) [3]uint64 {
	a, b := uint64(0), uint64(1)

	for a*b < prod {
		a, b = b, a+b
	}

	if a*b == prod {
		return [3]uint64{a, b, 1}
	}

	return [3]uint64{a, b, 0}
}

func TestProductFib(t *testing.T) {
	assert.Equal(t, [3]uint64{55, 89, 1}, ProductFib(4895))
	assert.Equal(t, [3]uint64{89, 144, 0}, ProductFib(5895))
	assert.Equal(t, [3]uint64{6765, 10946, 1}, ProductFib(74049690))
	assert.Equal(t, [3]uint64{10946, 17711, 0}, ProductFib(84049690))
}
