package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/50ee6b0bdeab583673000025

func Namevar() string {
	var a string = "code"
	var b string = "wa.rs"
	var name string = a + b
	return name
}

func TestNamevar(t *testing.T) {
	assert.Equal(t, "codewa.rs", Namevar())
}
