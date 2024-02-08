package kata

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/54d512e62a5e54c96200019e

func PrimeFactors(n int) string {
	var result strings.Builder
	current := 2

	for n > 1 {
		count := 0
		for n%current == 0 {
			n /= current
			count++
		}

		if count > 0 {
			result.WriteString(fmt.Sprintf("(%d", current))
			if count > 1 {
				result.WriteString(fmt.Sprintf("**%d", count))
			}
			result.WriteString(")")
		}

		current++
	}

	return result.String()
}

func TestPrimeFactors(t *testing.T) {
	assert.Equal(t, "(2**2)(3**3)(5)(7)(11**2)(17)", PrimeFactors(7775460))
	assert.Equal(t, "(2**2)(5)(3967)", PrimeFactors(79340))
}
