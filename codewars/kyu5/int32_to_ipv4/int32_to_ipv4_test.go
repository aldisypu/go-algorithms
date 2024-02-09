package kata

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/52e88b39ffb6ac53a400022e

func Int32ToIp(n uint32) string {
	ip := net.IPv4(byte(n>>24), byte(n>>16), byte(n>>8), byte(n))
	return ip.String()
}

func TestInt32ToIp(t *testing.T) {
	assert.Equal(t, "128.114.17.104", Int32ToIp(2154959208))
	assert.Equal(t, "128.32.10.1", Int32ToIp(2149583361))
	assert.Equal(t, "0.0.0.0", Int32ToIp(0))
}
