package bootstrapping

import (
	"github.com/tuneinsight/lattigo/v6/circuits/ckks/dft"
	"github.com/tuneinsight/lattigo/v6/circuits/ckks/mod1"
	"github.com/tuneinsight/lattigo/v6/core/rlwe"
	"github.com/tuneinsight/lattigo/v6/ring"
	"github.com/tuneinsight/lattigo/v6/schemes/ckks"
)

// Evaluator is a struct to store a memory buffer with the plaintext matrices,
// the polynomial approximation, and the keys for the bootstrapping.
// It is used to evaluate the bootstrapping circuit on single ciphertexts.
type Evaluator struct {
	Parameters
	*ckks.Evaluator
	DFTEvaluator  *dft.Evaluator
	Mod1Evaluator *mod1.Evaluator
	*EvaluationKeys

	ckks.DomainSwitcher

	// [1, x, x^2, x^4, ..., x^N1/2] / (X^N1 +1)
	xPow2N1 []ring.Poly
	// [1, x, x^2, x^4, ..., x^N2/2] / (X^N2 +1)
	xPow2N2 []ring.Poly
	// [1, x^-1, x^-2, x^-4, ..., x^-N1/2] / (X^N1 +1)
	xPow2InvN1 []ring.Poly
	// [1, x^-1, x^-2, x^-4, ..., x^-N2/2] / (X^N2 +1)
	xPow2InvN2 []ring.Poly

	Mod1Parameters mod1.Parameters
	S2CDFTMatrix   dft.Matrix
	C2SDFTMatrix   dft.Matrix

	SkDebug *rlwe.SecretKey

	pool *rlwe.BufferPool
}

