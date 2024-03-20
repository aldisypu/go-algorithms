package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/515e188a311df01cba000003

func GetPlanetName(ID int) string {
	switch ID {
	case 1:
		return "Mercury"
	case 2:
		return "Venus"
	case 3:
		return "Earth"
	case 4:
		return "Mars"
	case 5:
		return "Jupiter"
	case 6:
		return "Saturn"
	case 7:
		return "Uranus"
	case 8:
		return "Neptune"
	default:
		return "Pluto"
	}
}

func TestGetPlanetName(t *testing.T) {
	assert.Equal(t, "Venus", GetPlanetName(2))
	assert.Equal(t, "Jupiter", GetPlanetName(5))
	assert.Equal(t, "Earth", GetPlanetName(3))
	assert.Equal(t, "Mars", GetPlanetName(4))
	assert.Equal(t, "Neptune", GetPlanetName(8))
	assert.Equal(t, "Mercury", GetPlanetName(1))
}
