package minimax

import (
	"github.com/tuneinsight/lattigo/v6/circuits/ckks/bootstrapping"
	"github.com/tuneinsight/lattigo/v6/circuits/ckks/polynomial"
	"github.com/tuneinsight/lattigo/v6/core/rlwe"
	"github.com/tuneinsight/lattigo/v6/schemes/ckks"
)

// Evaluator is an evaluator used to evaluate composite polynomials on ciphertexts.
// All fields of this struct are publics, enabling custom instantiations.
type Evaluator struct {
	*ckks.Evaluator
	PolyEval   *polynomial.Evaluator
	BtsEval    bootstrapping.Bootstrapper
	Parameters ckks.Parameters
}

// NewEvaluator instantiates a new Evaluator.
// This method is allocation free.
func NewEvaluator(params ckks.Parameters, eval *ckks.Evaluator, btsEval bootstrapping.Bootstrapper) *Evaluator {
	_ = "STUB: not implemented"
	return nil
}

// Evaluate evaluates the provided MinimaxCompositePolynomial on the input ciphertext.
func (eval Evaluator) Evaluate(ct *rlwe.Ciphertext, mcp Polynomial) (res *rlwe.Ciphertext, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Checks that the number of levels available after the bootstrapping is enough to evaluate all polynomials

// Checks that res has enough level to evaluate the next polynomial, else bootstrap

// Define the scale that res must have after the polynomial evaluation.
// If we use the regular CKKS (with complex values), we chose a scale to be
// half of the desired scale, so that (x + conj(x)/2) has the correct scale.

// Evaluate the polynomial

// Clean the imaginary part (else it tends to explode)

// Reassigns the scale back to the original one

// Avoids float errors
