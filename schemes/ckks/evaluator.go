package ckks

import (
	"github.com/tuneinsight/lattigo/v6/core/rlwe"
	"github.com/tuneinsight/lattigo/v6/ring"
	"github.com/tuneinsight/lattigo/v6/ring/ringqp"
)

// Evaluator is a struct that holds the necessary elements to execute the homomorphic operations between Ciphertexts and/or Plaintexts.
// It also holds a memory buffer used to store intermediate computations.
type Evaluator struct {
	*Encoder
	*rlwe.Evaluator
	pool *rlwe.BufferPool
}

// NewEvaluator creates a new [Evaluator], that can be used to do homomorphic
// operations on the Ciphertexts and/or Plaintexts. It stores a memory buffer
// and Ciphertexts that will be used for intermediate values.
func NewEvaluator(parameters Parameters, evk rlwe.EvaluationKeySet) *Evaluator {
	_ = "STUB: not implemented"
	return nil
}

// GetParameters returns a pointer to the underlying [ckks.Parameters].
func (eval Evaluator) GetParameters() *Parameters { _ = "STUB: not implemented"; return nil }

// GetRLWEParameters returns a pointer to the underlying [rlwe.Parameters].
func (eval Evaluator) GetRLWEParameters() *rlwe.Parameters { _ = "STUB: not implemented"; return nil }

