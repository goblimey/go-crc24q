# go-crc24q: A Go implementation of the Qualcomm 24-bit Cyclic Redundancy Checksum (CRC) Algorithm

This is a Go implementation of the Qualcomm CRC-24Q cyclic redundancy checksum,
written by Mark Rafter.

This algorithm is used for many purposes, including the checksum value of RTCM messages.

(RTCM messages are defined by the Radio Technical Commission for Maritime Services (RTCM).
RTCM standard 10403 specifies the Differential GNSS
(Global Navigation Satellite Systems) Services,
currently at version 3.
RTCM3 messages are used to provide corrections to satellite navigation systems such as GPS,
allowing greater accuracy.
Each RTCM3 message is a stream of bits ending in a 24-bit checksum,
created and checked using this algorithm.)

The source code contains references to the original paper describing the algorithm.

There are various implementations of the algorithm in C,
including one in [rtklib](http://www.rtklib.com/)
and another in [gpsd](https://github.com/ukyg9e5r6k7gubiekd6/gpsd).

## Download

go get github.com/goblimey/go-crc24q

## Usage

```
include (
    "github.com/goblimey/go-crc24q/crc24q"
)
```

Given message, []byte, produce a 24-bit checksum like so:

```
crc := crc24q.Hash(message)
```

To check that a message with a checksum at the end is valid,
take the message up to but not including the checksum,
hash it and compare the result with the given checksum.

## Package crc24q
import "github.com/goblimey/go-crc24q/crc24q"

### Index
```
func Hash(data []byte) uint32
func HiByte(x uint32) byte
func LoByte(x uint32) byte
func MiByte(x uint32) byte
```
### Package files
crc24q.go poly.go

#### func Hash
```
func Hash(data []byte) uint32
Hash calculates the crc24q checksum of data.
```

#### func HiByte
```
func HiByte(x uint32) byte
HiByte returns the high byte of a 24-bit CRC value.
```

#### func LoByte
```
func LoByte(x uint32) byte
LoByte returns the lower byte of a 24-bit CRC value.
```

#### func MiByte
```
func MiByte(x uint32) byte
MiByte returns the middle byte of a 24-bit CRC value.
```
