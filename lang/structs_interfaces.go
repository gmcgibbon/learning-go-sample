// Go language fundamentals.
// Structures and Interfaces
package lang

import "strings"

// Person is a a type representing a human being
type Person struct {
	Name string
	Age  int
}

// Student is a type representing a Person actively learning
type Student struct {
	Person // embed another type anonymously to implement its structure
	Number int64
}

// ComputerScienceStudent is a type representing a Student of Computer Science
type ComputerScienceStudent struct {
	Student
}

// In Go, a "function" becomes a "method" when bound to a type
// Type methods are defined outside the type definition like this:
// (reference type pointer in brackets) FunctionName() returnType {

// By adding this method to the ComputerScienceStudent struct it
// implicitly satisfies the Coder interface without having to declare it

// Type allows a ComputerScienceStudent to type on a keyboard
func (c *ComputerScienceStudent) Type() string {
	return "*click*"
}

// Coder is an interface that represents anything that
// can write code by typing with a keyboard
type Coder interface {
	Type() string
}

// WriteCode takes a coder and gets them typing
func WriteCode(c Coder) string {
	return strings.Repeat(c.Type()+" ", 3)
}

// NewStudent acts a pointer constructor for new Students,
// if it was a value constructor it would be named makeStudent
func NewStudent(name string, age int, number int64) *Student {

	student := new(Student)

	student.Name = name // name and age are directly available due to embedding
	student.Age = age
	student.Number = number

	return student
}
