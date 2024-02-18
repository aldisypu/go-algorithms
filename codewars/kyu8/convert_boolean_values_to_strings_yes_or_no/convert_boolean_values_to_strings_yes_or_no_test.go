package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/53369039d7ab3ac506000467

func BoolToWord(word bool) string {
	switch word {
	case true:
		return "Yes"
	case false:
		return "No"
	default:
		return ""
	}

}

func TestBoolToWord(t *testing.T) {
	assert.Equal(t, "Yes", BoolToWord(true))
	assert.Equal(t, "No", BoolToWord(false))
}
