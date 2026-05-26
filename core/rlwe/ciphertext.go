package rlwe

import (
	"github.com/tuneinsight/lattigo/v6/ring"
	"github.com/tuneinsight/lattigo/v6/utils/sampling"
)

// Ciphertext is a generic type for RLWE ciphertexts.
type Ciphertext struct {
	Element[ring.Poly]
}

// NewCiphertext returns a new [Ciphertext] with zero values and an associated
// MetaData set to the Parameters default value.
func NewCiphertext(params ParameterProvider, degree int, level ...int) (ct *Ciphertext) {
	_ = "STUB: not implemented"
	return nil
}

// NewCiphertextAtLevelFromPoly constructs a new [Ciphertext] at a specific level
// where the message is set to the passed poly. No checks are performed on poly and
// the returned [Ciphertext] will share its backing array of coefficients.
// Returned [Ciphertext]'s MetaData is allocated but empty.
func NewCiphertextAtLevelFromPoly(level int, poly []ring.Poly) (*Ciphertext, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewCiphertextRandom generates a new uniformly distributed [Ciphertext] of degree, level.
func NewCiphertextRandom(prng sampling.PRNG, params ParameterProvider, degree, level int) (ciphertext *Ciphertext) {
	_ = "STUB: not implemented"
	return nil
}

// Plaintext casts the target ciphertext into a plaintext type.
// This method is allocation free.
func (ct Ciphertext) Plaintext() *Plaintext { _ = "STUB: not implemented"; return nil }

// CopyNew creates a new element as a copy of the target element.
func (ct Ciphertext) CopyNew() *Ciphertext { _ = "STUB: not implemented"; return nil }

// Copy copies the input element and its parameters on the target element.
func (ct Ciphertext) Copy(ctxCopy *Ciphertext) { _ = "STUB: not implemented"; return }

// Equal performs a deep equal.
func (ct Ciphertext) Equal(other *Ciphertext) bool { _ = "STUB: not implemented"; return false }
