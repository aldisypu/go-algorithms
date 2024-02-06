package kata

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/52685f7382004e774f0001f7

func HumanReadableTime(seconds int) string {
	duration := time.Second * time.Duration(seconds)
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	seconds = int(duration.Seconds()) % 60

	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}

func TestHumanReadableTime(t *testing.T) {
	assert.Equal(t, "00:00:00", HumanReadableTime(0))
	assert.Equal(t, "00:00:59", HumanReadableTime(59))
	assert.Equal(t, "00:01:00", HumanReadableTime(60))
	assert.Equal(t, "00:01:30", HumanReadableTime(90))
	assert.Equal(t, "00:59:59", HumanReadableTime(3599))
}
