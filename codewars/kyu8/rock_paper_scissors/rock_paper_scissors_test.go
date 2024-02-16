package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Rps(p1, p2 string) string {
	switch {
	case (p1 == "rock" && p2 == "scissors") || (p1 == "scissors" && p2 == "paper") || (p1 == "paper" && p2 == "rock"):
		return "Player 1 won!"
	case (p2 == "rock" && p1 == "scissors") || (p2 == "scissors" && p1 == "paper") || (p2 == "paper" && p1 == "rock"):
		return "Player 2 won!"
	default:
		return "Draw!"
	}
}

func TestRps(t *testing.T) {
	assert.Equal(t, "Player 1 won!", Rps("rock", "scissors"))
	assert.Equal(t, "Player 2 won!", Rps("scissors", "rock"))
	assert.Equal(t, "Draw!", Rps("rock", "rock"))
}
