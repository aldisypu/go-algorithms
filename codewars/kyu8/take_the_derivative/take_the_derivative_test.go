package kata

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/5963c18ecb97be020b0000a2

func Derive(coefficient, exponent int) string {
	return fmt.Sprintf("%dx^%d", coefficient*exponent, exponent-1)
}

func TestDerive(t *testing.T) {
	assert.Equal(t, "56x^7", Derive(7, 8))
	assert.Equal(t, "45x^8", Derive(5, 9))
}
