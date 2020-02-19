package fizz_buzz

import (
	"fmt"
)

// if divisible by 3 = Fizz
// if divisible by 5 = Buzz
// if divisible by 3 && 5 = Fizz Buzz
func FizzBuzz(n int) {
	var output string
	for i := 1; i<=n ; i++  {
		switch {
		case i%3 == 0 && i%5 == 0:
			output += "Fizz Buzz"
		case i%3 == 0:
			output += "Fizz"
		case i%5 == 0:
			output += "Buzz"
		default:
			output += fmt.Sprint(i)
		}

		// append comma if not the last number
		if i != n {
			output += ", "
		}
	}

	fmt.Println(output)
}
