package booleanexpr

import (
	"testing"
)

type MockElement struct {
	value bool
}

func (e *MockElement) Check() bool {
	return e.value
}

func TestEvaluate(t *testing.T) {
	elements := map[string]Element{
		"A": &MockElement{value: true},
		"B": &MockElement{value: false},
	}

	tests := []struct {
		input    string
		expected bool
	}{
		{"A", true},
		{"B", false},
		{"!A", false},
		{"!B", true},
		{"A & A", true},
		{"A & B", false},
		{"B & B", false},
		{"A | B", true},
		{"B | B", false},
		{"(A & B) | A", true},
		{"!(A & B) | B", true},
	}

	for _, test := range tests {
		result := Evaluate(test.input, elements)
		if result != test.expected {
			t.Errorf("Evaluate(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}
