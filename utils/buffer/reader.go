package buffer

// ReadAsUint64 reads an uint64 from r and stores the result into c with pointer type casting into type T.
func ReadAsUint64[T any](r Reader, c *T) (n int64, err error) {
	_ = "STUB: not implemented"
	/* #nosec G103 -- behavior and consequences well understood, pointer type cast */ return 0, nil
}

// ReadAsUint32 reads an uint32 from r and stores the result into c with pointer type casting into type T.
func ReadAsUint32[T any](r Reader, c *T) (n int64, err error) {
	_ = "STUB: not implemented"
	/* #nosec G103 -- behavior and consequences well understood, pointer type cast */ return 0, nil
}

// ReadAsUint16 reads an uint16 from r and stores the result into c with pointer type casting into type T.
func ReadAsUint16[T any](r Reader, c *T) (n int64, err error) {
	_ = "STUB: not implemented"
	/* #nosec G103 -- behavior and consequences well understood, pointer type cast */ return 0, nil
}

// ReadAsUint8 reads an uint8 from r and stores the result into c with pointer type casting into type T.
func ReadAsUint8[T any](r Reader, c *T) (n int64, err error) {
	_ = "STUB: not implemented"
	/* #nosec G103 -- behavior and consequences well understood, pointer type cast */ return 0, nil
}

// ReadAsUint64Slice reads a slice of uint64 from r and stores the result into c with pointer type casting into type T.
func ReadAsUint64Slice[T any](r Reader, c []T) (n int64, err error) {
	_ = "STUB: not implemented"
	/* #nosec G103 -- behavior and consequences well understood, pointer type cast */ return 0, nil
}

// ReadAsUint32Slice reads a slice of uint32 from r and stores the result into c with pointer type casting into type T.
func ReadAsUint32Slice[T any](r Reader, c []T) (n int64, err error) {
	_ = "STUB: not implemented"
	/* #nosec G103 -- behavior and consequences well understood, pointer type cast */ return 0, nil
}

// ReadAsUint16Slice reads a slice of uint16 from r and stores the result into c with pointer type casting into type T.
func ReadAsUint16Slice[T any](r Reader, c []T) (n int64, err error) {
	_ = "STUB: not implemented"
	/* #nosec G103 -- behavior and consequences well understood, pointer type cast */ return 0, nil
}

// ReadAsUint8Slice reads a slice of uint8 from r and stores the result into c with pointer type casting into type T.
func ReadAsUint8Slice[T any](r Reader, c []T) (n int64, err error) {
	_ = "STUB: not implemented"
	/* #nosec G103 -- behavior and consequences well understood, pointer type cast */ return 0, nil
}

// Read reads a slice of bytes from r and copies it on c.
func Read(r Reader, c []byte) (n int64, err error) { _ = "STUB: not implemented"; return 0, nil }

// ReadUint8 reads a byte from r and stores the result into *c.
func ReadUint8(r Reader, c *uint8) (n int64, err error) { _ = "STUB: not implemented"; return 0, nil }

// Reads one byte

// ReadUint8Slice reads a slice of byte from r and stores the result into c.
func ReadUint8Slice(r Reader, c []uint8) (n int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// ReadUint16 reads a uint16 from r and stores the result into *c.
func ReadUint16(r Reader, c *uint16) (n int64, err error) { _ = "STUB: not implemented"; return 0, nil }

// Reads one byte

// ReadUint16Slice reads a slice of uint16 from r and stores the result into c.
func ReadUint16Slice(r Reader, c []uint16) (n int64, err error) {
	_ = "STUB: not implemented"

	// c is empty, return
	return 0, nil
}

// Then returns the written bytes

// If the slice to write on is equal or smaller than the amount peaked

// Discards what was read

// Decodes the maximum

// Discard what was peeked

// Recurses on the remaining slice to fill

// ReadUint32 reads a uint32 from r and stores the result into *c.
func ReadUint32(r Reader, c *uint32) (n int64, err error) { _ = "STUB: not implemented"; return 0, nil }

// Reads one byte

// ReadUint32Slice reads a slice of uint32 from r and stores the result into c.
func ReadUint32Slice(r Reader, c []uint32) (n int64, err error) {
	_ = "STUB: not implemented"

	// c is empty, return
	return 0, nil
}

// Avoid EOF

// Then returns the written bytes

// If the slice to write on is equal or smaller than the amount peaked

// Discards what was read

// Decodes the maximum

// Discard what was peeked

// Recurses on the remaining slice to fill

// ReadUint64 reads a uint64 from r and stores the result into c.
func ReadUint64(r Reader, c *uint64) (n int64, err error) { _ = "STUB: not implemented"; return 0, nil }

// Reads one byte

// ReadUint64Slice reads a slice of uint64 from r and stores the result into c.
func ReadUint64Slice(r Reader, c []uint64) (n int64, err error) {
	_ = "STUB: not implemented"

	// c is empty, return
	return 0, nil
}

// Avoid EOF

// Then returns the written bytes

// If the slice to write on is equal or smaller than the amount peaked

// Discards what was read

// Decodes the maximum

// Discard what was peeked

// Recurses on the remaining slice to fill