// NewEvaluator creates a new [Evaluator].
func NewEvaluator(btpParams Parameters, evk *EvaluationKeys) (eval *Evaluator, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The switch to standard to conjugate invariant multiplies the scale by 2

// CheckKeys checks if all the necessary keys are present in the instantiated [Evaluator]
func (eval Evaluator) checkKeys(evk *EvaluationKeys) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (eval *Evaluator) initialize(btpParams Parameters) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// [-K, K]

// Correcting factor for approximate division by Q
// The second correcting factor for approximate multiplication by Q is included in the coefficients of the EvalMod polynomials

// If the scale used during the EvalMod step is smaller than Q0, then we cannot increase the scale during
// the EvalMod step to get a free division by MessageRatio, and we need to do this division (totally or partly)
// during the CoeffstoSlots step

// Sets qDiv to 1 if there is enough room for the division to happen using scale manipulation.

// CoeffsToSlots vectors
// Change of variable for the evaluation of the Chebyshev polynomial + cancelling factor for the DFT and SubSum + eventual scaling factor for the double angle formula

// For the GC

// Bootstrap bootstraps a single ciphertext and returns the bootstrapped ciphertext.
func (eval Evaluator) Bootstrap(ct *rlwe.Ciphertext) (*rlwe.Ciphertext, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BootstrapMany bootstraps a list of ciphertexts and returns the list of bootstrapped ciphertexts.
func (eval Evaluator) BootstrapMany(cts []rlwe.Ciphertext) ([]rlwe.Ciphertext, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Depth returns the multiplicative depth (number of levels consumed) of the bootstrapping circuit.
func (eval Evaluator) Depth() int { _ = "STUB: not implemented"; return 0 }

// OutputLevel returns the output level after the evaluation of the bootstrapping circuit.
func (eval Evaluator) OutputLevel() int { _ = "STUB: not implemented"; return 0 }

// MinimumInputLevel returns the minimum level at which a ciphertext must be to be bootstrapped.
func (eval Evaluator) MinimumInputLevel() int { _ = "STUB: not implemented"; return 0 }

// Evaluate re-encrypts a ciphertext to a ciphertext at MaxLevel - k where k is the depth of the bootstrapping circuit.
// If the input ciphertext level is zero, the input scale must be an exact power of two smaller than Q[0]/MessageRatio
// (it can't be equal since Q[0] is not a power of two).
// The message ratio is an optional field in the bootstrapping parameters, by default it set to 2^{LogMessageRatio = 8}.
// See the bootstrapping parameters for more information about the message ratio or other parameters related to the bootstrapping.
// If the input ciphertext is at level one or more, the input scale does not need to be an exact power of two as one level
// can be used to do a scale matching.
//
// The circuit consists in 5 steps.
//  1. ScaleDown: scales the ciphertext to q/|m| and bringing it down to q
//  2. ModUp: brings the modulus from q to Q
//  3. CoeffsToSlots: homomorphic encoding
//  4. EvalMod: homomorphic modular reduction
//  5. SlotsToCoeffs: homomorphic decoding
func (eval Evaluator) Evaluate(ctIn *rlwe.Ciphertext) (ctOut *rlwe.Ciphertext, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// [M^{d}/q1 + e^{d-logprec}]

// Stores by how much a ciphertext must be scaled to get back
// to the input scale
// Error correcting factor of the approximate division by q1
// diffScale = ctIn.Scale / (ctOut.Scale * errScale)

// [M^{d} + e^{d-logprec}]

// prec = round(2^{logprec})

// Corrects the last iteration 2^{logprec} such that diffScale / prec * QReserved is as close to an integer as possible.
// This is necessary to not lose bits of precision during the last iteration is a reserved prime is used.
// If this correct is not done, what can happen is that there is a loss of up to 2^{logprec/2} bits from the last iteration.

// 1) Computes the scale = diffScale / prec * QReserved

// 2) Finds the closest integer to scale with scale = round(scale)

// 3) Computes the corrected precision = diffScale * QReserved / round(scale)

// 4) Updates with the corrected precision

// round(q1/logprec)

// Checks that round(q1/logprec) >= 2^{logprec}

// [M^{d} + e^{d-logprec}] - [M^{d}] -> [e^{d-logprec}]

// prec * [e^{d-logprec}] -> [e^{d}]

// [e^{d}] -> [e^{d}/q1] -> [e^{d}/q1 + e'^{d-logprec}]

// [[e^{d}/q1 + e'^{d-logprec}] * q1/logprec -> [e^{d-logprec} + e'^{d-2logprec}*q1]

// Else we compute the floating point ratio

// Do a scaled multiplication by the last prime

// And rescale

// This is a given

// [M^{d} + e^{d-logprec}] - [e^{d-logprec} + e'^{d-2logprec}*q1] -> [M^{d} + e'^{d-2logprec}*q1]

// EvaluateConjugateInvariant takes two ciphertext in the Conjugate Invariant ring, repacks them in a single ciphertext in the standard ring
// using the real and imaginary part, bootstrap both ciphertext, and then extract back the real and imaginary part before repacking them
// individually in two new ciphertexts in the Conjugate Invariant ring.
func (eval Evaluator) EvaluateConjugateInvariant(ctLeftN1Q0, ctRightN1Q0 *rlwe.Ciphertext) (ctLeftN1QL, ctRightN1QL *rlwe.Ciphertext, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Switches ring from ring.ConjugateInvariant to ring.Standard

// Repacks ctRightN1Q0 into the imaginary part of ctLeftN1Q0
// which is zero since it comes from the Conjugate Invariant ring)

// Bootstraps in the ring.Standard

// The SlotsToCoeffs transformation scales the ciphertext by 0.5
// This is done to compensate for the 2x factor introduced by ringStandardToConjugate(*).

// Switches ring from ring.Standard to ring.ConjugateInvariant

// Extracts the imaginary part

// checks if the current message ratio is greater or equal to the last prime times the target message ratio.
func checkMessageRatio(ct *rlwe.Ciphertext, msgRatio float64, r *ring.Ring) bool {
	_ = "STUB: not implemented"
	return false
}

func (eval Evaluator) bootstrap(ctIn *rlwe.Ciphertext) (ctOut *rlwe.Ciphertext, errScale *rlwe.Scale, err error) {
	_ = "STUB: not implemented"

	// Step 1: scale to q/|m|
	return nil, nil, nil
}

// Step 2 : Extend the basis from q to Q

// Step 3 : CoeffsToSlots (Homomorphic encoding)
// ctReal = Ecd(real)
// ctImag = Ecd(imag)
// If n < N/2 then ctReal = Ecd(real||imag)

// Step 4 : EvalMod (Homomorphic modular reduction)

// Step 4 : EvalMod (Homomorphic modular reduction)

// Step 5 : SlotsToCoeffs (Homomorphic decoding)

// ScaleDown brings the ciphertext level to zero and scaling factor to Q[0]/MessageRatio
// It multiplies the ciphertexts by round(currentMessageRatio / targetMessageRatio) where:
//   - currentMessageRatio = Q/ctIn.Scale
//   - targetMessageRatio = q/|m|
//
// and updates the scale of ctIn accordingly
// It then rescales the ciphertext down to q if necessary and also returns the rescaling error from this process
func (eval Evaluator) ScaleDown(ctIn *rlwe.Ciphertext) (*rlwe.Ciphertext, *rlwe.Scale, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Removes unecessary primes

// Current Message Ratio

// Desired Message Ratio

// (Current Message Ratio) / (Desired Message Ratio)

// errScale = CtIn.Scale/(Q[0]/MessageRatio)

// Rescaling error (if any)

// ModUp raise the modulus from q to Q, scales the message  and applies the Trace if the ciphertext is sparsely packed.
func (eval Evaluator) ModUp(ctIn *rlwe.Ciphertext) (ctOut *rlwe.Ciphertext, err error) {
	_ = "STUB: not implemented"

	// Switch to the sparse key
	return nil, nil
}

// Extend the ciphertext from q to Q with zero values.

// ModUp q->Q for ctIn[0] centered around q

// ModUp q->QP for ctIn[1] centered around q

// Scale the message from Q0/|m| to QL/|m|, where QL is the largest modulus used during the bootstrapping.

// Switch back to the dense key

// Scale the message from Q0/|m| to QL/|m|, where QL is the largest modulus used during the bootstrapping.

//SubSum X -> (N/dslots) * Y^dslots

// CoeffsToSlots applies the homomorphic decoding
func (eval Evaluator) CoeffsToSlots(ctIn *rlwe.Ciphertext) (ctReal, ctImag *rlwe.Ciphertext, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// EvalMod applies the homomorphic modular reduction by q.
func (eval Evaluator) EvalMod(ctIn *rlwe.Ciphertext) (ctOut *rlwe.Ciphertext, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// EvalModAndScale applies the homomorphic modular reduction by q and scales the output value (without
// consuming an additional level).
func (eval Evaluator) EvalModAndScale(ctIn *rlwe.Ciphertext, scaling complex128) (ctOut *rlwe.Ciphertext, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (eval Evaluator) SlotsToCoeffs(ctReal, ctImag *rlwe.Ciphertext) (ctOut *rlwe.Ciphertext, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (eval Evaluator) switchRingDegreeN1ToN2New(ctN1 *rlwe.Ciphertext) (ctN2 *rlwe.Ciphertext) {
	_ = "STUB: not implemented"
	return nil
}

// Sanity check, this error should never happen unless this algorithm has been improperly
// modified to pass invalid inputs.

func (eval Evaluator) switchRingDegreeN2ToN1New(ctN2 *rlwe.Ciphertext) (ctN1 *rlwe.Ciphertext) {
	_ = "STUB: not implemented"
	return nil
}

// Sanity check, this error should never happen unless this algorithm has been improperly
// modified to pass invalid inputs.

func (eval Evaluator) ComplexToRealNew(ctCmplx *rlwe.Ciphertext) (ctReal *rlwe.Ciphertext) {
	_ = "STUB: not implemented"
	return nil
}

// Sanity check, this error should never happen unless this algorithm has been improperly
// modified to pass invalid inputs.

func (eval Evaluator) RealToComplexNew(ctReal *rlwe.Ciphertext) (ctCmplx *rlwe.Ciphertext) {
	_ = "STUB: not implemented"
	return nil
}

// Sanity check, this error should never happen unless this algorithm has been improperly
// modified to pass invalid inputs.

// packingContext contains the parameters used when packing (with Pack())
type packingContext struct {
	Params           *ckks.Parameters // Parameters of the ring we are packing to or unpacking from
	LogMaxDimensions ring.Dimensions  // maximum dimension of a packed ciphertext (logMaxDimensions <= params.LogMaxDimensions())
	LogSlots         int              // number of slots in a ct before packing (resp. after unpacking)
	NbPackedCTs      int              // number of cts to be packed (resp. to be unpacked into)
}

// PackAndSwitchN1ToN2 packs the ciphertexts into N1 and switch to N2 if N1 < N2
// then it packs the ciphertexts into N2.
func (eval Evaluator) PackAndSwitchN1ToN2(cts []rlwe.Ciphertext) ([]rlwe.Ciphertext, *packingContext, *packingContext, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

// If N1 < N2, we pack ciphertexts into N1 and then switch to N2

// If the bootstrapping max slots are smaller than the max slots of N1, we only pack up to the former

// Packing ciphertexts into N2 (up to eval.Parameters.LogMaxDimensions())

// UnpackAndSwitchN2ToN1 unpacks the ciphertexts into N2 and, if N1 < N2, it switches the ciphertexts
// to N1 and unpacks further into N1
func (eval Evaluator) UnpackAndSwitchN2ToN1(cts []rlwe.Ciphertext, ctxtN1, ctxtN2 *packingContext) ([]rlwe.Ciphertext, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Unpack ciphertexts in N2

// If N1 != N2 (i.e. ctxtN1 != nil): 1) switch cts to N1 2) unpack the cts in N1

// Set back the dimension of cts to its original value

// unpack unpacks one sparse ciphertext of (log) dimension ctxt.logMaxDimensions
// into ctxt.NbPackedCTs ciphertexts of (log) dimension {0, ctxt.LogSlots}
func (eval Evaluator) unpack(ct *rlwe.Ciphertext, ctxt packingContext, xPow2Inv []ring.Poly) ([]rlwe.Ciphertext, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// log of number of CTs that can be packed in one ct

// #cts to unpack from ct

// log gap of CTs with params.N (minus one)

/* #nosec G115 -- n-1 cannot be negative */

// pack packs ctxt.NbPackedCTs sparse ciphertexts of (log) dimension {0, ctxt.LogSlots}
// into one ciphertext of (log) dimension ctxt.logMaxDimensions
func (eval Evaluator) pack(cts []rlwe.Ciphertext, ctxt packingContext, xPow2 []ring.Poly) ([]rlwe.Ciphertext, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// log of number of CTs that can be packed in one ct
