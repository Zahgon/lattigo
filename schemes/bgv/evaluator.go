package bgv

import (
	"math/big"

	"github.com/tuneinsight/lattigo/v6/core/rlwe"
	"github.com/tuneinsight/lattigo/v6/ring"
	"github.com/tuneinsight/lattigo/v6/ring/ringqp"
)

// Evaluator is a struct that holds the necessary elements to perform the homomorphic operations between ciphertexts and/or plaintexts.
// The [Evaluator.ScaleInvariant] flag needs to be set in order to use a BFV-style
// version of the evaluator.
type Evaluator struct {
	*evaluatorBase
	*rlwe.Evaluator
	*Encoder

	// ScaleInvariant is a flag indicating whether the evaluator executes
	// scale-invariant multiplications (transforming the BGV evaluator into
	// BFV evaluator).
	ScaleInvariant bool
	pool           *rlwe.BufferPool
	poolQMul       *ring.BufferPool
}

type evaluatorBase struct {
	tMontgomery         ring.RNSScalar
	levelQMul           []int      // optimal #QiMul depending on #Qi (variable level)
	pHalf               []*big.Int // all prod(QiMul) / 2 depending on #Qi
	basisExtenderQ1toQ2 *ring.BasisExtender
}

func newEvaluatorPrecomp(parameters Parameters) *evaluatorBase {
	_ = "STUB: not implemented"
	return nil
}

// PlaintextModulus * 2^{64} mod Q

// NewEvaluator creates a new [Evaluator], that can be used to do homomorphic
// operations on ciphertexts and/or plaintexts. It stores a memory buffer
// and ciphertexts that will be used for intermediate values.
// The evaluator can optionally be initialized as scale-invariant, which
// transforms it into a BFV evaluator. See `schemes/bfv/README.md` for more
// information.
func NewEvaluator(parameters Parameters, evk rlwe.EvaluationKeySet, scaleInvariant ...bool) *Evaluator {
	_ = "STUB: not implemented"
	return nil
}

// GetParameters returns a pointer to the underlying [bgv.Parameters].
func (eval Evaluator) GetParameters() *Parameters { _ = "STUB: not implemented"; return nil }

// WithKey creates a shallow copy of this [Evaluator] in which the read-only data-structures are
// shared with the receiver but the evaluation key is set to the provided [rlwe.EvaluationKeySet].
func (eval Evaluator) WithKey(evk rlwe.EvaluationKeySet) *Evaluator {
	_ = "STUB: not implemented"
	return nil
}

