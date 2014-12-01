// Package kata provides simple "coding kata" functions.
// Here is where unit tests for this package live.
package kata

import "testing" // import the unit testing package

func TestWowzaCowza(t *testing.T) { // all test funcs must begin with Test and take in a testing type parameter

	subject := WowzaCowza(5)
	expected := "1\n2\nWowza\n4\nCowza\n"

	if subject != expected {

		t.Errorf("Expected \n%s got \n%s", expected, subject)
	}
}

func TestFactorial(t *testing.T) {

	subject := Factorial(5)
	expected := 120

	if subject != expected {

		t.Errorf("Expected %d got %d", expected, subject)
	}
}
