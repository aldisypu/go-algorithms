package kata

import (
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/5629db57620258aa9d000014

func Mix(s1, s2 string) string {
	counts1 := countLetters(s1)
	counts2 := countLetters(s2)

	var result []string

	for char := 'a'; char <= 'z'; char++ {
		maxCount := max(counts1[char], counts2[char])
		if maxCount > 1 {
			prefix := ""
			if counts1[char] == maxCount && counts2[char] == maxCount {
				prefix = "=:"
			} else if counts1[char] == maxCount {
				prefix = "1:"
			} else if counts2[char] == maxCount {
				prefix = "2:"
			}
			result = append(result, prefix+strings.Repeat(string(char), maxCount))
		}
	}

	sort.Slice(result, func(i, j int) bool {
		if len(result[i]) == len(result[j]) {
			return result[i] < result[j]
		}
		return len(result[i]) > len(result[j])
	})

	return strings.Join(result, "/")
}

func countLetters(s string) map[rune]int {
	counts := make(map[rune]int)
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			counts[char]++
		}
	}
	return counts
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestMix(t *testing.T) {
	assert.Equal(t, "2:eeeee/2:yy/=:hh/=:rr", Mix("Are they here", "yes, they are here"))
	assert.Equal(t, "=:uuuuuu", Mix("uuuuuu", "uuuuuu"))
	assert.Equal(t, "1:ooo/1:uuu/2:sss/=:nnn/1:ii/2:aa/2:dd/2:ee/=:gg", Mix("looping is fun but dangerous", "less dangerous than coding"))
}
