package lintrans

import (
	"github.com/tuneinsight/lattigo/v6/core/rlwe"
	"github.com/tuneinsight/lattigo/v6/ring/ringqp"
	"github.com/tuneinsight/lattigo/v6/schemes"
)

type Evaluator struct {
	schemes.Evaluator
	pool *rlwe.BufferPool
}

func NewEvaluator(eval schemes.Evaluator) *Evaluator { _ = "STUB: not implemented"; return nil }

// EvaluateMany takes as input a ciphertext ctIn, a list of linear transformations [M0, M1, M2, ...] and a list of pre-allocated receiver opOut
// and evaluates opOut: [M0(ctIn), M1(ctIn), M2(ctIn), ...]
func (eval Evaluator) EvaluateMany(ctIn *rlwe.Ciphertext, linearTransformations []LinearTransformation, opOut []*rlwe.Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// PreRotatedCiphertextForDiagonalMatrixMultiplication populates ctPreRot with the pre-rotated ciphertext for the rotations rots and deletes rotated ciphertexts that are not in rots.
func (eval Evaluator) PreRotatedCiphertextForDiagonalMatrixMultiplication(levelQ, levelP int, ctIn *rlwe.Ciphertext, BuffDecompQP []ringqp.Poly, rots []int, ctPreRot map[int]*rlwe.Element[ringqp.Poly]) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Cleans the map of non-required ciphertexts

// Computes the rotation only for the ones that are not already present.

// EvaluateSequential takes as input a ciphertext ctIn and a list of linear transformations [M0, M1, M2, ...] and evaluates opOut:...M2(M1(M0(ctIn))
func (eval Evaluator) EvaluateSequential(ctIn *rlwe.Ciphertext, linearTransformations []LinearTransformation, opOut *rlwe.Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// MultiplyByDiagMatrix multiplies the Ciphertext ctIn by the plaintext matrix and returns the result on the Ciphertext
// opOut. Memory buffers for the decomposed ciphertext BuffDecompQP, BuffDecompQP must be provided, those are list of poly of ringQ and ringP
// respectively, each of size params.Beta().
// The naive approach is used (single hoisting and no baby-step giant-step), which is faster than MultiplyByDiagMatrixBSGS
// for matrix of only a few non-zero diagonals but uses more keys.
func (eval Evaluator) MultiplyByDiagMatrix(ctIn *rlwe.Ciphertext, matrix LinearTransformation, BuffDecompQP []ringqp.Poly, opOut *rlwe.Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// ct0 * P mod Q

// P*c0

// keyswitch(c1_Q) = (d0_QP, d1_QP)

// keyswitch(c1_Q) = (d0_QP, d1_QP)

// sum(phi(c0 * P + d0_QP))/P
// sum(phi(d1_QP))/P

// Rotation by zero
// opOut += c0_Q * plaintext
// opOut += c1_Q * plaintext

// MultiplyByDiagMatrixBSGS multiplies the Ciphertext "ctIn" by the plaintext matrix "matrix" and returns the result on the Ciphertext "opOut".
// ctInPreRotated can be obtained with [PreRotatedCiphertextForDiagonalMatrixMultiplication].
// The BSGS approach is used (double hoisting with baby-step giant-step), which is faster than MultiplyByDiagMatrix
// for matrix with more than a few non-zero diagonals and uses significantly less keys.
func (eval Evaluator) MultiplyByDiagMatrixBSGS(ctIn *rlwe.Ciphertext, matrix LinearTransformation, ctInPreRot map[int]*rlwe.Element[ringqp.Poly], opOut *rlwe.Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Computes the N2 rotations indexes of the non-zero rows of the diagonalized DFT matrix for the baby-step giant-step algorithm

// Accumulator inner loop

// Accumulator outer loop

// Result in QP

// P*c0
// P*c1

// OUTER LOOP

// INNER LOOP

// If j != 0, then rotates ((tmp0QP.Q, tmp0QP.P), (tmp1QP.Q, tmp1QP.P)) by N1*j and adds the result on ((cQP.Value[0].Q, cQP.Value[0].P), (cQP.Value[1].Q, cQP.Value[1].P))

// Hoisting of the ModDown of sum(sum(phi(d1) * plaintext))
// c1 * plaintext + sum(phi(d1) * plaintext) + phi(c1) * plaintext mod Q

// EvaluationKey(P*phi(tmpRes_1)) = (d0, d1) in base QP

// Outer loop rotations

// Else directly adds on ((cQP.Value[0].Q, cQP.Value[0].P), (cQP.Value[1].Q, cQP.Value[1].P))

// sum(phi(c0 * P + d0_QP))/P
// sum(phi(d1_QP))/P
