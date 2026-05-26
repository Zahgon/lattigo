package rgsw

import (
	"github.com/tuneinsight/lattigo/v6/core/rlwe"
)

// Encryptor is a type for encrypting RGSW ciphertexts. It implements the [rlwe.Encryptor]
// interface overriding the [rlwe.Encryptor.Encrypt] and [rlwe.Encryptor.EncryptZero] methods to accept [rgsw.Ciphertext]
// types in addition to ciphertexts types in the rlwe package.
type Encryptor struct {
	*rlwe.Encryptor
	pool *rlwe.BufferPool
}

// NewEncryptor creates a new Encryptor type. Note that only secret-key encryption is
// supported at the moment.
func NewEncryptor(params rlwe.ParameterProvider, key rlwe.EncryptionKey) *Encryptor {
	_ = "STUB: not implemented"
	return nil
}

// Encrypt encrypts a plaintext pt into a ciphertext ct, which can be a [rgsw.Ciphertext]
// or any of the `rlwe` cipheretxt types.
func (enc Encryptor) Encrypt(pt *rlwe.Plaintext, ct interface{}) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Sanity check, this error should not happen.

// EncryptZero generates an encryption of zero into a ciphertext ct, which can be a [rgsw.Ciphertext]
// or any of the `rlwe` ciphertext types.
func (enc Encryptor) EncryptZero(ct interface{}) (err error) { _ = "STUB: not implemented"; return nil }

// extract the RingQ polynomial in case not P moduli have been provided
