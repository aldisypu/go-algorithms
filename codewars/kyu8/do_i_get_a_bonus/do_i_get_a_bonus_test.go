package kata

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/56f6ad906b88de513f000d96

func BonusTime(salary int, bonus bool) string {
	if bonus {
		salary *= 10
	}
	return fmt.Sprintf("£%d", salary)
}

func TestBonusTime(t *testing.T) {
	assert.Equal(t, "£100", BonusTime(100, false))
	assert.Equal(t, "£9", BonusTime(9, false))
	assert.Equal(t, "£55000", BonusTime(55000, false))
	assert.Equal(t, "£1000", BonusTime(100, true))
	assert.Equal(t, "£140000", BonusTime(14000, true))
}
