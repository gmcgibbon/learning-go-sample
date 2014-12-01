// Package kata provides simple "coding kata" functions.
// You will find some coding kata functions (WowzaCowza and Factorial) here.
package kata

import (
	"strconv"
)

// WowzaCowza prints out a number a specified amount of times and returns a
// \n delimited string with substituted values for multiples of 3 and 5
func WowzaCowza(times int) string {

	lines := "" // initialize lines to an empty string, shorthand style

	for i := 1; i < times+1; i++ {

		// word is initialized to a string's default value, an empty string
		var word string

		if i%3 == 0 && i%5 == 0 {

			word = "WowzaCowza"
		} else if i%3 == 0 {

			word += "Wowza"
		} else if i%5 == 0 {

			word += "Cowza"
		} else {

			// use the Itoa func from strconv package to
			// convert an int to a string
			word = strconv.Itoa(i)
		}

		lines += word + "\n" // append a newline to word and add to lines
	}

	return lines
}

// Factorial calculates the factorial of a given number
func Factorial(number int) int {

	var factorial int // default 0

	if number < 2 {

		factorial = 1
	} else {

		factorial = number * Factorial(number-1) // recursively calc factorial
	}

	return factorial
}
