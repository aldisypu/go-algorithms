package kata

import (
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/55c45be3b2079eccff00010f

func Order(sentence string) string {
	words := strings.Fields(sentence)

	sort.Slice(words, func(i, j int) bool {
		num1, _ := strconv.Atoi(strings.TrimFunc(words[i], func(r rune) bool { return !('1' <= r && r <= '9') }))
		num2, _ := strconv.Atoi(strings.TrimFunc(words[j], func(r rune) bool { return !('1' <= r && r <= '9') }))
		return num1 < num2
	})

	return strings.Join(words, " ")
}

func TestOrder(t *testing.T) {
	assert.Equal(t, "Thi1s is2 3a T4est", Order("is2 Thi1s T4est 3a"))
	assert.Equal(t, "Fo1r the2 g3ood 4of th5e pe6ople", Order("4of Fo1r pe6ople g3ood th5e the2"))
	assert.Equal(t, "", Order(""))
}
