package rlwe

import (
	"github.com/tuneinsight/lattigo/v6/ring"
)

// Decryptor is a structure used to decrypt [Ciphertext]. It stores the secret-key.
type Decryptor struct {
	params Parameters
	ringQ  *ring.Ring
	poolQ  *ring.BufferPool
	sk     *SecretKey
}

// NewDecryptor instantiates a new generic RLWE [Decryptor].
func NewDecryptor(params ParameterProvider, sk *SecretKey) *Decryptor {
	_ = "STUB: not implemented"
	return nil
}

// GetRLWEParameters returns the underlying [Parameters].
func (d Decryptor) GetRLWEParameters() *Parameters {
	_ = "STUB: not implemented"

	// DecryptNew decrypts the [Ciphertext] and returns the result in a new [Plaintext].
	// Output pt [MetaData] will match the input ct [MetaData].
	return nil
}

func (d Decryptor) DecryptNew(ct *Ciphertext) (pt *Plaintext) {
	_ = "STUB: not implemented"
	return nil
}

// Decrypt decrypts the [Ciphertext] and writes the result in pt.
// The level of the output [Plaintext] is min(ct.Level(), pt.Level())
// Output pt [MetaData] will match the input ct [MetaData].
func (d Decryptor) Decrypt(ct *Ciphertext, pt *Plaintext) { _ = "STUB: not implemented"; return }

// WithKey creates a shallow copy of [Decryptor] with a new decryption key, in which all the
// data-structures are shared with the receiver.
// The receiver and the returned [Decryptor] can be used concurrently.
func (d Decryptor) WithKey(sk *SecretKey) *Decryptor { _ = "STUB: not implemented"; return nil }
