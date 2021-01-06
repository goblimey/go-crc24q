# go-crc24q: A Go implementation of the Qualcomm 24-bit Cyclic Redundancy Checksum (CRC) Algorithm

This is a Go implementation of the Qualcomm CRC-24Q cyclic redundancy checksum,
written by Mark Rafter.

This algorithm is used for many purposes, including the checksum value of RTCM messages.
These are defined by the Radio Technical Commission for Maritime Services (RTCM).
RTCM standard 10403 specifies the Differential GNSS
(Global Navigation Satellite Systems) Services,
currently at version 3.
RTCM3 messages are used to provide corrections to satellite navigation systems such as GPS,
allowing greater accuracy.
Each RTCM3 message is a stream of bits ending in a 24-bit checksum,
created and checked using this algorithm.

The source code contains references to the original paper describing the algorithm.

There are various implementations of the algorithm in C,
including one in [rtklib](http://www.rtklib.com/)
and another in [gpsd](https://github.com/ukyg9e5r6k7gubiekd6/gpsd).

## Download

go get https://github.com/goblimey/go-crc24q

## Usage

```
include (
    crc24q "github.com/goblimey/go-crc24q"
)
```

Given message, a slice of unsigned 8-bit integers ([]uint8):

```
var crc uint32 = crc24q.Hash(message)
```

produces the checksum.

To test that a checksum value in a message is correct,
take the message up to but excluding the checksum,
hash it and compare the result with the hash in the message.

