package kata

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/5556282156230d0e5e000089

func DNAtoRNA(dna string) string {
	return strings.ReplaceAll(dna, "T", "U")
}

func TestDNAtoRNA(t *testing.T) {
	assert.Equal(t, "GCAU", DNAtoRNA("GCAT"))
}
