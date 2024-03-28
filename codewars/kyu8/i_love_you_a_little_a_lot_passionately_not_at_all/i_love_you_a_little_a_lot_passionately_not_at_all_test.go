package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/57f24e6a18e9fad8eb000296

func HowMuchILoveYou(i int) string {
	phrases := []string{"I love you", "a little", "a lot", "passionately", "madly", "not at all"}
	index := (i - 1) % len(phrases)
	return phrases[index]
}

func TestHowMuchILoveYou(t *testing.T) {
	assert.Equal(t, "I love you", HowMuchILoveYou(7))
	assert.Equal(t, "a lot", HowMuchILoveYou(3))
	assert.Equal(t, "not at all", HowMuchILoveYou(6))
}
