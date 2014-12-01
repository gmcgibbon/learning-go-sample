// Go language fundamentals.
// Functions and Closures
package lang

// privateFunc is visible only inside the lang package
// due to the first letter of its name being lowercase
func privateFunc() string {

	// private access, not available outside lang
	return "I'm private!"
}

// PublicFunc is visible inside and outside the lang package
// due to the first letter of its name being uppercase
func PublicFunc() string {

	// public access, available inside lang
	return "I'm public!"
}

// Multiplier multiplies each number passed together
// and returns the product
func Multiplier(nums ...int) int {

	product := 1

	for _, num := range nums { // range is like foreach in other languages

		product *= num
	}

	return product
}

// GetIncrementingClosure retruns a closure that returns
// an incremented integer value when invoked
func GetIncrementingClosure() func() int {

	// i is used inside the returned anonymous function,
	// the state of i is therefore persisted when the returned
	// function is called
	var i int

	return func() int {

		i += 1

		return i
	}
}
