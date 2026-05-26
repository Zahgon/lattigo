// Package mod1 implements a homomorphic mod1 circuit for the CKKS scheme.
package mod1

import (
	"github.com/tuneinsight/lattigo/v6/circuits/ckks/polynomial"
	"github.com/tuneinsight/lattigo/v6/core/rlwe"
	"github.com/tuneinsight/lattigo/v6/schemes/ckks"
)

// Evaluator is an evaluator providing an API for homomorphic evaluations of scaled x mod 1.
// All fields of this struct are public, enabling custom instantiations.
type Evaluator struct {
	*ckks.Evaluator
	PolynomialEvaluator *polynomial.Evaluator
	Parameters          Parameters
}

// NewEvaluator instantiates a new [Evaluator] evaluator from [ckks.Evaluator].
// This method is allocation free.
func NewEvaluator(eval *ckks.Evaluator, evalPoly *polynomial.Evaluator, Mod1Parameters Parameters) *Evaluator {
	_ = "STUB: not implemented"
	return nil
}

// EvaluateAndScaleNew calls [EvaluateNew] and scales the output values by `scaling` (without consuming additional depth).
// If `scaling` set to 1, then this is equivalent to simply calling [EvaluateNew].
func (eval Evaluator) EvaluateAndScaleNew(ct *rlwe.Ciphertext, scaling complex128) (res *rlwe.Ciphertext, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Normalize the modular reduction to mod by 1 (division by Q)

// Compute the scales that the ciphertext should have before the double angle
// formula such that after it it has the scale it had before the polynomial
// evaluation

// Division by 1/2^r and change of variable for the Chebyshev evaluation

// Double angle

// Chebyshev evaluation

// ArcSine

// Multiplies back by q

// EvaluateNew applies an homomorphic mod Q on a vector scaled by Delta, scaled down to mod 1:
//
//  1. Delta * (Q/Delta * I(X) + m(X)) (Delta = scaling factor, I(X) integer poly, m(X) message)
//  2. Delta * (I(X) + Delta/Q * m(X)) (divide by Q/Delta)
//  3. Delta * (Delta/Q * m(X)) (x mod 1)
//  4. Delta * (m(X)) (multiply back by Q/Delta)
//
// Since Q is not a power of two, but Delta is, then does an approximate division by the closest
// power of two to Q instead. Hence, it assumes that the input plaintext is already scaled by
// the correcting factor Q/2^{round(log(Q))}.
//
// !! Assumes that the input is normalized by 1/K for K the range of the approximation.
//
// Scaling back error correction by 2^{round(log(Q))}/Q afterward is included in the polynomial
func (eval Evaluator) EvaluateNew(ct *rlwe.Ciphertext) (*rlwe.Ciphertext, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
