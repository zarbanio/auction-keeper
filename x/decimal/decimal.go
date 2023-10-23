package decimal

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"

	"math/big"

	"github.com/cockroachdb/apd"
)

var ErrDivisionByZero = errors.New("division by zero")

type Decimal struct {
	value *apd.Decimal
	ctx   *apd.Context
}

var ctx = apd.BaseContext.WithPrecision(50)
var Zero = NewFromInt(0)

func NewFromInt(n int64) Decimal {
	return Decimal{
		value: apd.New(n, 0),
		ctx:   ctx,
	}
}

func NewFromBigInt(n *big.Int) Decimal {
	return Decimal{
		value: apd.NewWithBigInt(n, 0),
		ctx:   ctx,
	}
}

func NewFromFloat64(f float64) (Decimal, error) {
	d, err := new(apd.Decimal).SetFloat64(f)
	if err != nil {
		return Decimal{}, err
	}
	return Decimal{value: d, ctx: ctx}, nil
}
func NewFromFloat(f float64) Decimal {
	d, _ := NewFromFloat64(f)
	return d
}

func NewFromString(s string) (Decimal, error) {
	value, cond, err := apd.NewFromString(s)
	if err != nil {
		return Decimal{}, err
	}
	if cond != apd.Inexact {
		return Decimal{value: value, ctx: ctx}, nil
	}
	return Decimal{}, fmt.Errorf("can't convert %s to decimal: fractional part too long", s)
}

func (d Decimal) Cmp(d2 Decimal) int {
	return d.value.Cmp(d2.value)
}

// Equal returns whether the numbers represented by d and d2 are equal.
func (d Decimal) Equal(d2 Decimal) bool {
	return d.Cmp(d2) == 0
}

// Equals is deprecated, please use Equal method instead
func (d Decimal) Equals(d2 Decimal) bool {
	return d.Equal(d2)
}

// GreaterThan (GT) returns true when d is greater than d2.
func (d Decimal) GreaterThan(d2 Decimal) bool {
	return d.Cmp(d2) == 1
}

// GreaterThanOrEqual (GTE) returns true when d is greater than or equal to d2.
func (d Decimal) GreaterThanOrEqual(d2 Decimal) bool {
	cmp := d.Cmp(d2)
	return cmp == 1 || cmp == 0
}

// LessThan (LT) returns true when d is less than d2.
func (d Decimal) LessThan(d2 Decimal) bool {
	return d.Cmp(d2) == -1
}

// LessThanOrEqual (LTE) returns true when d is less than or equal to d2.
func (d Decimal) LessThanOrEqual(d2 Decimal) bool {
	cmp := d.Cmp(d2)
	return cmp == -1 || cmp == 0
}

func (d Decimal) Sign() int {
	if d.value.Cmp(Zero.value) == 0 {
		return 0
	} else if d.value.Sign() == -1 {
		return -1
	}
	return 1
}

// IsPositive return
//
//	true if d > 0
//	false if d == 0
//	false if d < 0
func (d Decimal) IsPositive() bool {
	return d.Sign() == 1
}

// IsNegative return
//
//	true if d < 0
//	false if d == 0
//	false if d > 0
func (d Decimal) IsNegative() bool {
	return d.Sign() == -1
}

func (d Decimal) Neg() Decimal {
	val := new(big.Int).Neg(d.BigInt())
	return Decimal{
		value: apd.NewWithBigInt(val, 0),
	}
}

// BigInt returns integer component of the decimal as a BigInt.
func (d Decimal) BigInt() *big.Int {
	i := &big.Int{}
	i.SetString(d.value.String(), 10)
	return i
}

// Min returns the smallest Decimal that was passed in the arguments.
//
// To call this function with an array, you must do:
//
//	Min(arr[0], arr[1:]...)
//
// This makes it harder to accidentally call Min with 0 arguments.
func Min(first Decimal, rest ...Decimal) Decimal {
	ans := first
	for _, item := range rest {
		if item.Cmp(ans) < 0 {
			ans = item
		}
	}
	return ans
}

// Max returns the largest Decimal that was passed in the arguments.
//
// To call this function with an array, you must do:
//
//	Max(arr[0], arr[1:]...)
//
// This makes it harder to accidentally call Max with 0 arguments.
func Max(first Decimal, rest ...Decimal) Decimal {
	ans := first
	for _, item := range rest {
		if item.Cmp(ans) > 0 {
			ans = item
		}
	}
	return ans
}

// Sum returns the combined total of the provided first and rest Decimals
func Sum(first Decimal, rest ...Decimal) Decimal {
	total := first
	for _, item := range rest {
		total = total.Add(item)
	}

	return total
}

// Avg returns the average value of the provided first and rest Decimals
func Avg(first Decimal, rest ...Decimal) Decimal {
	count := NewFromInt(int64(len(rest) + 1))
	sum := Sum(first, rest...)
	return sum.Div(count)
}

// IsZero return
//
//	true if d == 0
//	false if d > 0
//	false if d < 0
func (d Decimal) IsZero() bool {
	return d.Sign() == 0
}

func MustNewFromString(s string) Decimal {
	d, _, err := apd.NewFromString(s)
	if err != nil {
		panic(err)
	}
	return Decimal{
		value: d,
		ctx:   ctx,
	}
}

func (d Decimal) Add(b Decimal) Decimal {
	c := &apd.Decimal{}
	_, err := d.ctx.Add(c, d.value, b.value)
	if err != nil {
		panic(err)
	}
	return Decimal{value: c, ctx: d.ctx}
}

func (d Decimal) Div(b Decimal) Decimal {
	c := &apd.Decimal{}
	cond, err := d.ctx.Quo(c, d.value, b.value)
	if err != nil {
		panic(err)
	}
	if cond == apd.DivisionByZero {
		panic(ErrDivisionByZero)
	}
	return Decimal{value: c, ctx: d.ctx}
}

func (d Decimal) Sub(b Decimal) Decimal {
	c := &apd.Decimal{}
	_, err := d.ctx.Sub(c, d.value, b.value)
	if err != nil {
		panic(err)
	}
	return Decimal{value: c, ctx: d.ctx}
}

func (d Decimal) Mul(b Decimal) Decimal {
	c := &apd.Decimal{}
	_, err := d.ctx.Mul(c, d.value, b.value)
	if err != nil {
		panic(err)
	}
	return Decimal{value: c, ctx: d.ctx}
}

func (d Decimal) String() string {
	return d.value.Text('f')
}

func (d *Decimal) Scan(value interface{}) error {
	if value == nil {
		return errors.New("Scan received nil")
	}
	switch v := value.(type) {
	case []byte:
		decValue, err := NewFromString(string(v))
		if err != nil {
			return err
		}
		d.value = decValue.value
		d.ctx = decValue.ctx
		return nil
	default:
		return fmt.Errorf("unsupported scan type: %T", value)
	}
}

func (d Decimal) Value() (driver.Value, error) {
	if d.value == nil {
		return nil, nil
	}
	return d.String(), nil
}

func (d Decimal) MarshalJSON() ([]byte, error) {
	strValue := d.String()
	jsonValue := "\"" + strValue + "\""
	return []byte(jsonValue), nil
}

func (d *Decimal) UnmarshalJSON(b []byte) error {
	strValue := strings.Trim(string(b), "\"")
	decValue, err := NewFromString(strValue)
	if err != nil {
		return err
	}
	d.value = decValue.value
	d.ctx = decValue.ctx
	return nil
}
