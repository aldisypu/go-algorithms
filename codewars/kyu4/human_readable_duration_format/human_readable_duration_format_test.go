package kata

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/52742f58faf5485cae000b9a

func FormatDuration(seconds int64) string {
	if seconds == 0 {
		return "now"
	}

	duration := time.Duration(seconds) * time.Second
	years := int(duration.Hours() / 24 / 365)
	days := int(duration.Hours()/24) % 365
	hours := int(duration.Hours()) % 24
	minutes := int(duration.Minutes()) % 60
	seconds = int64(duration.Seconds()) % 60

	units := []struct {
		value int
		unit  string
	}{
		{years, "year"},
		{days, "day"},
		{hours, "hour"},
		{minutes, "minute"},
		{int(seconds), "second"},
	}

	result := ""
	for _, u := range units {
		if u.value > 0 {
			result += fmt.Sprintf("%d %s", u.value, u.unit)
			if u.value > 1 {
				result += "s"
			}
			result += ", "
		}
	}

	result = result[:len(result)-2]
	if index := lastCommaIndex(result); index != -1 {
		result = result[:index] + " and" + result[index+1:]
	}

	return result
}

func lastCommaIndex(s string) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ',' {
			return i
		}
	}
	return -1
}

func TestFormatDuration(t *testing.T) {
	assert.Equal(t, "1 second", FormatDuration(1))
	assert.Equal(t, "1 minute and 2 seconds", FormatDuration(62))
	assert.Equal(t, "2 minutes", FormatDuration(120))
	assert.Equal(t, "1 hour", FormatDuration(3600))
	assert.Equal(t, "1 hour, 1 minute and 2 seconds", FormatDuration(3662))
}
