package rgsw

import (
	"github.com/tuneinsight/lattigo/v6/core/rlwe"
	"github.com/tuneinsight/lattigo/v6/ring"
	"github.com/tuneinsight/lattigo/v6/ring/ringqp"
)

// Evaluator is a type for evaluating homomorphic operations involving RGSW ciphertexts.
// It currently supports the external product between a RLWE and a RGSW ciphertext (see
// [Evaluator.ExternalProduct]).
type Evaluator struct {
	rlwe.Evaluator
	pool *rlwe.BufferPool
}

// NewEvaluator creates a new [Evaluator] type supporting RGSW operations in addition
// to [rlwe.Evaluator] operations.
func NewEvaluator(params rlwe.ParameterProvider, evk rlwe.EvaluationKeySet) *Evaluator {
	_ = "STUB: not implemented"
	return nil
}

// WithKey creates a shallow copy of the receiver [Evaluator] for which the evaluation key is set to the provided [rlwe.EvaluationKeySet]
// and where the temporary buffers are shared. The receiver and the returned Evaluators cannot be used concurrently.
func (eval Evaluator) WithKey(evk rlwe.EvaluationKeySet) *Evaluator {
	_ = "STUB: not implemented"
	return nil
}

// ExternalProduct computes RLWE x RGSW -> RLWE
//
//	RLWE : (-as + m + e, a)
//	x
//	RGSW : [(-as + P*w*m1 + e, a), (-bs + e, b + P*w*m1)]
//	=
//	RLWE : (<RLWE, RGSW[0]>, <RLWE, RGSW[1]>)
func (eval Evaluator) ExternalProduct(op0 *rlwe.Ciphertext, op1 *Ciphertext, opOut *rlwe.Ciphertext) {
	_ = "STUB: not implemented"
	return
}

// If log(Q) * (Q-1)**2 < 2^{64}-1

func (eval Evaluator) externalProduct32Bit(ct0 *rlwe.Ciphertext, rgsw *Ciphertext, c0, c1 ring.Poly) {
	_ = "STUB: not implemented"

	// rgsw = [(-as + P*w*m1 + e, a), (-bs + e, b + P*w*m1)]
	// ct = [-cs + m0 + e, c]
	// opOut = [<ct, rgsw[0]>, <ct, rgsw[1]>] = [ct[0] * rgsw[0][0] + ct[1] * rgsw[0][1], ct[0] * rgsw[1][0] + ct[1] * rgsw[1][1]]
	return
}

// (a, b) + (c0 * rgsw[0][0], c0 * rgsw[0][1])
// (a, b) + (c1 * rgsw[1][0], c1 * rgsw[1][1])

// TODO: center values if mask = 0

func (eval Evaluator) externalProductInPlaceSinglePAndBitDecomp(ct0 *rlwe.Ciphertext, rgsw *Ciphertext, c0QP, c1QP ringqp.Poly) {
	_ = "STUB: not implemented"

	// rgsw = [(-as + P*w*m1 + e, a), (-bs + e, b + P*w*m1)]
	// ct = [-cs + m0 + e, c]
	// opOut = [<ct, rgsw[0]>, <ct, rgsw[1]>] = [ct[0] * rgsw[0][0] + ct[1] * rgsw[0][1], ct[0] * rgsw[1][0] + ct[1] * rgsw[1][1]]
	return
}

// (a, b) + (c0 * rgsw[k][0], c0 * rgsw[k][1])

// TODO: center values if mask == 0

func (eval Evaluator) externalProductInPlaceMultipleP(levelQ, levelP int, ct0 *rlwe.Ciphertext, rgsw *Ciphertext, c0OutQ, c0OutP, c1OutQ, c1OutP ring.Poly) {
	_ = "STUB: not implemented"
	return
}

// (a, b) + (c0 * rgsw[0][0], c0 * rgsw[0][1])

// AddLazy adds op to opOut, without modular reduction.
func AddLazy(op interface{}, ringQP ringqp.Ring, opOut *Ciphertext) {
	_ = "STUB: not implemented"
	return
}

// Doesn't matter which one since we add without modular reduction

// Reduce applies the modular reduction on ctIn and returns the result on opOut.
func Reduce(ctIn *Ciphertext, ringQP ringqp.Ring, opOut *Ciphertext) {
	_ = "STUB: not implemented"
	return
}

// MulByXPowAlphaMinusOneLazy multiplies opOut by (X^alpha - 1) and returns the result on opOut.
func MulByXPowAlphaMinusOneLazy(ctIn *Ciphertext, powXMinusOne ringqp.Poly, ringQP ringqp.Ring, opOut *Ciphertext) {
	_ = "STUB: not implemented"
	return
}

// MulByXPowAlphaMinusOneThenAddLazy multiplies opOut by (X^alpha - 1) and adds the result on opOut.
func MulByXPowAlphaMinusOneThenAddLazy(ctIn *Ciphertext, powXMinusOne ringqp.Poly, ringQP ringqp.Ring, opOut *Ciphertext) {
	_ = "STUB: not implemented"
	return
}
