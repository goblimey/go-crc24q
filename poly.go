package crc24q

// Author: Mark Rafter
// License: MIT License

// Below is some code for calculation of polynomial remainders.
// It been written to be a simple to understand, direct
// implementation of the mathematics.
//
// It's written for clarity, not speed - but that doesn't matter as it's only
// used in initialising a small table at program startup.

// We are dealling with polynomials who's coefficients
// are just 0 or 1 - i.e. binary polynomials.
//
// We can represent a binary polynomial using the bits
// of an unsigned integer.  That way polynomial
// addition and subtraction are just xor, and polynomial
// multiplication by a power of x is just a left shift.
//
// Implementation restriction:
// only implement poynomials that can be represented in 32 bits.
//
// This is (just) enough for the crc24q table initialisation.
//
type poly uint32

// initPoly initialises a polnomial from its coefficients.
//
func initPoly(coeffs []int) (res poly) {
	for _, b := range coeffs {
		res = res | poly(1<<b)
	}
	return res
}

// polyRem returns the remainder of the binary polynomial division t/b.
//
func polyRem(t, b poly) (res poly) {
	_, res = polyDivRem(t, b)
	return res
}

// polyDivRem returns the quotient and remainder of the binary polynomial division t/b.
// after this we will have   t == b*q + r
//
func polyDivRem(t, b poly) (q, r poly) {
	switch {
	case b == 0:
		panic("Poly Div by Zero")
	case t == 0:
		return 0, 0
	case deg(t) < deg(b):
		return 0, t
	default:
		// at this point we know that b and t are non zero
		// so we are safe to take their degrees.
		// we also know that q is non zero.

		// align b prior to subtraction.
		align := deg(t) - deg(b)
		t = t ^ (b << align) // subtract shifted (multiplied) form of b.
		q = 1 << align       // q is now the number by which we have multiplied b before subtracting from t.

		q2, r2 := polyDivRem(t, b)

		return q2 ^ q, r2 // don't forget to add q to q2.
	}
}

// polyMul(x,y) returns the product of x and y.
//
func polyMul(x, y poly) (res poly) {
	for y != 0 {
		if (y & 1) != 0 {
			res = res ^ x
		}
		y = y >> 1
		x = x << 1
	}
	return res
}

// deg(p) returns the degree of the polynomial,
// aka the most significant bit set in p when
// p is viewed simply as an unsigned.
//
func deg(p poly) int {
	switch p {
	case 0:
		panic("You can't take the degree of the zero polynomial")
	case 1:
		return 0
	default:
		return 1 + deg(p>>1)
	}
}
