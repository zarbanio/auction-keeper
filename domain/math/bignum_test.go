package math

import (
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
	"time"
)

func TestPow10(t *testing.T) {
	assert.Equal(t, math.BigPow(10, 0), Pow10(0))
	assert.Equal(t, math.BigPow(10, 1), Pow10(1))
	assert.Equal(t, math.BigPow(10, 27), Pow10(27))
	s1 := time.Now()
	Pow10(26)
	s2 := time.Now()
	Pow10(26)
	s3 := time.Now()
	assert.Equal(t, math.BigPow(10, 26), Pow10(26))
	assert.True(t, s3.Sub(s2).Nanoseconds() < s2.Sub(s1).Nanoseconds())
}

func TestNormalize(t *testing.T) {
	assert.Equal(t, "0.01234567", Normalize(big.NewInt(123456789), 10).String())
	assert.Equal(t, "0.12345678", Normalize(big.NewInt(123456789), 9).String())
	assert.Equal(t, "12345.6789", Normalize(big.NewInt(123456789), 4).String())
	assert.Equal(t, "123456.789", Normalize(big.NewInt(123456789), 3).String())
	assert.Equal(t, "1234567.89", Normalize(big.NewInt(123456789), 2).String())
	assert.Equal(t, "12345678.9", Normalize(big.NewInt(123456789), 1).String())
	assert.Equal(t, "123456789", Normalize(big.NewInt(123456789), 0).String())
}
