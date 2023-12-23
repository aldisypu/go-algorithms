package odd_event

import (
	"fmt"
	"testing"
)

func oddEven(n int) {
	if n%2 != 0 {
		fmt.Printf("%d is Odd Number\n", n)
	} else {
		fmt.Printf("%d is Even Number\n", n)
	}
}

func TestOddEvent(t *testing.T) {
	oddEven(1)
	oddEven(2)
}
