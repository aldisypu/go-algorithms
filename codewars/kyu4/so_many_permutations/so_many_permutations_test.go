package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/5254ca2719453dcc0b00027d

func Permutations(s string) []string {
	result := make(map[string]struct{})
	characters := []byte(s)
	n := len(characters)

	swap := func(i, j int) {
		characters[i], characters[j] = characters[j], characters[i]
	}

	var generate func(int)
	generate = func(k int) {
		if k == 1 {
			result[string(characters)] = struct{}{}
		} else {
			for i := 0; i < k; i++ {
				generate(k - 1)
				if k%2 == 0 {
					swap(i, k-1)
				} else {
					swap(0, k-1)
				}
			}
		}
	}

	generate(n)

	uniqueResult := make([]string, 0, len(result))
	for perm := range result {
		uniqueResult = append(uniqueResult, perm)
	}

	return uniqueResult
}

func TestPermutations(t *testing.T) {
	assert.ElementsMatch(t, []string{"a"}, Permutations("a"))
	assert.ElementsMatch(t, []string{"ab", "ba"}, Permutations("ab"))
	assert.ElementsMatch(t, []string{"abc", "acb", "bac", "bca", "cab", "cba"}, Permutations("abc"))
	assert.ElementsMatch(t, []string{"aa"}, Permutations("aa"))
	assert.ElementsMatch(t, []string{"aabb", "abab", "abba", "baab", "baba", "bbaa"}, Permutations("aabb"))
}
