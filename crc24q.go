package crc24q

// Author: Mark Rafter
// License: MIT License

// This implementation of the crc calculations is a direct implementaion
// of the table lookup algorithm on pages 66 and 67 of the article
//     "A Tutorial on CRC Computations" by
//      Tenkasi V. Ramabadran and Sunil S. Gaitonde
// in the August 1988 issue of IEEE micro.

// This is the Qualcomm CRC-24Q polynomial, used by RTCM104V3.
//
//    x^24 + x^23 + x^18 + x^17 + x^14 + x^11 + x^10 + x^7 + x^6 + x^5 + x^4 + x^3 + x + 1
//

// Go, I would like the following as const :(
var crcPoly = initPoly([]int{24, 23, 18, 17, 14, 11, 10, 7, 6, 5, 4, 3, 1, 0})
var crcDeg = deg(poly(crcPoly))
var crcMask = uint32((1 << crcDeg) - 1)

// Hash calculates the crc24q checksum of data.
//
func Hash(data []byte) uint32 {
	crc := uint32(0)
	for i := range data {
		//
		// We look up precomputed data in table that is
		// based on the current data (data[i])
		// and the hiByte of the current crc.
		//
		// See the paper for an explanation of the next three lines.
		//
		t := table[(data[i] ^ HiByte(crc))]
		crc = crc << bitsInUint8
		crc = crc ^ t
	}
	return crc & crcMask
}

// table contains precomputed values for the crc calculation.
//
// We are doing a byte (uint8) oriented table lookup implementation.
// So table has length 2**bitsInUint8  i.e. (1 << bitsInUint8)
//
var table [tableLen]uint32

const tableLen = (1 << bitsInUint8) // 2**bitsInUint8
const bitsInUint8 = 8               // The number of bits in a uint8

// Initialise the crc table.
//
func init() {

	for i := range table {
		//
		// See the paper for an explanation of the next three lines.
		//
		t := poly(i) << crcDeg
		r := polyRem(t, poly(crcPoly))
		table[i] = uint32(r) & crcMask
	}
}

// LoByte returns the lower byte of a 24-bit CRC value.
func LoByte(x uint32) byte { return uint8(x & 0xff) }

// MiByte returns the middle byte of a 24-bit CRC value.
func MiByte(x uint32) byte { return uint8((x >> 8) & 0xff) }

// HiByte returns the high byte of a 24-bit CRC value.
func HiByte(x uint32) byte { return uint8((x >> 16) & 0xff) }
