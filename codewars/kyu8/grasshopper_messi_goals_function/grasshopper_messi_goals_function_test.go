package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/55f73be6e12baaa5900000d4

func Goals(laLigaGoals, copaDelReyGoals, championsLeagueGoals int) int {
	return laLigaGoals + copaDelReyGoals + championsLeagueGoals
}

func TestGoals(t *testing.T) {
	assert.Equal(t, 0, Goals(0, 0, 0))
	assert.Equal(t, 58, Goals(43, 10, 5))
}
