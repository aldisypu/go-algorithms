package kata

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/56541980fa08ab47a0000040

func PrinterError(s string) string {
	errorCount := 0
	for _, char := range s {
		if char < 'a' || char > 'm' {
			errorCount++
		}
	}

	return fmt.Sprintf("%d/%d", errorCount, len(s))
}

func TestPrinterError(t *testing.T) {
	assert.Equal(t, "3/56", PrinterError("aaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyz"))
	assert.Equal(t, "6/60", PrinterError("kkkwwwaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyz"))
	assert.Equal(t, "11/65", PrinterError("kkkwwwaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyzuuuuu"))
}
