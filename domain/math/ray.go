package math

import (
	"math/big"
)

var (
	MinusOne = big.NewInt(-1)
	Zero     = big.NewInt(0)

	One      = big.NewInt(1)
	OneFloat = new(big.Float).SetInt(One)
	Two      = big.NewInt(2)
	Five     = big.NewInt(5)
	Six      = big.NewInt(6)
	Ten      = big.NewInt(10)

	Wad         = Pow10(18)
	WadDecimals = 18
	HalfWad     = new(big.Int).Div(Wad, Two)

	Ray     = Pow10(27)
	HalfRay = new(big.Int).Div(Ray, Two)

	RayDecimals = 27

	WadRayRatio = Pow10(9)

	SecondsPerYear       = big.NewInt(31536000)
	SecondsPerYearFloat  = new(big.Float).SetInt(SecondsPerYear)
	LoanToValuePrecision = 4

	Rad         = Pow10(45)
	RadDecimals = 45
)

func WadMul(a, b *big.Int) *big.Int {
	mul := new(big.Int).Mul(a, b)
	sum := new(big.Int).Add(HalfWad, mul)
	return new(big.Int).Div(sum, Wad)
}

func WadDiv(a, b *big.Int) *big.Int {
	halfB := new(big.Int).Div(b, big.NewInt(2))
	return new(big.Int).
		Div(
			new(big.Int).Add(
				halfB,
				new(big.Int).Mul(a, Wad)),
			b)

}

func RayMul(a, b *big.Int) *big.Int {
	mul := new(big.Int).Mul(a, b)
	sum := new(big.Int).Add(HalfRay, mul)
	return new(big.Int).Div(sum, Ray)
}

func RayDiv(a, b *big.Int) *big.Int {
	halfB := new(big.Int).Div(b, big.NewInt(2))
	return new(big.Int).
		Div(
			new(big.Int).Add(
				halfB,
				new(big.Int).Mul(a, Ray)),
			b)

}

func RayToWad(a *big.Int) *big.Int {
	halfRatio := new(big.Int).Div(WadRayRatio, big.NewInt(2))
	return new(big.Int).
		Div(
			new(big.Int).Add(halfRatio, a),
			WadRayRatio)
}

func WadToRay(a *big.Int) *big.Int {
	return new(big.Int).Mul(a, WadRayRatio)
}

func RayPow(base, exp *big.Int) *big.Int {
	z := Ray
	if new(big.Int).Mod(exp, big.NewInt(2)).Cmp(big.NewInt(0)) != 0 {
		z = base
	}
	for exp = new(big.Int).Div(exp, big.NewInt(2)); exp.Cmp(big.NewInt(0)) != 0; exp = new(big.Int).Div(exp, big.NewInt(2)) {
		base = RayMul(base, base)
		if new(big.Int).Mod(exp, big.NewInt(2)).Cmp(big.NewInt(0)) != 0 {
			z = RayMul(z, base)
		}
	}
	return z
}

func BinomialApproximatedRayPow(base, exp *big.Int) *big.Int {
	if exp.Cmp(Zero) == 0 {
		return Ray
	}
	expMinusOne := new(big.Int).Sub(exp, One)
	expMinusTwo := new(big.Int).Sub(exp, Two)
	if exp.Cmp(Two) <= 0 {
		expMinusTwo = Zero
	}

	basePowerTwo := RayMul(base, base)
	basePowerThree := RayMul(basePowerTwo, base)

	firstTerm := new(big.Int).Mul(exp, base)

	x1 := new(big.Int).Mul(exp, expMinusOne)
	x2 := new(big.Int).Mul(x1, basePowerTwo)
	secondTerm := new(big.Int).Div(x2, Two)

	x1 = new(big.Int).Mul(exp, expMinusOne)
	x2 = new(big.Int).Mul(x1, expMinusTwo)
	x3 := new(big.Int).Mul(x2, basePowerThree)
	thirdTerm := new(big.Int).Div(x3, Six)

	sum1 := new(big.Int).Add(Ray, firstTerm)
	sum2 := new(big.Int).Add(sum1, secondTerm)
	return new(big.Int).Add(sum2, thirdTerm)
}
