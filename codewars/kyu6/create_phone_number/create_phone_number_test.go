package kata

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/525f50e3b73515a6db000b83

func CreatePhoneNumber(numbers [10]uint) string {
	return fmt.Sprintf("(%d%d%d) %d%d%d-%d%d%d%d", numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6], numbers[7], numbers[8], numbers[9])
}

func TestCreatePhoneNumber(t *testing.T) {
	assert.Equal(t, "(123) 456-7890", CreatePhoneNumber([10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}))
}
