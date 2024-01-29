package kata

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/554e4a2f232cdd87d9000038

func DNAStrand(dna string) string {
	replacer := strings.NewReplacer("A", "T", "T", "A", "C", "G", "G", "C")
	return replacer.Replace(dna)
}

func TestDNAStrand(t *testing.T) {
	assert.Equal(t, "TTTT", DNAStrand("AAAA"))
	assert.Equal(t, "TAACG", DNAStrand("ATTGC"))
	assert.Equal(t, "CATA", DNAStrand("GTAT"))
}
