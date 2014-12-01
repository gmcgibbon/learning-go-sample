// Go language fundamentals.
// Panics and Defers
package lang

import "fmt"

// Prints first action
func firstDoThis() {

	fmt.Println("First")
}

// Prints second action or panics
func thenDoThis(shouldPanic bool) {

	if shouldPanic {
		panic("A problem occurred!")
	}

	fmt.Println("Second") // will never be called if panicing
}

// Prints third action and recovers from the panic
func finallyThis() {

	recover() // when called in a deferred func the panic is handled

	fmt.Println("Third")
}

// DoStuff prints out an example  panicked and recovered control flow
func DoStuff() {

	firstDoThis() // print first action

	defer finallyThis() // defer to call afer DoStuff is done executing

	thenDoThis(true) // called before finally due to the defer keyowrd
}
