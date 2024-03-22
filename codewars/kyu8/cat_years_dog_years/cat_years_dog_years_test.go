package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/5a6663e9fd56cb5ab800008b

func CalculateYears(years int) (result [3]int) {
	catYears := 0
	dogYears := 0

	for i := 1; i <= years; i++ {
		switch i {
		case 1:
			catYears += 15
			dogYears += 15
		case 2:
			catYears += 9
			dogYears += 9
		default:
			catYears += 4
			dogYears += 5
		}
	}

	return [3]int{years, catYears, dogYears}
}

func TestCalculateYears(t *testing.T) {
	assert.Equal(t, [3]int{1, 15, 15}, CalculateYears(1))
	assert.Equal(t, [3]int{2, 24, 24}, CalculateYears(2))
	assert.Equal(t, [3]int{10, 56, 64}, CalculateYears(10))
}
