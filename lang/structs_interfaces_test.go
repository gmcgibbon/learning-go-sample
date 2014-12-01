// Go language fundamentals.
// Structures and Interfaces unit tests
package lang

import "testing"

func TestInterface(t *testing.T) {

	subject := WriteCode(new(ComputerScienceStudent))
	expected := "*click* *click* *click* "

	if subject != expected {
		t.Errorf("Expected %s got %s", expected, subject)
	}
}

func TestEmbedding(t *testing.T) {

	subject := NewStudent("Gannon", 20, 1234567)
	expected := &Student{Person{"Gannon", 20}, 1234567} // & returns a pointer to type

	if subject.Name != expected.Name ||
		subject.Age != expected.Age ||
		subject.Number != expected.Number {

		t.Errorf("Expected %v got %v", expected, subject)
	}
}
