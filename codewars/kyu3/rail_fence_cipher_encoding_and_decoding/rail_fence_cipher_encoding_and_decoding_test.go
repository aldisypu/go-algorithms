package kata

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/58c5577d61aefcf3ff000081

func Encode(s string, n int) string {
	if n <= 1 {
		return s
	}

	rails := make([]strings.Builder, n)
	downward := true
	railIndex := 0

	for _, char := range s {
		rails[railIndex].WriteRune(char)

		if downward {
			railIndex++
		} else {
			railIndex--
		}

		if railIndex == 0 || railIndex == n-1 {
			downward = !downward
		}
	}

	var result strings.Builder
	for _, rail := range rails {
		result.WriteString(rail.String())
	}

	return result.String()
}

func Decode(s string, n int) string {
	if n <= 1 {
		return s
	}

	rails := make([][]rune, n)
	downward := true
	railIndex := 0

	for i := 0; i < len(s); i++ {
		rails[railIndex] = append(rails[railIndex], rune(0)) // placeholder to maintain correct rail lengths

		if downward {
			railIndex++
		} else {
			railIndex--
		}

		if railIndex == 0 || railIndex == n-1 {
			downward = !downward
		}
	}

	index := 0
	for i := 0; i < n; i++ {
		for j := range rails[i] {
			rails[i][j] = rune(s[index])
			index++
		}
	}

	var result strings.Builder
	railIndex = 0
	downward = true

	for i := 0; i < len(s); i++ {
		result.WriteRune(rails[railIndex][0])
		rails[railIndex] = rails[railIndex][1:]

		if downward {
			railIndex++
		} else {
			railIndex--
		}

		if railIndex == 0 || railIndex == n-1 {
			downward = !downward
		}
	}

	return result.String()
}

func TestEncodeDecode(t *testing.T) {
	assert.Equal(t, "WECRLTEERDSOEEFEAOCAIVDEN", Encode("WEAREDISCOVEREDFLEEATONCE", 3))
	assert.Equal(t, "WEAREDISCOVEREDFLEEATONCE", Decode("WECRLTEERDSOEEFEAOCAIVDEN", 3))
	assert.Equal(t, "Hoo!el,Wrdl l", Encode("Hello, World!", 3))
}
