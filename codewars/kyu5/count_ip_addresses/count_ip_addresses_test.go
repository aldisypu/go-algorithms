package kata

import (
	"math/big"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/526989a41034285187000de4

func IpsBetween(start, end string) int {
	startIP := ipToBigInt(net.ParseIP(start))
	endIP := ipToBigInt(net.ParseIP(end))

	diff := new(big.Int).Sub(endIP, startIP)

	return int(diff.Int64())
}

func ipToBigInt(ip net.IP) *big.Int {
	ipInt := big.NewInt(0)
	ipInt.SetBytes(ip.To4())
	return ipInt
}

func TestIpsBetween(t *testing.T) {
	assert.Equal(t, 50, IpsBetween("10.0.0.0", "10.0.0.50"))
	assert.Equal(t, 246, IpsBetween("20.0.0.10", "20.0.1.0"))
}
