// Package polynomial implements a homomorphic polynomial evaluator for the BFV/BGV schemes.
package polynomial

import (
	"github.com/tuneinsight/lattigo/v6/circuits/common/polynomial"
	"github.com/tuneinsight/lattigo/v6/schemes/bgv"
)

// Polynomial is a type wrapping the type [polynomial.Polynomial].
type Polynomial polynomial.Polynomial

// NewPolynomial creates a new Polynomial from a list of coefficients []T.
func NewPolynomial[T bgv.Integer](coeffs []T) Polynomial {
	_ = "STUB: not implemented"
	return *new(Polynomial)
}

// PolynomialVector is a type wrapping the type [polynomial.PolynomialVector].
type PolynomialVector polynomial.PolynomialVector

// Depth returns the depth of the target [PolynomialVector].
func (p PolynomialVector) Depth() int { _ = "STUB: not implemented"; return 0 }

func NewPolynomialVector[T bgv.Integer](polys [][]T, mapping map[int][]int) (PolynomialVector, error) {
	_ = "STUB: not implemented"
	return *new(PolynomialVector), nil
}
