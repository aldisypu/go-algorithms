package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func DigitalRoot(n int) int {
	return (n-1)%9 + 1
}

func TestDigitalRoot(t *testing.T) {
	assert.Equal(t, 7, DigitalRoot(16))
	assert.Equal(t, 6, DigitalRoot(195))
	assert.Equal(t, 2, DigitalRoot(992))
	assert.Equal(t, 9, DigitalRoot(167346))
	assert.Equal(t, 0, DigitalRoot(0))
}
