package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/55bf01e5a717a0d57e0000ec

func Persistence(n int) int {
	persistenceCount := 0

	for n >= 10 {
		product := 1

		for ; n > 0; n /= 10 {
			product *= n % 10
		}

		n = product
		persistenceCount++
	}

	return persistenceCount
}

func TestPersistence(t *testing.T) {
	assert.Equal(t, 3, Persistence(39))
	assert.Equal(t, 0, Persistence(4))
	assert.Equal(t, 2, Persistence(25))
	assert.Equal(t, 4, Persistence(999))
}