// Add adds op1 to op0 and returns the result in opOut.
// The following types are accepted for op1:
//   - [rlwe.ElementInterface][[ring.Poly]]
//   - complex128, float64, int, int64, uint, uint64, *[big.Int], *[big.Float], *[bignum.Complex]
//   - []complex128, []float64, []*[big.Float] or []*[bignum.Complex] of size at most params.MaxSlots()
//
// Passing an invalid type will return an error.
func (eval Evaluator) Add(op0 *rlwe.Ciphertext, op1 rlwe.Operand, opOut *rlwe.Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Checks operand validity and retrieves minimum level

// Generic inplace evaluation

// Convertes the scalar to a complex RNS scalar

// Generic inplace evaluation

// Resize step ensures identical size

// Instantiates new plaintext from buffer

// Sets the metadata, notably matches scales

// Encodes the vector on the plaintext

// Generic in place evaluation

// AddNew adds op1 to op0 and returns the result in a newly created element opOut.
// The following types are accepted for op1:
//   - [rlwe.ElementInterface][[ring.Poly]]
//   - complex128, float64, int, int64, uint, uint64, *[big.Int], *[big.Float], *[bignum.Complex]
//   - []complex128, []float64, []*[big.Float] or []*[bignum.Complex] of size at most params.MaxSlots()
//
// Passing an invalid type will return an error.
func (eval Evaluator) AddNew(op0 *rlwe.Ciphertext, op1 rlwe.Operand) (opOut *rlwe.Ciphertext, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Sub subtracts op1 from op0 and returns the result in opOut.
// The following types are accepted for op1:
//   - [rlwe.ElementInterface][[ring.Poly]]
//   - complex128, float64, int, int64, uint, uint64, *[big.Int], *[big.Float], *[bignum.Complex]
//   - []complex128, []float64, []*[big.Float] or []*[bignum.Complex] of size at most params.MaxSlots()
//
// Passing an invalid type will return an error.
func (eval Evaluator) Sub(op0 *rlwe.Ciphertext, op1 rlwe.Operand, opOut *rlwe.Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Checks operand validity and retrieves minimum level

// Generic inplace evaluation

// Negates high degree ciphertext coefficients if the degree of the second operand is larger than the first operand

// Convertes the scalar to a complex RNS scalar

// Generic inplace evaluation

// Resize step ensures identical size

// Instantiates new plaintext from buffer

// Encodes the vector on the plaintext

// Generic inplace evaluation

// SubNew subtracts op1 from op0 and returns the result in a newly created element opOut.
// The following types are accepted for op1:
//   - [rlwe.ElementInterface][[ring.Poly]]
//   - complex128, float64, int, int64, uint, uint64, *[big.Int], *[big.Float], *[bignum.Complex]
//   - []complex128, []float64, []*[big.Float] or []*[bignum.Complex] of size at most params.MaxSlots()
//
// Passing an invalid type will return an error.
func (eval Evaluator) SubNew(op0 *rlwe.Ciphertext, op1 rlwe.Operand) (opOut *rlwe.Ciphertext, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (eval Evaluator) evaluateInPlace(level int, c0 *rlwe.Ciphertext, c1 *rlwe.Element[ring.Poly], opOut *rlwe.Ciphertext, evaluate func(ring.Poly, ring.Poly, ring.Poly)) {
	_ = "STUB: not implemented"
	return
}

// Checks whether or not the receiver element is the same as one of the input elements
// and acts accordingly to avoid unnecessary element creation or element overwriting,
// and scales properly the element before the evaluation.

// Sanity check, this error should not happen unless the evaluator's buffers
// were improperly tempered with.

// Will avoid resizing on the output

// Sanity check, this error should not happen unless the evaluator's buffers
// were improperly tempered with.

// Will avoid resizing on the output

// Sanity check, this error should not happen unless the evaluator's buffers
// were improperly tempered with.

// Sanity check, this error should not happen unless the evaluator's buffers
// were improperly tempered with.

// If the inputs degrees differ, it copies the remaining degree on the receiver.
// Also checks that the receiver is not one of the inputs to avoid unnecessary work.

func (eval Evaluator) evaluateWithScalar(level int, p0 []ring.Poly, RNSReal, RNSImag ring.RNSScalar, p1 []ring.Poly, evaluate func(ring.Poly, ring.RNSScalar, ring.RNSScalar, ring.Poly)) {
	_ = "STUB: not implemented"

	// Component wise operation with the following vector:
	// [a + b*psi_qi^2, ....., a + b*psi_qi^2, a - b*psi_qi^2, ...., a - b*psi_qi^2] mod Qi
	// [{                  N/2                }{                N/2               }]
	// Which is equivalent outside of the NTT domain to evaluating a to the first coefficient of op0 and b to the N/2-th coefficient of op0.
	return
}

// ScaleUpNew multiplies op0 by scale and sets its scale to its previous scale times scale returns the result in opOut.
func (eval Evaluator) ScaleUpNew(op0 *rlwe.Ciphertext, scale rlwe.Scale) (opOut *rlwe.Ciphertext, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ScaleUp multiplies op0 by scale and sets its scale to its previous scale times scale returns the result in opOut.
func (eval Evaluator) ScaleUp(op0 *rlwe.Ciphertext, scale rlwe.Scale, opOut *rlwe.Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// SetScale sets the scale of the ciphertext to the input scale (consumes a level).
func (eval Evaluator) SetScale(ct *rlwe.Ciphertext, scale rlwe.Scale) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// DropLevelNew reduces the level of op0 by levels and returns the result in a newly created element.
// No rescaling is applied during this procedure.
func (eval Evaluator) DropLevelNew(op0 *rlwe.Ciphertext, levels int) (opOut *rlwe.Ciphertext) {
	_ = "STUB: not implemented"
	return nil
}

// DropLevel reduces the level of op0 by levels and returns the result in op0.
// No rescaling is applied during this procedure.
func (eval Evaluator) DropLevel(op0 *rlwe.Ciphertext, levels int) {
	_ = "STUB: not implemented"
	return
}

// Rescale divides op0 by the last prime of the moduli chain and repeats this procedure
// params.LevelsConsumedPerRescaling() times.
//
// Returns an error if:
//   - Either op0 or opOut MetaData are nil
//   - The level of op0 is too low to enable a rescale
func (eval Evaluator) Rescale(op0, opOut *rlwe.Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// RescaleTo divides op0 by the last prime in the moduli chain, and repeats this procedure (consuming one level each time)
// and stops if the scale reaches `minScale` or if it would go below `minscale/2`, and returns the result in opOut.
// Returns an error if:
// - minScale <= 0
// - ct.Scale <= 0
// - ct.Level() = 0
func (eval Evaluator) RescaleTo(op0 *rlwe.Ciphertext, minScale rlwe.Scale, opOut *rlwe.Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Divides the scale by each moduli of the modulus chain as long as the scale isn't smaller than minScale/2
// or until the output Level() would be zero

// MulNew multiplies op0 with op1 without relinearization and returns the result in a newly created element opOut.
//
// op1.(type) can be
//   - [rlwe.ElementInterface][[ring.Poly]]
//   - complex128, float64, int, int64, uint, uint64, *[big.Int], *[big.Float], *[bignum.Complex]
//   - []complex128, []float64, []*[big.Float] or []*[bignum.Complex] of size at most params.MaxSlots()
//
// If op1.(type) == rlwe.ElementInterface[ring.Poly]:
//   - The procedure will return an error if either op0.Degree or op1.Degree > 1.
func (eval Evaluator) MulNew(op0 *rlwe.Ciphertext, op1 rlwe.Operand) (opOut *rlwe.Ciphertext, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Mul multiplies op0 with op1 without relinearization and returns the result in opOut.
//
// The following types are accepted for op1:
//   - [rlwe.ElementInterface][[ring.Poly]]
//   - complex128, float64, int, int64, uint, uint64, *[big.Int], *[big.Float], *[bignum.Complex]
//   - []complex128, []float64, []*[big.Float] or []*[bignum.Complex] of size at most params.MaxSlots()
//
// Passing an invalid type will return an error.
//
// If op1.(type) == [rlwe.ElementInterface][[ring.Poly]]:
//   - The procedure will return an error if either op0 or op1 are have a degree higher than 1.
//   - The procedure will return an error if opOut.Degree != op0.Degree + op1.Degree.
func (eval Evaluator) Mul(op0 *rlwe.Ciphertext, op1 rlwe.Operand, opOut *rlwe.Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Generic in place evaluation

// Convertes the scalar to a *bignum.Complex

// Gets the ring at the target level

// Scalar is a GaussianInteger, thus no scaling required

// Current modulus scaling factor

// If DefaultScalingFactor > 2^60, then multiple moduli are used per single rescale
// thus continues multiplying the scale with the appropriate number of moduli

// Convertes the *bignum.Complex to a complex RNS scalar

// Generic in place evaluation

// Copies the metadata on the output
// updates the scaling factor

// Gets the ring at the target level

// Instantiates new plaintext from buffer

// If DefaultScalingFactor > 2^60, then multiple moduli are used per single rescale
// thus continues multiplying the scale with the appropriate number of moduli

// Encodes the vector on the plaintext

// Generic in place evaluation

// MulRelinNew multiplies op0 with op1 with relinearization and returns the result in a newly created element.
//
// The following types are accepted for op1:
//   - [rlwe.ElementInterface][[ring.Poly]]
//   - complex128, float64, int, int64, uint, uint64, *[big.Int], *[big.Float], *[bignum.Complex]
//   - []complex128, []float64, []*[big.Float] or []*[bignum.Complex] of size at most params.MaxSlots()
//
// Passing an invalid type will return an error.
//
// The procedure will return an error if either op0.Degree or op1.Degree > 1.
// The procedure will return an error if the evaluator was not created with an relinearization key.
func (eval Evaluator) MulRelinNew(op0 *rlwe.Ciphertext, op1 rlwe.Operand) (opOut *rlwe.Ciphertext, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MulRelin multiplies op0 with op1 with relinearization and returns the result in opOut.
//
// The following types are accepted for op1:
//   - [rlwe.ElementInterface][[ring.Poly]]
//   - complex128, float64, int, int64, uint, uint64, *[big.Int], *[big.Float], *[bignum.Complex]
//   - []complex128, []float64, []*[big.Float] or []*[bignum.Complex] of size at most params.MaxSlots()
//
// Passing an invalid type will return an error.
//
// The procedure will return an error if either op0.Degree or op1.Degree > 1.
// The procedure will return an error if opOut.Degree != op0.Degree + op1.Degree.
// The procedure will return an error if the evaluator was not created with an relinearization key.
func (eval Evaluator) MulRelin(op0 *rlwe.Ciphertext, op1 rlwe.Operand, opOut *rlwe.Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (eval Evaluator) mulRelin(op0 *rlwe.Ciphertext, op1 *rlwe.Element[ring.Poly], relin bool, opOut *rlwe.Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Case Ciphertext (x) Ciphertext

// Avoid overwriting if the second input is the output

// squaring case
// c0 = c[0]*c[0]
// c2 = c[1]*c[1]
// c1 = 2*c[0]*c[1]

// regular case
// c0 = c0[0]*c0[0]
// c2 = c0[1]*c1[1]

// c1 = c0[0]*c1[1] + c0[1]*c1[0]

// Case Plaintext (x) Ciphertext or Ciphertext (x) Plaintext

// MulThenAdd evaluate opOut = opOut + op0 * op1.
//
// The following types are accepted for op1:
//   - rlwe.ElementInterface[ring.Poly]
//   - complex128, float64, int, int64, uint, uint64, *[big.Int], *[big.Float], *[bignum.Complex]
//   - []complex128, []float64, []*[big.Float] or []*[bignum.Complex] of size at most params.MaxSlots()
//
// Passing an invalid type will return an error.
//
// If op1.(type) is complex128, float64, int, int64, uint64. *[big.Float], *[big.Int] or *[ring.Complex]:
//
// This function will not modify op0 but will multiply opOut by Q[min(op0.Level(), opOut.Level())] if:
//   - op0.Scale == opOut.Scale
//   - constant is not a Gaussian integer.
//
// If op0.Scale == opOut.Scale, and constant is not a Gaussian integer, then the constant will be scaled by
// Q[min(op0.Level(), opOut.Level())] else if opOut.Scale > op0.Scale, the constant will be scaled by opOut.Scale/op0.Scale.
//
// To correctly use this function, make sure that either op0.Scale == opOut.Scale or
// opOut.Scale = op0.Scale * Q[min(op0.Level(), opOut.Level())].
//
// If op1.(type) is []complex128, []float64, []*[big.Float] or []*[bignum.Complex]:
//   - If opOut.Scale == op0.Scale, op1 will be encoded and scaled by Q[min(op0.Level(), opOut.Level())]
//   - If opOut.Scale > op0.Scale, op1 will be encoded ans scaled by opOut.Scale/op1.Scale.
//
// Then the method will recurse with op1 given as [rlwe.ElementInterface][[ring.Poly]].
//
// If op1.(type) is [rlwe.ElementInterface][[ring.Poly]], the multiplication is carried outwithout relinearization and:
//
// This function will return an error if op0.Scale > opOut.Scale and user must ensure that opOut.Scale <= op0.Scale * op1.Scale.
// If opOut.Scale < op0.Scale * op1.Scale, then scales up opOut before adding the result.
// Additionally, the procedure will return an error if:
//   - either op0 or op1 are have a degree higher than 1.
//   - opOut.Degree != op0.Degree + op1.Degree.
//   - opOut = op0 or op1.
func (eval Evaluator) MulThenAdd(op0 *rlwe.Ciphertext, op1 rlwe.Operand, opOut *rlwe.Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Gets the ring at the minimum level

// Convertes the scalar to a *bignum.Complex

// If op0 and opOut scales are identical, but the op1 is not a Gaussian integer then multiplies opOut by scaleRLWE.
// This ensures noiseless addition with opOut = scaleRLWE * opOut + op0 * round(scalar * scaleRLWE).

// opOut.Scale > op0.Scale then the scaling factor for op1 becomes the quotient between the two scales

// Gets the ring at the target level

// If op0 and opOut scales are identical then multiplies opOut by scaleRLWE.

// opOut.Scale > op0.Scale then the scaling factor for op1 becomes the quotient between the two scales

// Instantiates new plaintext from buffer

// Encodes the vector on the plaintext

// MulRelinThenAdd multiplies op0 with op1 with relinearization and adds the result on opOut.
//
// The following types are accepted for op1:
//   - [rlwe.ElementInterface][[ring.Poly]]
//   - complex128, float64, int, int64, uint, uint64, *[big.Int], *[big.Float], *[bignum.Complex]
//   - []complex128, []float64, []*[big.Float] or []*[bignum.Complex] of size at most params.MaxSlots()
//
// Passing an invalid type will return an error.
//
// User must ensure that opOut.Scale <= op0.Scale * op1.Scale.
//
// If opOut.Scale < op0.Scale * op1.Scale, then scales up opOut before adding the result.
//
// The procedure will return an error if either op0.Degree or op1.Degree > 1.
// The procedure will return an error if opOut.Degree != op0.Degree + op1.Degree.
// The procedure will return an error if the evaluator was not created with an relinearization key.
// The procedure will return an error if opOut = op0 or op1.
func (eval Evaluator) MulRelinThenAdd(op0 *rlwe.Ciphertext, op1 rlwe.Operand, opOut *rlwe.Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (eval Evaluator) mulRelinThenAdd(op0 *rlwe.Ciphertext, op1 *rlwe.Element[ring.Poly], relin bool, opOut *rlwe.Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Only scales up if int(ratio) >= 2

// Case Ciphertext (x) Ciphertext

// c0 += c[0]*c[0]
// c1 += c[0]*c[1]
// c1 += c[1]*c[0]

// c2 += c[1]*c[1]

// c2 += c[1]*c[1]

// Case Plaintext (x) Ciphertext or Ciphertext (x) Plaintext

// RelinearizeNew applies the relinearization procedure on op0 and returns the result in a newly
// created Ciphertext. The input Ciphertext must be of degree two.
func (eval Evaluator) RelinearizeNew(op0 *rlwe.Ciphertext) (opOut *rlwe.Ciphertext, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ApplyEvaluationKeyNew applies the rlwe.EvaluationKey on op0 and returns the result on a new ciphertext opOut.
func (eval Evaluator) ApplyEvaluationKeyNew(op0 *rlwe.Ciphertext, evk *rlwe.EvaluationKey) (opOut *rlwe.Ciphertext, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RotateNew rotates the columns of op0 by k positions to the left, and returns the result in a newly created element.
// The method will return an error if the evaluator hasn't been given an evaluation key set with the appropriate GaloisKey.
func (eval Evaluator) RotateNew(op0 *rlwe.Ciphertext, k int) (opOut *rlwe.Ciphertext, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Rotate rotates the columns of op0 by k positions to the left and returns the result in opOut.
// The method will return an error if the evaluator hasn't been given an evaluation key set with the appropriate GaloisKey.
func (eval Evaluator) Rotate(op0 *rlwe.Ciphertext, k int, opOut *rlwe.Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// ConjugateNew conjugates op0 (which is equivalent to a row rotation) and returns the result in a newly created element.
// The method will return an error if the evaluator hasn't been given an evaluation key set with the appropriate GaloisKey.
func (eval Evaluator) ConjugateNew(op0 *rlwe.Ciphertext) (opOut *rlwe.Ciphertext, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Conjugate conjugates op0 (which is equivalent to a row rotation) and returns the result in opOut.
// The method will return an error if the evaluator hasn't been given an evaluation key set with the appropriate GaloisKey.
func (eval Evaluator) Conjugate(op0 *rlwe.Ciphertext, opOut *rlwe.Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// RotateHoistedNew takes an input Ciphertext and a list of rotations and returns a map of Ciphertext, where each element of the map is the input Ciphertext
// rotation by one element of the list. It is much faster than sequential calls to [Evaluator.Rotate].
func (eval Evaluator) RotateHoistedNew(ctIn *rlwe.Ciphertext, rotations []int) (opOut map[int]*rlwe.Ciphertext, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RotateHoisted takes an input Ciphertext and a list of rotations and populates a map of pre-allocated Ciphertexts,
// where each element of the map is the input Ciphertext rotation by one element of the list.
// It is much faster than sequential calls to [Evaluator.Rotate].
func (eval Evaluator) RotateHoisted(ctIn *rlwe.Ciphertext, rotations []int, opOut map[int]*rlwe.Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (eval Evaluator) RotateHoistedLazyNew(level int, rotations []int, ct *rlwe.Ciphertext, c2DecompQP []ringqp.Poly) (cOut map[int]*rlwe.Element[ringqp.Poly], err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// InnerSum divides each row of the underlying plaintext in sub-vectors of size batchSize and add n of these together.
//
// WARNING: 0 < n*batchSize <= ctIn.Slots() must divide the number of slots ctIn.Slots(). For other parameters, consider using [Evaluator.RotateAndAdd].
//
// Example for batchSize=2, n=4 and 16 slots (garbage slots are marked as X):
//
// Input:
//
// [{a, b}, {c, d}, {e, f}, {g, h}, {i, j}, {k, l}, {m, n}, {o, p}]
//
// Output:
//
// [{a+c+e+g, b+d+f+h}, {X, X}, {X, X}, {X, X}, {i+k+m+o, j+l+n+p}, {X, X}, {X, X}, {X, X}]
func (eval Evaluator) InnerSum(ctIn *rlwe.Ciphertext, batchSize, n int, opOut *rlwe.Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// RotateAndAdd computes the sum of pt_i, 0 <= i < n, where pt_i is the underlying plaintext rotated ([Evaluator.Rotate]) by batchSize*i slots.
//
// Example: for batchSize=3, n=2, ctIn.Slots()=8:
//
// Input:
//
//	[a, b, c, d, e, f, g, h]
//
// Output:
//
//	[a, b, c, d, e, f, g, h] + [d, e, f, g, h, a, b, c] = [a+d, b+e, c+f, d+g, e+h, f+a, g+b, h+c]
//
// Calling RotateAndAdd(ctIn, 1, n, opOut) can be used to compute the inner sum of the first n slots of a plaintext.
func (eval Evaluator) RotateAndAdd(ctIn *rlwe.Ciphertext, batchSize, n int, opOut *rlwe.Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// WithKey creates a shallow copy of the receiver Evaluator for which the new EvaluationKey is evaluationKey
// and where the temporary buffers are shared. The receiver and the returned Evaluators cannot be used concurrently.
func (eval Evaluator) WithKey(evk rlwe.EvaluationKeySet) *Evaluator {
	_ = "STUB: not implemented"
	return nil
}
