// Package utils implements various helper functions.
package utils

import (
	"golang.org/x/exp/constraints"
)

// Min returns the minimum value of the two inputs.
func Min[V constraints.Ordered](a, b V) (r V) { _ = "STUB: not implemented"; return *new(V) }

// Max returns the maximum value of the two inputs.
func Max[V constraints.Ordered](a, b V) (r V) { _ = "STUB: not implemented"; return *new(V) }

// IsNil returns true either type or value are nil.
// Only interfaces or pointers to objects should be passed as argument.
func IsNil(i interface{}) bool { _ = "STUB: not implemented"; return false }

// BitReverse64 returns the bit-reverse value of the input value, within a context of 2^bitLen.
func BitReverse64[V uint64 | uint32 | int | int64](index V, bitLen int) uint64 {
	_ = "STUB: not implemented"
	return 0
}

// HammingWeight64 returns the hamming weight if the input value.
func HammingWeight64[V uint64 | uint32 | int | int64](x V) V {
	_ = "STUB: not implemented"
	return *new(V)
}

// AllDistinct returns true if all elements in s are distinct, and false otherwise.
func AllDistinct[V comparable](s []V) bool { _ = "STUB: not implemented"; return false }

// GCD computes the greatest common divisor between a and b.
func GCD[V uint64 | uint32 | int | int64](a, b V) V { _ = "STUB: not implemented"; return *new(V) }
