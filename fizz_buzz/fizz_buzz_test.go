package fizz_buzz

import (
	"fmt"
	"testing"
)

func fizzBuzz(total int) {
	for i := 1; i <= total; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("Fizz Buzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}

	}
}

func TestFizzBuzz(t *testing.T) {
	fizzBuzz(15)
}
