package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/525c65e51bf619685c000059

func Cakes(recipe, available map[string]int) int {
	maxCount := int(^uint(0) >> 1)

	for ingredient, requiredAmount := range recipe {
		availableAmount, ingredientExists := available[ingredient]

		if !ingredientExists {
			return 0
		}

		count := availableAmount / requiredAmount

		if count < maxCount {
			maxCount = count
		}
	}

	return maxCount
}

func TestCakes(t *testing.T) {
	assert.Equal(t, 2, Cakes(map[string]int{"flour": 500, "sugar": 200, "eggs": 1}, map[string]int{"flour": 1200, "sugar": 1200, "eggs": 5, "milk": 200}))
	assert.Equal(t, 0, Cakes(map[string]int{"apples": 3, "flour": 300, "sugar": 150, "milk": 100, "oil": 100}, map[string]int{"sugar": 500, "flour": 2000, "milk": 2000}))
}
