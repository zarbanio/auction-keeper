package math

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestWadMul(t *testing.T) {
	type test struct {
		a        *big.Int
		b        *big.Int
		expected string
	}
	tests := []test{
		{
			new(big.Int).Mul(big.NewInt(5), Pow10(18)),
			new(big.Int).Mul(big.NewInt(6), Pow10(18)),
			"30000000000000000000",
		},
		{
			new(big.Int).Mul(big.NewInt(7), Pow10(9)),
			new(big.Int).Mul(big.NewInt(8), Pow10(9)),
			"56",
		},
	}
	for _, tt := range tests {
		c := WadMul(tt.a, tt.b)
		assert.Equal(t, tt.expected, c.String())
	}
}

func TestWadDiv(t *testing.T) {
	type test struct {
		a        *big.Int
		b        *big.Int
		expected string
	}
	tests := []test{
		{
			new(big.Int).Mul(big.NewInt(5), Pow10(18)),
			new(big.Int).Mul(big.NewInt(6), Pow10(18)),
			"833333333333333333",
		},
		{
			new(big.Int).Mul(big.NewInt(7), Pow10(9)),
			new(big.Int).Mul(big.NewInt(8), Pow10(9)),
			"875000000000000000",
		},
	}
	for _, tt := range tests {
		c := WadDiv(tt.a, tt.b)
		assert.Equal(t, tt.expected, c.String())
	}
}

func TestRayMul(t *testing.T) {
	type test struct {
		a        *big.Int
		b        *big.Int
		expected string
	}
	tests := []test{
		{
			new(big.Int).Mul(big.NewInt(5), Pow10(18)),
			new(big.Int).Mul(big.NewInt(6), Pow10(18)),
			"30000000000",
		},
		{
			new(big.Int).Mul(big.NewInt(7), Pow10(9)),
			new(big.Int).Mul(big.NewInt(8), Pow10(9)),
			"0",
		},
	}
	for _, tt := range tests {
		c := RayMul(tt.a, tt.b)
		assert.Equal(t, tt.expected, c.String())
	}
}

func TestRayDiv(t *testing.T) {
	type test struct {
		a        *big.Int
		b        *big.Int
		expected string
	}
	tests := []test{
		{
			new(big.Int).Mul(big.NewInt(5), Pow10(18)),
			new(big.Int).Mul(big.NewInt(6), Pow10(18)),
			"833333333333333333333333333",
		},
		{
			new(big.Int).Mul(big.NewInt(7), Pow10(9)),
			new(big.Int).Mul(big.NewInt(8), Pow10(9)),
			"875000000000000000000000000",
		},
	}
	for _, tt := range tests {
		c := RayDiv(tt.a, tt.b)
		assert.Equal(t, tt.expected, c.String())
	}
}

func TestWadToRay(t *testing.T) {
	type test struct {
		a        *big.Int
		expected string
	}
	tests := []test{
		{
			new(big.Int).Mul(big.NewInt(73813893131), Pow10(9)),
			"73813893131000000000000000000",
		},
		{
			new(big.Int).Mul(big.NewInt(3617836817), Pow10(6)),
			"3617836817000000000000000",
		},
	}
	for _, tt := range tests {
		c := WadToRay(tt.a)
		assert.Equal(t, tt.expected, c.String())
	}
}

func TestRayToWad(t *testing.T) {
	type test struct {
		a        *big.Int
		expected string
	}
	tests := []test{
		{
			new(big.Int).Mul(big.NewInt(19128901), Pow10(6)),
			"19129",
		},
		{
			new(big.Int).Mul(big.NewInt(638163891), Pow10(23)),
			"63816389100000000000000",
		},
	}
	for _, tt := range tests {
		c := RayToWad(tt.a)
		assert.Equal(t, tt.expected, c.String())
	}
}

func TestRayPow(t *testing.T) {
	type test struct {
		a        *big.Int
		b        *big.Int
		expected string
	}
	tests := []test{
		{
			new(big.Int).Mul(big.NewInt(761), Pow10(3)),
			big.NewInt(4),
			"1000000000000000000003044000",
		},
		{
			new(big.Int).Mul(big.NewInt(941), Pow10(8)),
			big.NewInt(78),
			"1000000000000007339800000000",
		},
		{
			Ray,
			Ray,
			"166666666666666666666666666666666666666666666666666667500000000000000000000000001000000000000000000000000000",
		},
		{
			Ray,
			Wad,
			"166666666666666666666666666666666667500000000000000001000000000000000000000000000",
		},
	}
	for _, tt := range tests {
		c := BinomialApproximatedRayPow(tt.a, tt.b)
		assert.Equal(t, tt.expected, c.String())
	}
}
