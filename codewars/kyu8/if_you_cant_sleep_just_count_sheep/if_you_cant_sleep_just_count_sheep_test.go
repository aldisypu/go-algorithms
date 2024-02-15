package kata

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/5b077ebdaf15be5c7f000077

func countSheep(num int) string {
	var builder strings.Builder

	for i := 1; i <= num; i++ {
		builder.WriteString(strconv.Itoa(i))
		builder.WriteString(" sheep...")
	}

	return builder.String()
}

func TestCountSheep(t *testing.T) {
	assert.Equal(t, "1 sheep...2 sheep...", countSheep(2))
	assert.Equal(t, "", countSheep(0))
	assert.Equal(t, "1 sheep...", countSheep(1))
}
