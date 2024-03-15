package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/568dcc3c7f12767a62000038

func SetAlarm(employed, vacation bool) bool {
	return employed && !vacation
}

func TestSetAlarm(t *testing.T) {
	assert.False(t, SetAlarm(true, true))
	assert.False(t, SetAlarm(false, true))
	assert.False(t, SetAlarm(false, false))
	assert.True(t, SetAlarm(true, false))
}
