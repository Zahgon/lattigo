package buffer

// WriteAsUint64 casts &T to an *uint64 and writes it to w.
// User must ensure that T can be stored in an uint64.
func WriteAsUint64[T any](w Writer, c T) (n int64, err error) {
	_ = "STUB: not implemented"
	/* #nosec G103 -- behavior and consequences well understood, pointer type cast */ return 0, nil
}

// WriteAsUint32 casts &T to an *uint32 and writes it to w.
// User must ensure that T can be stored in an uint32.
func WriteAsUint32[T any](w Writer, c T) (n int64, err error) {
	_ = "STUB: not implemented"
	/* #nosec G103 -- behavior and consequences well understood, pointer type cast */ return 0, nil
}

// WriteAsUint16 casts &T to an *uint16 and writes it to w.
// User must ensure that T can be stored in an uint16.
func WriteAsUint16[T any](w Writer, c T) (n int64, err error) {
	_ = "STUB: not implemented"
	/* #nosec G103 -- behavior and consequences well understood, pointer type cast */ return 0, nil
}

// WriteAsUint8 casts &T to an *uint8 and writes it to w.
// User must ensure that T can be stored in an uint8.
func WriteAsUint8[T any](w Writer, c T) (n int64, err error) {
	_ = "STUB: not implemented"
	/* #nosec G103 -- behavior and consequences well understood, pointer type cast */ return 0, nil
}

// WriteAsUint64Slice casts &[]T into *[]uint64 and writes it to w.
// User must ensure that T can be stored in an uint64.
func WriteAsUint64Slice[T any](w Writer, c []T) (n int64, err error) {
	_ = "STUB: not implemented"
	/* #nosec G103 -- behavior and consequences well understood, pointer type cast */ return 0, nil
}

// WriteAsUint32Slice casts &[]T into *[]uint32 and writes it to w.
// User must ensure that T can be stored in an uint32.
func WriteAsUint32Slice[T any](w Writer, c []T) (n int64, err error) {
	_ = "STUB: not implemented"
	/* #nosec G103 -- behavior and consequences well understood, pointer type cast */ return 0, nil
}

// WriteAsUint16Slice casts &[]T into *[]uint16 and writes it to w.
// User must ensure that T can be stored in an uint16.
func WriteAsUint16Slice[T any](w Writer, c []T) (n int64, err error) {
	_ = "STUB: not implemented"
	/* #nosec G103 -- behavior and consequences well understood, pointer type cast */ return 0, nil
}

// WriteAsUint8Slice casts &[]T into *[]uint8 and writes it to w.
// User must ensure that T can be stored in an uint8.
func WriteAsUint8Slice[T any](w Writer, c []T) (n int64, err error) {
	_ = "STUB: not implemented"
	/* #nosec G103 -- behavior and consequences well understood, pointer type cast */ return 0, nil
}

// Write writes a slice of bytes to w.
func Write(w Writer, c []byte) (n int64, err error) { _ = "STUB: not implemented"; return 0, nil }

// WriteUint8 writes a byte c to w.
func WriteUint8(w Writer, c uint8) (n int64, err error) { _ = "STUB: not implemented"; return 0, nil }

// WriteUint8Slice writes a slice of bytes c to w.
func WriteUint8Slice(w Writer, c []uint8) (n int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil

	// Remaining available space in the internal buffer
}

// If there is enough space in the available buffer

// First fills the space

// Flushes

// Then recurses on itself with the remaining slice

// WriteUint16 writes a uint16 c to w.
func WriteUint16(w Writer, c uint16) (n int64, err error) { _ = "STUB: not implemented"; return 0, nil }

// WriteUint16Slice writes a slice of uint16 c to w.
func WriteUint16Slice(w Writer, c []uint16) (n int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil

	// Remaining available space in the internal buffer
}

// If there is enough space in the available buffer

// First fills the space

// Flushes

// Then recurses on itself with the remaining slice

// WriteUint32 writes a uint32 c into w.
func WriteUint32(w Writer, c uint32) (n int64, err error) { _ = "STUB: not implemented"; return 0, nil }

// WriteUint32Slice writes a slice of uint32 c into w.
func WriteUint32Slice(w Writer, c []uint32) (n int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil

	// Remaining available space in the internal buffer
}

// If there is enough space in the available buffer

// First fills the space

// Flushes

// Then recurses on itself with the remaining slice

// WriteUint64 writes a uint64 c into w.
func WriteUint64(w Writer, c uint64) (n int64, err error) { _ = "STUB: not implemented"; return 0, nil }

// WriteUint64Slice writes a slice of uint64 into w.
func WriteUint64Slice(w Writer, c []uint64) (n int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil

	// Remaining available space in the internal buffer
}

// If there is enough space in the available buffer

// First fills the space

// Flushes

// Then recurses on itself with the remaining slice
