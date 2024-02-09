package kata

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/52223df9e8f98c7aa7000062

func Rot13(msg string) string {
	var result strings.Builder

	for _, char := range msg {
		switch {
		case 'a' <= char && char <= 'z':
			result.WriteByte(byte((char-'a'+13)%26) + 'a')
		case 'A' <= char && char <= 'Z':
			result.WriteByte(byte((char-'A'+13)%26) + 'A')
		default:
			result.WriteRune(char)
		}
	}

	return result.String()
}

func TestRot13(t *testing.T) {
	assert.Equal(t, "ROT13 example.", Rot13("EBG13 rknzcyr."))
	assert.Equal(t, "Ubj pna lbh gryy na rkgebireg sebz na\nvagebireg ng AFN? In the elevators,\nthe extrovert looks at the OTHER guy's shoes.", Rot13("How can you tell an extrovert from an\nintrovert at NSA? Va gur ryringbef,\ngur rkgebireg ybbxf ng gur BGURE thl'f fubrf."))
	assert.Equal(t, "123", Rot13("123"))
	assert.Equal(t, "This is actually the first kata I ever made. Thanks for finishing it! :)", Rot13("Guvf vf npghnyyl gur svefg xngn V rire znqr. Gunaxf sbe svavfuvat vg! :)"))
	assert.Equal(t, "@[`{", Rot13("@[`{"))
}
