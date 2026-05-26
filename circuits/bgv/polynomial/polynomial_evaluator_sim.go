package polynomial

import (
	"github.com/tuneinsight/lattigo/v6/circuits/common/polynomial"
	"github.com/tuneinsight/lattigo/v6/core/rlwe"
	"github.com/tuneinsight/lattigo/v6/schemes/bgv"
)

// simEvaluator is a struct used to pre-computed the scaling
// factors of the polynomial coefficients used by the inlined
// polynomial evaluation by running the polynomial evaluation
// with dummy operands.
// This struct implements the interface [polynomial.SimEvaluator].
type simEvaluator struct {
	params             bgv.Parameters
	InvariantTensoring bool
}

// PolynomialDepth returns the depth of the polynomial.
func (d simEvaluator) PolynomialDepth(degree int) int { _ = "STUB: not implemented"; return 0 }

// Rescale rescales the target polynomial.SimOperand n times and returns it.
func (d simEvaluator) Rescale(op0 *polynomial.SimOperand) { _ = "STUB: not implemented"; return }

// MulNew multiplies two polynomial.SimOperand, stores the result the target [polynomial.SimOperand] and returns the result.
func (d simEvaluator) MulNew(op0, op1 *polynomial.SimOperand) (opOut *polynomial.SimOperand) {
	_ = "STUB: not implemented"
	return nil
}

// UpdateLevelAndScaleBabyStep returns the updated level and scale for a baby-step.
func (d simEvaluator) UpdateLevelAndScaleBabyStep(lead bool, tLevelOld int, tScaleOld rlwe.Scale) (tLevelNew int, tScaleNew rlwe.Scale) {
	_ = "STUB: not implemented"
	return 0, *new(rlwe.Scale)
}

// UpdateLevelAndScaleGiantStep returns the updated level and scale for a giant-step.
func (d simEvaluator) UpdateLevelAndScaleGiantStep(lead bool, tLevelOld int, tScaleOld, xPowScale rlwe.Scale) (tLevelNew int, tScaleNew rlwe.Scale) {
	_ = "STUB: not implemented"
	return 0, *new(rlwe.Scale)
}

// tScaleNew = targetScale*currentQi/XPow.Scale

// -Q mod T
