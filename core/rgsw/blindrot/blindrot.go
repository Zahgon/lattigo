// Package blindrot implements blind rotations evaluation for RLWE schemes.
package blindrot

import (
	"github.com/tuneinsight/lattigo/v6/core/rlwe"
	"github.com/tuneinsight/lattigo/v6/ring"
)

// InitTestPolynomial takes a function g, and creates a test polynomial polynomial for the function in the interval [a, b].
// Inputs to the blind rotation evaluation are assumed to have been normalized with the change of basis (2*x - a - b)/(b-a).
// Interval [a, b] should take into account the "drift" of the value x, caused by the change of modulus from Q to 2N.
func InitTestPolynomial(g func(x float64) (y float64), scale rlwe.Scale, ringQ *ring.Ring, a, b float64) (F ring.Poly) {
	_ = "STUB: not implemented"
	return *new(ring.Poly)
}

// Discretization interval

// Interval [-1, 0] of g(x)

// Interval ]0, 1[ of g(x)
