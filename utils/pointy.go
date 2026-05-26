package utils

import (
	cs "golang.org/x/exp/constraints"
)

type Number interface {
	cs.Complex | cs.Float | cs.Integer
}

// Pointy creates a new T variable and returns its pointer.
func Pointy[T Number](x T) *T {
	_ = "STUB: not implemented"

	// PointyIntToPointUint64 converts *int to *uint64.
	return nil
}

func PointyIntToPointUint64(x *int) *uint64 {
	_ = "STUB: not implemented"
	/* #nosec G103 -- behavior and consequences well understood, pointer type cast */ return nil
}
