package crc24q

// Author: Mark Rafter
// License: MIT License

import "testing"

func TestPolyDivRem(test *testing.T) {

	for t := poly(0); t < 10; t++ {
		for b := poly(1); b < 10; b++ {

			q, r := polyDivRem(poly(t), poly(b))

			m := polyMul(q, b) ^ r
			if m != t {
				test.Fatal("after division with remainder  t != q*b + r")
			}
		}
	}
}
