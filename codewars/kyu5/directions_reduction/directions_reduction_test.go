package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/550f22f4d758534c1100025a

func DirReduc(arr []string) []string {
	opposites := map[string]string{
		"NORTH": "SOUTH",
		"SOUTH": "NORTH",
		"EAST":  "WEST",
		"WEST":  "EAST",
	}

	var result []string

	for _, dir := range arr {
		if len(result) > 0 && opposites[result[len(result)-1]] == dir {
			result = result[:len(result)-1]
		} else {
			result = append(result, dir)
		}
	}

	return result
}

func TestDirReduc(t *testing.T) {
	assert.Equal(t, []string{}, DirReduc([]string{"NORTH", "SOUTH", "EAST", "WEST"}))
	assert.Equal(t, []string{"WEST"}, DirReduc([]string{"NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "WEST"}))
	assert.Equal(t, []string{"NORTH", "WEST", "SOUTH", "EAST"}, DirReduc([]string{"NORTH", "WEST", "SOUTH", "EAST"}))
	assert.Equal(t, []string{"NORTH"}, DirReduc([]string{"NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "NORTH"}))
}
