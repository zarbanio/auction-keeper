package math

import (
	"math/big"
	"testing"
)

func TestNormalize(t *testing.T) {
	testCases := []struct {
		input     *big.Int
		decimals  int64
		expectStr string
	}{
		{
			input:     BigIntFromString("123456789"),
			decimals:  8,
			expectStr: "1.23456789",
		},
		{
			input:     BigIntFromString("123456789"),
			decimals:  1,
			expectStr: "12345678.9",
		},
		{
			input:     BigIntFromString("123456700"),
			decimals:  1,
			expectStr: "12345670",
		},
		{
			input:     BigIntFromString("1234"),
			decimals:  4,
			expectStr: "0.1234",
		},
		{
			input:     BigIntFromString("1234"),
			decimals:  8,
			expectStr: "0.00001234",
		},
		{
			input:     BigIntFromString("123456789012345678901234"),
			decimals:  24,
			expectStr: "0.123456789012345678901234",
		},
		{
			input:     BigIntFromString("1234567890123456789012345"),
			decimals:  25,
			expectStr: "0.1234567890123456789012345",
		},
		{
			input:     BigIntFromString("12345678901234567890123456789012345678901234567890"),
			decimals:  24,
			expectStr: "12345678901234567890123456.78901234567890123456789",
		},
		{
			input:     BigIntFromString("12345678901234567890123456789012345678901234567899"),
			decimals:  24,
			expectStr: "12345678901234567890123456.789012345678901234567899",
		},
	}

	for _, tc := range testCases {
		result := Normalize(tc.input, tc.decimals)

		resultStr := result.String()
		if resultStr != tc.expectStr {
			t.Errorf("For input %s and decimals %d, expected %s, but got %s", tc.input, tc.decimals, tc.expectStr, resultStr)
		}
	}
}