// Add adds op1 to op0 and returns the result in opOut.
// inputs:
//   - op0: an *[rlwe.Ciphertext]
//   - op1:
//   - [rlwe.ElementInterface][[ring.Poly]]
//   - *big.Int, uint64, int64, int
//   - []uint64 or []int64 (of size at most N where N is the smallest integer satisfying PlaintextModulus = 1 mod 2N)
//   - opOut: an *[rlwe.Ciphertext]
//
// If op1 is an [rlwe.ElementInterface][[ring.Poly]] and the scales of op0, op1 and opOut do not match, then a scale matching operation will
// be automatically carried out to ensure that addition is performed between operands of the same scale.
// This scale matching operation will increase the noise by a small factor.
// For this reason it is preferable to ensure that all operands are already at the same scale when calling this method.
func (eval Evaluator) Add(op0 *rlwe.Ciphertext, op1 rlwe.Operand, opOut *rlwe.Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Sets op1 to the scale of op0

// If op1 > T/2 -> op1 -= T

// Scales op0 by T^{-1} mod Q

// Instantiates new plaintext from buffer

// Sets the metadata, notably matches scalses

// Encodes the vector on the plaintext

// Generic in place evaluation

func (eval Evaluator) evaluateInPlace(level int, el0 *rlwe.Ciphertext, el1 *rlwe.Element[ring.Poly], elOut *rlwe.Ciphertext, evaluate func(ring.Poly, ring.Poly, ring.Poly)) {
	_ = "STUB: not implemented"
	return
}

// If the inputs degrees differ, it copies the remaining degree on the receiver.
// checks to avoid unnecessary work.

func (eval Evaluator) matchScaleThenEvaluateInPlace(level int, el0 *rlwe.Ciphertext, el1 *rlwe.Element[ring.Poly], elOut *rlwe.Ciphertext, evaluate func(ring.Poly, uint64, ring.Poly)) {
	_ = "STUB: not implemented"
	return
}

func (eval Evaluator) newCiphertextBinary(op0, op1 rlwe.ElementInterface[ring.Poly]) (opOut *rlwe.Ciphertext) {
	_ = "STUB: not implemented"
	return nil
}

// AddNew adds op1 to op0 and returns the result on a new *[rlwe.Ciphertext] opOut.
// inputs:
//   - op0: an *[rlwe.Ciphertext]
//   - op1:
//   - [rlwe.ElementInterface][[ring.Poly]]
//   - *big.Int, uint64, int64, int
//   - []uint64 or []int64 (of size at most N where N is the smallest integer satisfying PlaintextModulus = 1 mod 2N)
//
// If op1 is an [rlwe.ElementInterface][[ring.Poly]] and the scales of op0 and op1 not match, then a scale matching operation will
// be automatically carried out to ensure that addition is performed between operands of the same scale.
// This scale matching operation will increase the noise by a small factor.
// For this reason it is preferable to ensure that all operands are already at the same scale when calling this method.
func (eval Evaluator) AddNew(op0 *rlwe.Ciphertext, op1 rlwe.Operand) (opOut *rlwe.Ciphertext, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Sub subtracts op1 to op0 and returns the result in opOut.
// inputs:
//   - op0: an *[rlwe.Ciphertext]
//   - op1:
//   - [rlwe.ElementInterface][[ring.Poly]]
//   - *big.Int, uint64, int64, int
//   - []uint64 or []int64 (of size at most N where N is the smallest integer satisfying PlaintextModulus = 1 mod 2N)
//   - opOut: an *[rlwe.Ciphertext]
//
// If op1 is an [rlwe.ElementInterface][[ring.Poly]] and the scales of op0, op1 and opOut do not match, then a scale matching operation will
// be automatically carried out to ensure that the subtraction is performed between operands of the same scale.
// This scale matching operation will increase the noise by a small factor.
// For this reason it is preferable to ensure that all operands are already at the same scale when calling this method.
func (eval Evaluator) Sub(op0 *rlwe.Ciphertext, op1 rlwe.Operand, opOut *rlwe.Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Instantiates new plaintext from buffer

// Sets the metadata, notably matches scales

// Encodes the vector on the plaintext

// Generic in place evaluation

// SubNew subtracts op1 to op0 and returns the result in a new *[rlwe.Ciphertext] opOut.
// inputs:
//   - op0: an *[rlwe.Ciphertext]
//   - op1:
//   - [rlwe.ElementInterface][[ring.Poly]]
//   - *big.Int, uint64, int64, int
//   - []uint64 or []int64 (of size at most N where N is the smallest integer satisfying PlaintextModulus = 1 mod 2N)
//
// If op1 is an [rlwe.ElementInterface][[ring.Poly]] and the scales of op0, op1 and opOut do not match, then a scale matching operation will
// be automatically carried out to ensure that the subtraction is performed between operands of the same scale.
// This scale matching operation will increase the noise by a small factor.
// For this reason it is preferable to ensure that all operands are already at the same scale when calling this method.
func (eval Evaluator) SubNew(op0 *rlwe.Ciphertext, op1 rlwe.Operand) (opOut *rlwe.Ciphertext, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DropLevel reduces the level of op0 by levels.
// No rescaling is applied during this procedure.
func (eval Evaluator) DropLevel(op0 *rlwe.Ciphertext, levels int) {
	_ = "STUB: not implemented"
	return
}

// Mul multiplies op0 with op1 without relinearization using either standard tensoring (BGV/CKKS-style) when [Evaluator.ScaleInvariant]
// is set to false or scale-invariant tensoring (BFV-style) otherwise, i.e., [Evaluator.MulScaleInvariant], and returns the result in opOut.
// This tensoring increases the noise by a multiplicative factor of the plaintext and noise norms of the operands and will usually
// require to be followed by a rescaling operation to avoid an exponential growth of the noise from subsequent multiplications.
// The procedure will return an error if either op0 or op1 are have a degree higher than 1.
// The procedure will return an error if opOut.Degree != op0.Degree + op1.Degree.
//
// inputs:
//   - op0: an *[rlwe.Ciphertext]
//   - op1:
//   - [rlwe.ElementInterface][[ring.Poly]]
//   - *big.Int, uint64, int64, int
//   - []uint64 or []int64 (of size at most N where N is the smallest integer satisfying PlaintextModulus = 1 mod 2N)
//   - opOut: an *[rlwe.Ciphertext]
//
// If op1 is an [rlwe.ElementInterface][[ring.Poly]]:
//   - the level of opOut will be updated to min(op0.Level(), op1.Level())
//   - the scale of opOut will be updated to op0.Scale * op1.Scale
func (eval Evaluator) Mul(op0 *rlwe.Ciphertext, op1 rlwe.Operand, opOut *rlwe.Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// If op1 > T/2 then subtract T to minimize the noise

// Instantiates new plaintext from buffer

// Sets the metadata, notably matches scales

// Encodes the vector on the plaintext

// MulNew multiplies op0 with op1 without relinearization using standard tensoring (BGV/CKKS-style) when [Evaluator.ScaleInvariant]
// is set to false or scale-invariant tensoring (BFV-style) otherwise, i.e., [Evaluator.MulScaleInvariantNew], and returns
// the result in a new *[rlwe.Ciphertext] opOut. This tensoring increases the noise by a multiplicative factor of the plaintext
// and noise norms of the operands and will usually require to be followed by a rescaling operation to avoid an exponential
// growth of the noise from subsequent multiplications. The procedure will return an error if either op0 or op1 are have a
// degree higher than 1.
//
// inputs:
//   - op0: an *[rlwe.Ciphertext]
//   - op1:
//   - [rlwe.ElementInterface][[ring.Poly]]
//   - *big.Int, uint64, int64, int
//   - []uint64 or []int64 (of size at most N where N is the smallest integer satisfying PlaintextModulus = 1 mod 2N)
//
// If op1 is an [rlwe.ElementInterface][[ring.Poly]]:
//   - the degree of opOut will be op0.Degree() + op1.Degree()
//   - the level of opOut will be to min(op0.Level(), op1.Level())
//   - the scale of opOut will be to op0.Scale * op1.Scale
func (eval Evaluator) MulNew(op0 *rlwe.Ciphertext, op1 rlwe.Operand) (opOut *rlwe.Ciphertext, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MulRelin multiplies op0 with op1 with relinearization using standard tensoring (BGV/CKKS-style) when [Evaluator.ScaleInvariant]
// is set to false or scale-invariant tensoring (BFV-style) otherwise, i.e., [Evaluator.MulRelinScaleInvariant], and returns the result in
// opOut. This tensoring increases the noise by a multiplicative factor of the plaintext and noise norms of the operands and will usually
// require to be followed by a rescaling operation to avoid an exponential growth of the noise from subsequent multiplications.
// The procedure will return an error if either op0.Degree or op1.Degree > 1.
// The procedure will return an error if opOut.Degree != op0.Degree + op1.Degree.
// The procedure will return an error if the evaluator was not created with an relinearization key.
//
// inputs:
//   - op0: an *[rlwe.Ciphertext]
//   - op1:
//   - [rlwe.ElementInterface][[ring.Poly]]
//   - *big.Int, uint64, int64, int
//   - []uint64 or []int64 (of size at most N where N is the smallest integer satisfying PlaintextModulus = 1 mod 2N)
//   - opOut: an *[rlwe.Ciphertext]
//
// If op1 is an [rlwe.ElementInterface][[ring.Poly]]:
//   - the level of opOut will be updated to min(op0.Level(), op1.Level())
//   - the scale of opOut will be updated to op0.Scale * op1.Scale
func (eval Evaluator) MulRelin(op0 *rlwe.Ciphertext, op1 rlwe.Operand, opOut *rlwe.Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// MulRelinNew multiplies op0 with op1 with relinearization using standard tensoring (BGV/CKKS-style) when [Evaluator.ScaleInvariant]
// is set to false or scale-invariant tensoring (BFV-style) otherwise, i.e., [Evaluator.MulRelinScaleInvariantNew], and returns the result
// in a new *[rlwe.Ciphertext] opOut. This tensoring increases the noise by a multiplicative factor of the plaintext and noise norms
// of the operands and will usually require to be followed by a rescaling operation to avoid an exponential growth of the noise from
// subsequent multiplications.
// The procedure will return an error if either op0.Degree or op1.Degree > 1.
// The procedure will return an error if the evaluator was not created with an relinearization key.
//
// inputs:
//   - op0: an *[rlwe.Ciphertext]
//   - op1:
//   - [rlwe.ElementInterface][[ring.Poly]]
//   - *big.Int, uint64, int64, int
//   - []uint64 or []int64 (of size at most N where N is the smallest integer satisfying PlaintextModulus = 1 mod 2N)
//
// If op1 is an [rlwe.ElementInterface][[ring.Poly]]:
//   - the level of opOut will be to min(op0.Level(), op1.Level())
//   - the scale of opOut will be to op0.Scale * op1.Scale
func (eval Evaluator) MulRelinNew(op0 *rlwe.Ciphertext, op1 rlwe.Operand) (opOut *rlwe.Ciphertext, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (eval Evaluator) tensorStandard(op0 *rlwe.Ciphertext, op1 *rlwe.Element[ring.Poly], relin bool, opOut *rlwe.Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Case Ciphertext (x) Ciphertext

// Avoid overwriting if the second input is the output

// Multiply by T * 2^{64} * 2^{64} -> result multipled by T and switched in the Montgomery domain

// squaring case
// c0 = c[0]*c[0]
// c2 = c[1]*c[1]
// c1 = 2*c[0]*c[1]

// regular case
// c0 = c0[0]*c0[0]
// c2 = c0[1]*c1[1]

// c1 = c0[0]*c1[1] + c0[1]*c1[0]

// Case Plaintext (x) Ciphertext or Ciphertext (x) Plaintext

// Multiply by T * 2^{64} * 2^{64} -> result multipled by T and switched in the Montgomery domain

// MulScaleInvariant multiplies op0 with op1 without relinearization and using scale invariant tensoring (BFV-style), and returns the result in opOut.
// This tensoring increases the noise by a constant factor regardless of the current noise, thus no rescaling is required with subsequent multiplications if they are
// performed with the invariant tensoring procedure. Rescaling can still be useful to reduce the size of the ciphertext, once the noise is higher than the prime
// that will be used for the rescaling or to ensure that the noise is minimal before using the regular tensoring.
// The procedure will return an error if either op0.Degree or op1.Degree > 1.
// The procedure will return an error if the evaluator was not created with an relinearization key.
//
// inputs:
//   - op0: an *[rlwe.Ciphertext]
//   - op1:
//   - [rlwe.ElementInterface][[ring.Poly]]
//   - *big.Int, uint64, int64, int
//   - []uint64 or []int64 (of size at most N where N is the smallest integer satisfying PlaintextModulus = 1 mod 2N)
//   - opOut: an *[rlwe.Ciphertext]
//
// If op1 is an [rlwe.ElementInterface][[ring.Poly]]:
//   - the level of opOut will be updated to min(op0.Level(), op1.Level())
//   - the scale of opOut will be to op0.Scale * op1.Scale * (-Q mod T)^{-1} mod T
func (eval Evaluator) MulScaleInvariant(op0 *rlwe.Ciphertext, op1 rlwe.Operand, opOut *rlwe.Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Instantiates new plaintext from buffer

// Sets the metadata, notably matches scales

// Encodes the vector on the plaintext

// MulScaleInvariantNew multiplies op0 with op1 without relinearization and using scale invariant tensoring (BFV-style), and returns the result in a new *[rlwe.Ciphertext] opOut.
// This tensoring increases the noise by a constant factor regardless of the current noise, thus no rescaling is required with subsequent multiplications if they are
// performed with the invariant tensoring procedure. Rescaling can still be useful to reduce the size of the ciphertext, once the noise is higher than the prime
// that will be used for the rescaling or to ensure that the noise is minimal before using the regular tensoring.
// The procedure will return an error if either op0.Degree or op1.Degree > 1.
// The procedure will return an error if the evaluator was not created with an relinearization key.
//
// inputs:
//   - op0: an *[rlwe.Ciphertext]
//   - op1:
//   - [rlwe.ElementInterface][[ring.Poly]]
//   - *big.Int, uint64, int64, int
//   - []uint64 or []int64 (of size at most N where N is the smallest integer satisfying PlaintextModulus = 1 mod 2N)
//
// If op1 is an [rlwe.ElementInterface][[ring.Poly]]:
//   - the level of opOut will be to min(op0.Level(), op1.Level())
//   - the scale of opOut will be to op0.Scale * op1.Scale * (-Q mod PlaintextModulus)^{-1} mod PlaintextModulus
func (eval Evaluator) MulScaleInvariantNew(op0 *rlwe.Ciphertext, op1 rlwe.Operand) (opOut *rlwe.Ciphertext, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MulRelinScaleInvariant multiplies op0 with op1 with relinearization and using scale invariant tensoring (BFV-style), and returns the result in opOut.
// This tensoring increases the noise by a constant factor regardless of the current noise, thus no rescaling is required with subsequent multiplications if they are
// performed with the invariant tensoring procedure. Rescaling can still be useful to reduce the size of the ciphertext, once the noise is higher than the prime
// that will be used for the rescaling or to ensure that the noise is minimal before using the regular tensoring.
// The procedure will return an error if either op0.Degree or op1.Degree > 1.
// The procedure will return an error if the evaluator was not created with an relinearization key.
//
// inputs:
//   - op0: an *[rlwe.Ciphertext]
//   - op1:
//   - [rlwe.ElementInterface][[ring.Poly]]
//   - *big.Int, uint64, int64, int
//   - []uint64 or []int64 (of size at most N where N is the smallest integer satisfying PlaintextModulus = 1 mod 2N)
//   - opOut: an *[rlwe.Ciphertext]
//
// If op1 is an [rlwe.ElementInterface][[ring.Poly]]:
//   - the level of opOut will be updated to min(op0.Level(), op1.Level())
//   - the scale of opOut will be to op0.Scale * op1.Scale * (-Q mod PlaintextModulus)^{-1} mod PlaintextModulus
func (eval Evaluator) MulRelinScaleInvariant(op0 *rlwe.Ciphertext, op1 rlwe.Operand, opOut *rlwe.Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Instantiates new plaintext from buffer

// Sets the metadata, notably matches scales

// Encodes the vector on the plaintext

// MulRelinScaleInvariantNew multiplies op0 with op1 with relinearization and using scale invariant tensoring (BFV-style), and returns the result in a new *[rlwe.Ciphertext] opOut.
// This tensoring increases the noise by a constant factor regardless of the current noise, thus no rescaling is required with subsequent multiplications if they are
// performed with the invariant tensoring procedure. Rescaling can still be useful to reduce the size of the ciphertext, once the noise is higher than the prime
// that will be used for the rescaling or to ensure that the noise is minimal before using the regular tensoring.
// The procedure will return an error if either op0.Degree or op1.Degree > 1.
// The procedure will return an error if the evaluator was not created with an relinearization key.
//
// inputs:
//   - op0: an *[rlwe.Ciphertext]
//   - op1:
//   - [rlwe.ElementInterface][[ring.Poly]]
//   - *big.Int, uint64, int64, int
//   - []uint64 or []int64 (of size at most N where N is the smallest integer satisfying PlaintextModulus = 1 mod 2N)
//
// If op1 is an [rlwe.ElementInterface][[ring.Poly]]:
//   - the level of opOut will be to min(op0.Level(), op1.Level())
//   - the scale of opOut will be to op0.Scale * op1.Scale * (-Q mod PlaintextModulus)^{-1} mod PlaintextModulus
func (eval Evaluator) MulRelinScaleInvariantNew(op0 *rlwe.Ciphertext, op1 rlwe.Operand) (opOut *rlwe.Ciphertext, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// tensorScaleInvariant computes (ct0 x ct1) * (t/Q) and stores the result in opOut.
func (eval Evaluator) tensorScaleInvariant(ct0 *rlwe.Ciphertext, ct1 *rlwe.Element[ring.Poly], relin bool, opOut *rlwe.Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Avoid overwriting if the second input is the output

// MulScaleInvariant returns c = a * b / (-Q[level] mod PlaintextModulus), where a, b are the input scale,
// level the level at which the operation is carried out and and c is the new scale after performing the
// invariant tensoring (BFV-style).
func MulScaleInvariant(params Parameters, a, b rlwe.Scale, level int) (c rlwe.Scale) {
	_ = "STUB: not implemented"
	return *new(rlwe.Scale)
}

func (eval Evaluator) modUpAndNTT(level, levelQMul int, ctQ0, ctQ1 *rlwe.Element[ring.Poly]) {
	_ = "STUB: not implemented"
	return
}

func (eval Evaluator) tensorLowDeg(level, levelQMul int, ct0Q0, ct1Q0, ct2Q0, ct0Q1, ct1Q1, ct2Q1 *rlwe.Element[ring.Poly]) {
	_ = "STUB: not implemented"
	return
}

// Squaring case

// c0 = c0[0]*c0[0]
// c2 = c0[1]*c0[1]
// c1 = 2*c0[0]*c0[1]

// Normal case

// c0 = c0[0]*c1[0]
// c2 = c0[1]*c1[1]
// c1 = c0[0]*c1[1] + c0[1]*c1[0]

func (eval Evaluator) quantize(level, levelQMul int, c2Q1, c2Q2 ring.Poly) {
	_ = "STUB: not implemented"
	return
}

// Applies the inverse NTT to the ciphertext, scales down the ciphertext
// by t/q and reduces its basis from QP to Q

// Extends the basis Q of ct(x) to the basis P and Divides (ct(x)Q -> P) by Q
// QP / Q -> P

// Centers ct(x)P by (P-1)/2 and extends ct(x)P to the basis Q

// (ct(x)/Q)*T, doing so only requires that Q*P > Q*Q, faster but adds error ~|T|

// MulThenAdd multiplies op0 with op1 using standard tensoring and without relinearization, and adds the result on opOut.
// The procedure will return an error if either op0.Degree() or op1.Degree() > 1.
// The procedure will return an error if either op0 == opOut or op1 == opOut.
//
// inputs:
//   - op0: an *[rlwe.Ciphertext]
//   - op1:
//   - [rlwe.ElementInterface][[ring.Poly]]
//   - *big.Int, uint64, int64, int
//   - []uint64 or []int64 (of size at most N where N is the smallest integer satisfying PlaintextModulus = 1 mod 2N)
//   - opOut: an *[rlwe.Ciphertext]
//
// If op1 is an [rlwe.ElementInterface][[ring.Poly]] and opOut.Scale != op1.Scale * op0.Scale, then a scale matching operation will
// be automatically carried out to ensure that addition is performed between operands of the same scale.
// This scale matching operation will increase the noise by a small factor.
// For this reason it is preferable to ensure that opOut.Scale == op1.Scale * op0.Scale when calling this method.
func (eval Evaluator) MulThenAdd(op0 *rlwe.Ciphertext, op1 rlwe.Operand, opOut *rlwe.Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// op1 *= (op1.Scale / opOut.Scale)

// If op1 > T/2 then subtract T to minimize the noise

// Instantiates new plaintext from buffer

// Sets the metadata, notably matches scales

// op1 *= (op1.Scale / opOut.Scale)

// Encodes the vector on the plaintext

// MulRelinThenAdd multiplies op0 with op1 using standard tensoring and with relinearization, and adds the result on opOut.
// The procedure will return an error if either op0.Degree() or op1.Degree() > 1.
// The procedure will return an error if either op0 == opOut or op1 == opOut.
//
// inputs:
//   - op0: an *[rlwe.Ciphertext]
//   - op1:
//   - [rlwe.ElementInterface][[ring.Poly]]
//   - *big.Int, uint64, int64, int
//   - []uint64 or []int64 (of size at most N where N is the smallest integer satisfying PlaintextModulus = 1 mod 2N)
//   - opOut: an *[rlwe.Ciphertext]
//
// If op1 is an [rlwe.ElementInterface][[ring.Poly]] and opOut.Scale != op1.Scale * op0.Scale, then a scale matching operation will
// be automatically carried out to ensure that addition is performed between operands of the same scale.
// This scale matching operation will increase the noise by a small factor.
// For this reason it is preferable to ensure that opOut.Scale == op1.Scale * op0.Scale when calling this method.
func (eval Evaluator) MulRelinThenAdd(op0 *rlwe.Ciphertext, op1 rlwe.Operand, opOut *rlwe.Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (eval Evaluator) mulRelinThenAdd(op0 *rlwe.Ciphertext, op1 *rlwe.Element[ring.Poly], relin bool, opOut *rlwe.Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Case Ciphertext (x) Ciphertext

// If op0.Scale * op1.Scale != opOut.Scale then
// updates op1.Scale and opOut.Scale

// Multiply by T * 2^{64} * 2^{64} -> result multipled by T and switched in the Montgomery domain

// Scales the input to the output scale

// c0 += c[0]*c[0]
// c1 += c[0]*c[1]
// c1 += c[1]*c[0]

// c2 += c[1]*c[1]

// c2 += c[1]*c[1]

// Case Plaintext (x) Ciphertext or Ciphertext (x) Plaintext

// Multiply by T * 2^{64} * 2^{64} -> result multipled by T and switched in the Montgomery domain

// If op0.Scale * op1.Scale != opOut.Scale then
// updates op1.Scale and opOut.Scale

// Rescale divides (rounded) op0 by the last prime of the moduli chain and returns the result on opOut.
// This procedure divides the noise by the last prime of the moduli chain while preserving
// the MSB-plaintext bits.
// The procedure will return an error if:
//   - op0.Level() == 0 (the input ciphertext is already at the last prime)
//   - opOut.Level() < op0.Level() - 1 (not enough space to store the result)
//
// The scale of opOut will be updated to op0.Scale * qi^{-1} mod PlaintextModulus where qi is the prime consumed by
// the rescaling operation.
// Note that if the evaluator has been instantiated as scale-invariant (BFV-style), then Rescale is a nop.
func (eval Evaluator) Rescale(op0, opOut *rlwe.Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// RelinearizeNew applies the relinearization procedure on op0 and returns the result in a new opOut.
func (eval Evaluator) RelinearizeNew(op0 *rlwe.Ciphertext) (opOut *rlwe.Ciphertext, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ApplyEvaluationKeyNew re-encrypts op0 under a different key and returns the result in a new opOut.
// It requires a [rlwe.EvaluationKey], which is computed from the key under which the Ciphertext is currently encrypted,
// and the key under which the Ciphertext will be re-encrypted.
// The procedure will return an error if either op0.Degree() or opOut.Degree() != 1.
func (eval Evaluator) ApplyEvaluationKeyNew(op0 *rlwe.Ciphertext, evk *rlwe.EvaluationKey) (opOut *rlwe.Ciphertext, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RotateColumnsNew rotates the columns of op0 by k positions to the left, and returns the result in a newly created element.
// The procedure will return an error if the corresponding Galois key has not been generated and attributed to the evaluator.
// The procedure will return an error if op0.Degree() != 1.
func (eval Evaluator) RotateColumnsNew(op0 *rlwe.Ciphertext, k int) (opOut *rlwe.Ciphertext, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RotateColumns rotates the columns of op0 by k positions to the left and returns the result in opOut.
// The procedure will return an error if the corresponding Galois key has not been generated and attributed to the evaluator.
// The procedure will return an error if either op0.Degree() or opOut.Degree() != 1.
func (eval Evaluator) RotateColumns(op0 *rlwe.Ciphertext, k int, opOut *rlwe.Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// RotateRowsNew swaps the rows of op0 and returns the result in a new opOut.
// The procedure will return an error if the corresponding Galois key has not been generated and attributed to the evaluator.
// The procedure will return an error if op0.Degree() != 1.
func (eval Evaluator) RotateRowsNew(op0 *rlwe.Ciphertext) (opOut *rlwe.Ciphertext, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RotateRows swaps the rows of op0 and returns the result in op1.
// The procedure will return an error if the corresponding Galois key has not been generated and attributed to the evaluator.
// The procedure will return an error if either op0.Degree() or op1.Degree() != 1.
func (eval Evaluator) RotateRows(op0, opOut *rlwe.Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// RotateHoistedLazyNew applies a series of rotations on the same ciphertext and returns each different rotation in a map indexed by the rotation.
// Results are not rescaled by P.
func (eval Evaluator) RotateHoistedLazyNew(level int, rotations []int, op0 *rlwe.Ciphertext, c2DecompQP []ringqp.Poly) (opOut map[int]*rlwe.Element[ringqp.Poly], err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// InnerSum divides each row of the underlying plaintext in sub-vectors of size batchSize and add n of these together.
// If n*batchSize = ctIn.Slots(), the inner sum is computed as if the plaintext was a 1-D vector of dimension ctIn.Slots()
// (we recall that a BGV/BFV plaintext is represented as a 2 x ctIn.Slots()/2 matrix).
//
// WARNING: 0 < n*batchSize <= ctIn.Slots() must divide the number of slots ctIn.Slots(). For other parameters, consider using [Evaluator.RotateAndAdd].
//
// Example for batchSize=2, n=4 and 32 slots (garbage slots are marked as X):
//
// Input:
//
// [[{a1, b1}, {c1, d1}, {e1, f1}, {g1, h1}, {i1, j1}, {k1, l1}, {m1, n1}, {o1, p1}]
//
//	[{a2, b2}, {c2, d2}, {e2, f2}, {g2, h2}, {i2, j2}, {k2, l2}, {m2, n2}, {o2, p2}]]
//
// Output:
//
// [[{a1+c1+e1+g1, b1+d1+f1+h1}, {X, X}, {X, X}, {X, X}, {i1+k1+m1+o1, j1+l1+n1+p1}, {X, X}, {X, X}, {X, X}]
//
//	[{a2+c2+e2+g2, b2+d2+f2+h2}, {X, X}, {X, X}, {X, X}, {i2+k2+m2+o2, j2+l2+n2+p2}, {X, X}, {X, X}, {X, X}]]
func (eval Evaluator) InnerSum(ctIn *rlwe.Ciphertext, batchSize, n int, opOut *rlwe.Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// RotateAndAdd computes the sum of pt_i, 0 <= i < n, where pt_i is the underlying plaintext rotated ([Evaluator.RotateRows]) by batchSize*i slots.
//
// Example: for batchSize=3, n=2, ctIn.Slots()=16:
//
// Input (recall that a BGV/BFV plaintext is represented as a 2 x ctIn.Slots()/2 matrix):
//
//	[[a, b, c, d, e, f, g, h]
//	[i, j, k, l, m, n, o, p]]
//
// Output:
//
//	[[a, b, c, d, e, f, g, h] + [[d, e, f, g, h, a, b, c] = [[a+d, b+e, c+f, d+g, e+h, f+a, g+b, h+c]
//	[i, j, k, l, m, n, o, p]]   [l, m, n, o, p, i, j, k]]   [i+l, j+m, k+n, l+o, m+p, n+i, o+j, p+k]]
//
// Calling RotateAndAdd(ctIn, 1, n, opOut) can be used to compute the inner sum of the first n slots of a plaintext.
func (eval Evaluator) RotateAndAdd(ctIn *rlwe.Ciphertext, batchSize, n int, opOut *rlwe.Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// MatchScalesAndLevel updates both input ciphertexts to ensure that their scale matches.
// To do so it computes ct0 * a = opOut * b such that:
//   - ct0.Scale * a = opOut.Scale: make the scales match.
//   - gcd(a, PlaintextModulus) == gcd(b, PlaintextModulus) == 1: ensure that the new scale is not a zero divisor if PlaintextModulus is not prime.
//   - |a+b| is minimal: minimize the added noise by the procedure.
func (eval Evaluator) MatchScalesAndLevel(ct0, opOut *rlwe.Ciphertext) {
	_ = "STUB: not implemented"
	return
}

func (eval Evaluator) GetRLWEParameters() *rlwe.Parameters { _ = "STUB: not implemented"; return nil }

func (eval Evaluator) matchScalesBinary(scale0, scale1 uint64) (r0, r1, e uint64) {
	_ = "STUB: not implemented"
	return 0, 0, 0
}

// This should never happen and if it were to happen,
// there is no way to recover from it.

func center(x, thalf, t uint64) uint64 { _ = "STUB: not implemented"; return 0 }
