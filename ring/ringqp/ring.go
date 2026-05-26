// Package ringqp is implements a wrapper for both the ringQ and ringP.
package ringqp

import (
	"math/big"

	"github.com/tuneinsight/lattigo/v6/ring"
)

// Ring is a structure that implements the operation in the ring R_QP.
// This type is simply a union type between the two Ring types representing
// R_Q and R_P.
type Ring struct {
	RingQ, RingP *ring.Ring
}

func (r Ring) N() int { _ = "STUB: not implemented"; return 0 }

// AtLevel returns a shallow copy of the target ring configured to
// carry on operations at the specified levels.
func (r Ring) AtLevel(levelQ, levelP int) Ring { _ = "STUB: not implemented"; return *new(Ring) }

// PolyToBigintCentered reconstructs p1 and returns the result in an array of Int.
// Coefficients are centered around Q/2
// gap defines coefficients X^{i*gap} that will be reconstructed.
// For example, if gap = 1, then all coefficients are reconstructed, while
// if gap = 2 then only coefficients X^{2*i} are reconstructed.
func (r Ring) PolyToBigintCentered(p1 Poly, gap int, coeffsBigint []*big.Int) {
	_ = "STUB: not implemented"
	return
}

// Q

// P

// Centers the coefficients

// Log2OfStandardDeviation returns base 2 logarithm of the standard deviation of the coefficients
// of the polynomial.
func (r Ring) Log2OfStandardDeviation(poly Poly) (std float64) { _ = "STUB: not implemented"; return 0 }

// LevelQ returns the level at which the target
// ring operates for the modulus Q.
func (r Ring) LevelQ() int { _ = "STUB: not implemented"; return 0 }

// LevelP returns the level at which the target
// ring operates for the modulus P.
func (r Ring) LevelP() int { _ = "STUB: not implemented"; return 0 }

func (r Ring) Equal(p1, p2 Poly) (v bool) { _ = "STUB: not implemented"; return false }

// NewPoly creates a new polynomial with all coefficients set to 0.
func (r Ring) NewPoly() Poly { _ = "STUB: not implemented"; return *new(Poly) }
