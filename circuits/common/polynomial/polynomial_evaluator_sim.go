package polynomial

import "github.com/tuneinsight/lattigo/v6/core/rlwe"

// SimOperand is a dummy operand that
// only stores its level and scale.
type SimOperand struct {
	Level int
	Scale rlwe.Scale
}

// SimEvaluator defines a set of method on SimOperands.
type SimEvaluator interface {
	MulNew(op0, op1 *SimOperand) *SimOperand
	Rescale(op0 *SimOperand)
	PolynomialDepth(degree int) int
	UpdateLevelAndScaleGiantStep(lead bool, tLevelOld int, tScaleOld, xPowScale rlwe.Scale) (tLevelNew int, tScaleNew rlwe.Scale)
	UpdateLevelAndScaleBabyStep(lead bool, tLevelOld int, tScaleOld rlwe.Scale) (tLevelNew int, tScaleNew rlwe.Scale)
}

// SimPowerBasis is a map storing powers of SimOperands indexed by their power.
type SimPowerBasis map[int]*SimOperand

// GenPower populates the target SimPowerBasis with the nth power.
func (d SimPowerBasis) GenPower(params rlwe.ParameterProvider, n int, eval SimEvaluator) {
	_ = "STUB: not implemented"
	return
}
