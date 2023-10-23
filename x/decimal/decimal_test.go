package decimal

import (
	"testing"
)

func TestDecimal_Add(t *testing.T) {
	tests := []struct {
		name     string
		a, b     string
		expected string
	}{
		{"PositiveAddition", "10", "5", "15"},
		{"NegativeAddition", "-10", "-5", "-15"},
		{"MixedSignAddition", "10", "-5", "5"},
		{"AdditionWithZero", "10", "0", "10"},
		{"ZeroAddition", "0", "10", "10"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a, _ := NewFromString(tt.a)
			b, _ := NewFromString(tt.b)
			expected, _ := NewFromString(tt.expected)

			result := a.Add(b)
			if result.String() != expected.String() {
				t.Errorf("Addition result is incorrect. Got: %s, Expected: %s", result.String(), expected.String())
			}
		})
	}
}
func TestDecimal_Div(t *testing.T) {
	tests := []struct {
		name     string
		a, b     string
		expected error
	}{
		{"PositiveDivision", "10", "2", nil},
		{"NegativeDivision", "-10", "-2", nil},
		{"MixedSignDivision", "10", "-2", nil},
		{"DivisionByZero", "10", "0", ErrDivisionByZero},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a, _ := NewFromString(tt.a)
			b, _ := NewFromString(tt.b)

			if tt.name == "DivisionByZero" {
				defer func() {
					if r := recover(); r != nil {
						if err, ok := r.(error); ok {
							if err.Error() != ErrDivisionByZero.Error() {
								t.Errorf("Expected division by zero error, but got: %v", err)
							}
						} else {
							t.Errorf("Expected an error type, but got: %v", r)
						}
					}
				}()
				a.Div(b)
			} else {
				result := a.Div(b)
				if tt.expected != nil {
					t.Errorf("Expected error: %v, but got result: %s", tt.expected, result.String())
				}
			}
		})
	}
}

func TestDecimal_Sub(t *testing.T) {
	tests := []struct {
		name     string
		a, b     string
		expected string
	}{
		{"PositiveSubtraction", "10", "5", "5"},
		{"NegativeSubtraction", "-10", "-5", "-5"},
		{"MixedSignSubtraction", "10", "-5", "15"},
		{"SubtractionWithZero", "10", "0", "10"},
		{"ZeroSubtraction", "0", "10", "-10"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a, _ := NewFromString(tt.a)
			b, _ := NewFromString(tt.b)
			expected, _ := NewFromString(tt.expected)

			result := a.Sub(b)
			if result.String() != expected.String() {
				t.Errorf("Subtraction result is incorrect. Got: %s, Expected: %s", result.String(), expected.String())
			}
		})
	}
}

func TestDecimal_Mul(t *testing.T) {
	tests := []struct {
		name     string
		a, b     string
		expected string
	}{
		{"PositiveMultiplication", "2", "3", "6"},
		{"NegativeMultiplication", "-2", "-3", "6"},
		{"MixedSignMultiplication", "2", "-3", "-6"},
		{"MultiplicationWithZero", "2", "0", "0"},
		{"MultiplicationByZero", "0", "2", "0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a, _ := NewFromString(tt.a)
			b, _ := NewFromString(tt.b)
			expected, _ := NewFromString(tt.expected)

			result := a.Mul(b)
			if result.String() != expected.String() {
				t.Errorf("Multiplication result is incorrect. Got: %s, Expected: %s", result.String(), expected.String())
			}
		})
	}
}
