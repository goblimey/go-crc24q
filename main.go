// go-crc24q reads a stream of bytes from standard input, appends
// 3 bytes (24 bits) of CRC and writes the result to standard
// output.  The input data should be 2048 bytes or less.  Any
// more is ignored.
//
// The input and output files are binary.  Under UNIX, they
// can be viewed like so:
//
// $ od -A x -t x1z -v {file}
//
// Running the program against the supplied test input like so:
//
//     $ go-crc24q <test_input >result
//
// will display:
//
//     {date} {time} crc e51ed8, HiByte e5, MiByte 1e, LoByte d8
//
// and the file "result" will look like this:
//
//    $ od -A x -t x1z -v result
//    000000 d3 00 aa 46 70 00 66 ff bc a0 00 00 00 04 00 26  >...Fp.f........&<
//    000010 18 00 00 00 20 02 00 00 75 53 fa 82 42 62 9a 80  >.... ...uS..Bb..<
//    000020 00 00 06 95 4e a7 a0 bf 1e 78 7f 0a 10 08 18 7f  >....N....x......<
//    000030 35 04 ab ee 50 77 8a 86 f0 51 f1 4d 82 46 38 29  >5...Pw...Q.M.F8)<
//    000040 0a 8c 35 57 23 87 82 24 2a 01 b5 40 07 eb c5 01  >..5W#..$*..@....<
//    000050 37 a8 80 b3 88 03 23 c4 fc 61 e0 4f 33 c4 73 31  >7.....#..a.O3.s1<
//    000060 cd 90 54 b2 02 70 90 26 0b 42 d0 9c 2b 0c 02 97  >..T..p.&.B..+...<
//    000070 f4 08 3d 9e c7 b2 6e 44 0f 19 48 00 00 00 00 00  >..=...nD..H.....<
//    000080 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  >................<
//    000090 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  >................<
//    0000a0 00 00 00 00 00 00 00 00 00 00 00 00 00 e5 1e d8  >................<
//    0000b0
//
// This is the data from the input file with a three-byte
// checksum value (bytes e5, 1e and d8) added to the end.
//
// The result is an RTCM message frame, composed of a 24-bit
// header, a variable-length message and a 24-bit CRC value.
// The first 8 bits of any RTCM message header is always d3.
// The lower 10 bits of the header give the message length.
// The message follows.  The first 12 bits of the message
// give the message type, in this case type 1127 (BeiDou Full
// Pseudoranges and PhaseRanges plus CNR (high resolution)).
// The 24-bit checksum follows the message.
//
// RTCM messages can be sent over lossy media such as radio
// links and are thus prone to becoming scrambled in transit.
// When a message arrives, the receiver should recalculate the
// CRC value and check that it matches the CRC in the message
// frame.  If so, then it's highly likely that the message has
// survived intact.
//
package main

import (
	"io"
	"log"
	"os"

	"github.com/goblimey/go-crc24q/crc24q"
)

func main() {
	buffer := make([]byte, 2048)
	n, err := io.ReadAtLeast(os.Stdin, buffer, 2048)
	if err != nil && err != io.ErrUnexpectedEOF {
		log.Fatal("error reading input - %v", err)
	}

	buffer = buffer[:n]

	crc := crc24q.Hash(buffer)

	// 0xe5, 0x1e, 0xd8,

	log.Printf("crc %x, HiByte %x, MiByte %x, LoByte %x", crc, crc24q.HiByte(crc),
		crc24q.MiByte(crc), crc24q.LoByte(crc))

	buffer = append(buffer, crc24q.HiByte(crc))
	buffer = append(buffer, crc24q.MiByte(crc))
	buffer = append(buffer, crc24q.LoByte(crc))

	os.Stdout.Write(buffer)
}
