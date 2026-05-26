// Package polynomial implements a homomorphic polynomial evaluator for the CKKS scheme.
package polynomial

import (
	"math/big"

	"github.com/tuneinsight/lattigo/v6/circuits/common/polynomial"
	"github.com/tuneinsight/lattigo/v6/utils/bignum"
)

// Polynomial is a type wrapping the type [polynomial.Polynomial].
type Polynomial polynomial.Polynomial

// NewPolynomial creates a new Polynomial from a [bignum.Polynomial].
func NewPolynomial(poly bignum.Polynomial) Polynomial {
	_ = "STUB: not implemented"
	return *new(Polynomial)
}

// PolynomialVector is a type wrapping the type [polynomial.PolynomialVector].
type PolynomialVector polynomial.PolynomialVector

// Depth returns the depth of the target [PolynomialVector].
func (p PolynomialVector) Depth() int { _ = "STUB: not implemented"; return 0 }

// NewPolynomialVector creates a new PolynomialVector from a list of [bignum.Polynomial] and a mapping
// map[poly_index][slots_index] which stores which polynomial has to be evaluated on which slot.
// Slots that are not referenced in this mapping will be evaluated to zero.
// User must ensure that a same slot is not referenced twice.
func NewPolynomialVector(polys []bignum.Polynomial, mapping map[int][]int) (PolynomialVector, error) {
	_ = "STUB: not implemented"
	return *new(PolynomialVector), nil
}

func (p PolynomialVector) ChangeOfBasis(slots int) (scalar, constant []*big.Float) {
	_ = "STUB: not implemented"
	return nil, nil
}
