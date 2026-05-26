// Package sampling implements secure sampling of bytes and integers.
package sampling

import (
	"math/big"
)

// RandUint64 return a random value between 0 and 0xFFFFFFFFFFFFFFFF.
func RandUint64() uint64 { _ = "STUB: not implemented"; return 0 }

// RandFloat64 returns a random float between min and max.
func RandFloat64(min, max float64) float64 { _ = "STUB: not implemented"; return 0 }

// RandComplex128 returns a random complex with the real and imaginary part between min and max.
func RandComplex128(min, max float64) complex128 { _ = "STUB: not implemented"; return 0 }

// RandInt generates a random Int in [0, max-1].
func RandInt(max *big.Int) (n *big.Int) { _ = "STUB: not implemented"; return nil }
