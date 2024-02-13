package kata

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/55cf3b567fc0e02b0b00000b

func Part(n int) string {
	partitions := generatePartitions(n)
	products := calculateProducts(partitions)
	result := calculateStats(products)

	return result
}

func generatePartitions(n int) [][]int {
	if n == 0 {
		return [][]int{{}}
	}

	var partitions [][]int
	generatePartitionsHelper(n, 1, []int{}, &partitions)
	return partitions
}

func generatePartitionsHelper(n, start int, current []int, partitions *[][]int) {
	if n == 0 {
		*partitions = append(*partitions, append([]int{}, current...))
		return
	}

	for i := start; i <= n; i++ {
		generatePartitionsHelper(n-i, i, append(current, i), partitions)
	}
}

func calculateProducts(partitions [][]int) []int {
	productSet := make(map[int]struct{})
	for _, partition := range partitions {
		product := 1
		for _, num := range partition {
			product *= num
		}
		productSet[product] = struct{}{}
	}

	products := make([]int, 0, len(productSet))
	for product := range productSet {
		products = append(products, product)
	}

	sort.Ints(products)
	return products
}

func calculateStats(products []int) string {
	rangeValue := products[len(products)-1] - products[0]

	sum := 0
	for _, num := range products {
		sum += num
	}
	average := float64(sum) / float64(len(products))

	median := 0.0
	if len(products)%2 == 0 {
		median = float64(products[len(products)/2-1]+products[len(products)/2]) / 2.0
	} else {
		median = float64(products[len(products)/2])
	}

	return fmt.Sprintf("Range: %d Average: %.2f Median: %.2f", rangeValue, average, median)
}

func TestPart(t *testing.T) {
	assert.Equal(t, "Range: 0 Average: 1.00 Median: 1.00", Part(1))
	assert.Equal(t, "Range: 1 Average: 1.50 Median: 1.50", Part(2))
	assert.Equal(t, "Range: 2 Average: 2.00 Median: 2.00", Part(3))
	assert.Equal(t, "Range: 3 Average: 2.50 Median: 2.50", Part(4))
	assert.Equal(t, "Range: 5 Average: 3.50 Median: 3.50", Part(5))
}
