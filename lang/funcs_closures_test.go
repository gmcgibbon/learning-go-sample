// Go language fundamentals.
// Functions and Closures
package lang

import "testing"

func TestMultiplier(t *testing.T) {

	subject := Multiplier(5, 5, 10, 2)
	expected := 500

	if subject != expected {
		t.Errorf("Expected %d got %d", expected, subject)
	}
}

func TestGetIncrementingClosure(t *testing.T) {

	inc := GetIncrementingClosure()

	subject := [...]int{inc(), inc(), inc(), inc()}
	expected := [...]int{1, 2, 3, 4}

	if subject != expected {
		t.Errorf("Expected %s got %s", expected, subject)
	}
}
