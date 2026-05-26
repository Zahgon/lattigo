package ckks

import (
	"github.com/tuneinsight/lattigo/v6/core/rlwe"
	"github.com/tuneinsight/lattigo/v6/ring"
)

// DomainSwitcher is a type for switching between the standard CKKS domain (which encrypts vectors of complex numbers)
// and the conjugate invariant variant of CKKS (which encrypts vectors of real numbers).
type DomainSwitcher struct {
	stdRingQ, conjugateRingQ *ring.Ring

	stdToci, ciToStd  *rlwe.EvaluationKey
	automorphismIndex []uint64
	poolQ             *ring.BufferPool
}

// NewDomainSwitcher instantiate a new [DomainSwitcher] type. It may be instantiated from parameters from either RingType.
// The method returns an error if the parameters cannot support the switching (e.g., the NTTs are undefined for
// either of the two ring types).
// The comlexToRealEvk and comlexToRealEvk EvaluationKeys can be generated using [rlwe.KeyGenerator.GenEvaluationKeysForRingSwap].
func NewDomainSwitcher(params Parameters, comlexToRealEvk, realToComplexEvk *rlwe.EvaluationKey) (DomainSwitcher, error) {
	_ = "STUB: not implemented"
	return *new(DomainSwitcher), nil
}

// Sanity check, this error should not happen unless the
// algorithm has been modified to provide invalid inputs.

// ComplexToReal switches the provided ciphertext ctIn from the standard domain to the conjugate
// invariant domain and writes the result into opOut.
// Given ctInCKKS = enc(real(m) + imag(m)) in Z[X](X^N + 1), returns opOutCI = enc(real(m))
// in Z[X+X^-1]/(X^N + 1) in compressed form (N/2 coefficients).
// The scale of the output ciphertext is twice the scale of the input one.
// Requires the ring degree of opOut to be half the ring degree of ctIn.
// The security is changed from Z[X]/(X^N+1) to Z[X]/(X^N/2+1).
// The method will return an error if the DomainSwitcher was not initialized with a the appropriate EvaluationKeys.
func (switcher DomainSwitcher) ComplexToReal(eval *Evaluator, ctIn, opOut *rlwe.Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// RealToComplex switches the provided ciphertext ctIn from the conjugate invariant domain to the
// standard domain and writes the result into opOut.
// Given ctInCI = enc(real(m)) in Z[X+X^-1]/(X^2N+1) in compressed form (N coefficients), returns
// opOutCKKS = enc(real(m) + imag(0)) in Z[X]/(X^2N+1).
// Requires the ring degree of opOut to be twice the ring degree of ctIn.
// The security is changed from Z[X]/(X^N+1) to Z[X]/(X^2N+1).
// The method will return an error if the [DomainSwitcher] was not initialized with a the appropriate EvaluationKeys.
func (switcher DomainSwitcher) RealToComplex(eval *Evaluator, ctIn, opOut *rlwe.Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Switches the RCKswitcher key [X+X^-1] to a CKswitcher key [X]
