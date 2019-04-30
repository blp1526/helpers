package binaryh

import (
	"encoding/binary"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	PutUvarintKey = "PutUvarint"

	PutVarintKey = "PutVarint"

	ReadKey = "Read"

	ReadUvarintKey = "ReadUvarint"

	ReadVarintKey = "ReadVarint"

	SizeKey = "Size"

	UvarintKey = "Uvarint"

	VarintKey = "Varint"

	WriteKey = "Write"
)

func New() hctx.Map {
	return hctx.Map{

		PutUvarintKey: PutUvarint,

		PutVarintKey: PutVarint,

		ReadKey: Read,

		ReadUvarintKey: ReadUvarint,

		ReadVarintKey: ReadVarint,

		SizeKey: Size,

		UvarintKey: Uvarint,

		VarintKey: Varint,

		WriteKey: Write,
	}
}

// PutUvarint encodes a uint64 into buf and returns the number of bytes written.
// If the buffer is too small, PutUvarint will panic.
var PutUvarint = binary.PutUvarint

// PutVarint encodes an int64 into buf and returns the number of bytes written.
// If the buffer is too small, PutVarint will panic.
var PutVarint = binary.PutVarint

// Read reads structured binary data from r into data.
// Data must be a pointer to a fixed-size value or a slice
// of fixed-size values.
// Bytes read from r are decoded using the specified byte order
// and written to successive fields of the data.
// When decoding boolean values, a zero byte is decoded as false, and
// any other non-zero byte is decoded as true.
// When reading into structs, the field data for fields with
// blank (_) field names is skipped; i.e., blank field names
// may be used for padding.
// When reading into a struct, all non-blank fields must be exported
// or Read may panic.
//
// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Read returns ErrUnexpectedEOF.
var Read = binary.Read

// ReadUvarint reads an encoded unsigned integer from r and returns it as a uint64.
var ReadUvarint = binary.ReadUvarint

// ReadVarint reads an encoded signed integer from r and returns it as an int64.
var ReadVarint = binary.ReadVarint

// Size returns how many bytes Write would generate to encode the value v, which
// must be a fixed-size value or a slice of fixed-size values, or a pointer to such data.
// If v is neither of these, Size returns -1.
var Size = binary.Size

// Uvarint decodes a uint64 from buf and returns that value and the
// number of bytes read (&gt; 0). If an error occurred, the value is 0
// and the number of bytes n is &lt;= 0 meaning:
//
// 	n == 0: buf too small
// 	n  &lt; 0: value larger than 64 bits (overflow)
// 	        and -n is the number of bytes read
var Uvarint = binary.Uvarint

// Varint decodes an int64 from buf and returns that value and the
// number of bytes read (&gt; 0). If an error occurred, the value is 0
// and the number of bytes n is &lt;= 0 with the following meaning:
//
// 	n == 0: buf too small
// 	n  &lt; 0: value larger than 64 bits (overflow)
// 	        and -n is the number of bytes read
var Varint = binary.Varint

// Write writes the binary representation of data into w.
// Data must be a fixed-size value or a slice of fixed-size
// values, or a pointer to such data.
// Boolean values encode as one byte: 1 for true, and 0 for false.
// Bytes written to w are encoded using the specified byte order
// and read from successive fields of the data.
// When writing structs, zero values are written for fields
// with blank (_) field names.
var Write = binary.Write
