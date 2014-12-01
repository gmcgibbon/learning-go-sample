// Download Go for your platform to get started here: http://golang.org/doc/install
// See the official Go documentation here: https://golang.org/doc/
// Check out the Go standard library docs here: https://golang.org/pkg/

// Package main allows your Go application to make an executable
// of the same file name that runs the main function.
package main

// import allows borrowing public functionality of other packages
import (
	format "fmt"                                   // packages can have aliased names
	"github.com/gmcgibbon/learning-go-sample/kata" // packages are referenced by name, not their path
	//"github.com/gmcgibbon/learning-go-sample/lang"
)

// Main is the sample application's main function,
// it will run when the compiled sample app is executed
func main() {

	//lang.DoStuff() // demonstrates control flow of defer, panic, and rescue

	format.Println(kata.WowzaCowza(100)) // functions are called using dot notation
}
