package bignum

import (
	"io"
	"math/big"
)

// NewInt allocates a new *big.Int.
// Accepted types are: string, uint, uint64, int64, int, *big.Float or *big.Int.
func NewInt(x interface{}) (y *big.Int) { _ = "STUB: not implemented"; return nil }

// RandInt generates a random Int in [0, max-1].
func RandInt(reader io.Reader, max *big.Int) (n *big.Int) { _ = "STUB: not implemented"; return nil }

// DivRound sets the target i to round(a/b).
func DivRound(a, b, i *big.Int) { _ = "STUB: not implemented"; return }
