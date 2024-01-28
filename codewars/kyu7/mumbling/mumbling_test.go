package kata

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/5667e8f4e3f572a8f2000039

func Accum(s string) string {
	var result string

	for i, char := range s {
		result += strings.ToUpper(string(char))

		result += strings.Repeat(strings.ToLower(string(char)), i)

		if i < len(s)-1 {
			result += "-"
		}
	}

	return result
}

func TestAccum(t *testing.T) {
	assert.Equal(t, "Z-Pp-Ggg-Llll-Nnnnn-Rrrrrr-Xxxxxxx-Qqqqqqqq-Eeeeeeeee-Nnnnnnnnnn-Uuuuuuuuuuu", Accum("ZpglnRxqenU"))
	assert.Equal(t, "N-Yy-Fff-Ffff-Sssss-Gggggg-Eeeeeee-Yyyyyyyy-Yyyyyyyyy-Llllllllll-Bbbbbbbbbbb", Accum("NyffsGeyylB"))
	assert.Equal(t, "M-Jj-Ttt-Kkkk-Uuuuu-Bbbbbb-Ooooooo-Vvvvvvvv-Qqqqqqqqq-Rrrrrrrrrr-Uuuuuuuuuuu", Accum("MjtkuBovqrU"))
	assert.Equal(t, "E-Vv-Iii-Dddd-Jjjjj-Uuuuuu-Nnnnnnn-Oooooooo-Kkkkkkkkk-Mmmmmmmmmm-Mmmmmmmmmmm", Accum("EvidjUnokmM"))
	assert.Equal(t, "H-Bb-Iii-Dddd-Eeeee-Vvvvvv-Bbbbbbb-Xxxxxxxx-Nnnnnnnnn-Cccccccccc-Ccccccccccc", Accum("HbideVbxncC"))
}
