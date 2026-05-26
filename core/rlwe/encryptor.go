package rlwe

import (
	"github.com/tuneinsight/lattigo/v6/ring"
	"github.com/tuneinsight/lattigo/v6/ring/ringqp"
	"github.com/tuneinsight/lattigo/v6/utils/sampling"
)

// EncryptionKey is an interface for encryption keys. Valid encryption
// keys are the [SecretKey] and [PublicKey] types.
type EncryptionKey interface {
	isEncryptionKey()
}

// NewEncryptor creates a new [Encryptor] from either a public key or a private key.
func NewEncryptor(params ParameterProvider, key EncryptionKey) *Encryptor {
	_ = "STUB: not implemented"
	return nil
}

// Sanity check

// Sanity check, this error should not happen.

type Encryptor struct {
	params Parameters

	encKey         EncryptionKey
	prng           sampling.PRNG
	xeSampler      ring.Sampler
	xsSampler      ring.Sampler
	basisextender  *ring.BasisExtender
	uniformSampler ringqp.UniformSampler
	pool           *BufferPool
}

// GetRLWEParameters returns the underlying [Parameters].
func (enc Encryptor) GetRLWEParameters() *Parameters { _ = "STUB: not implemented"; return nil }

func newEncryptor(params Parameters) *Encryptor { _ = "STUB: not implemented"; return nil }

// Sanity check, this error should not happen.

// Sanity check, this error should not happen.

// newTestEncryptorWithKeyedPRNG creates a new [Encryptor] that uses the provided prng for randomness.
// CAUTION: THIS FUNCTION SHOULD BE USED FOR TESTING PURPOSES ONLY.
// WARNING: The resulting encryptor is not meant to be used concurrently.
func newTestEncryptorWithKeyedPRNG(params ParameterProvider, key EncryptionKey, prng *sampling.KeyedPRNG) *Encryptor {
	_ = "STUB: not implemented"
	return nil
}

// Encrypt encrypts the input plaintext using the stored encryption key and writes the result on ct.
// The method currently accepts only *[Ciphertext] as ct.
// If a [Plaintext] is given, then the output [Ciphertext] [MetaData] will match the [Plaintext] [MetaData].
// The method returns an error if the ct has an unsupported type or if no encryption key is stored
// in the [Encryptor].
//
// The encryption procedure masks the plaintext by adding a fresh encryption of zero.
// The encryption procedure depends on the parameters: If the auxiliary modulus P is defined, the
// encryption of zero is sampled in QP before being rescaled by P; otherwise, it is directly sampled in Q.
func (enc Encryptor) Encrypt(pt *Plaintext, ct interface{}) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// EncryptNew encrypts the input plaintext using the stored encryption key and returns a newly
// allocated [Ciphertext] containing the result.
// If a [Plaintext] is provided, then the output [Ciphertext] [MetaData] will match the [Plaintext] [MetaData].
// The method returns an error if the ct has an unsupported type or if no encryption key is stored
// in the [Encryptor].
//
// The encryption procedure masks the plaintext by adding a fresh encryption of zero.
// The encryption procedure depends on the parameters: If the auxiliary modulus P is defined, the
// encryption of zero is sampled in QP before being rescaled by P; otherwise, it is directly sampled in Q.
func (enc Encryptor) EncryptNew(pt *Plaintext) (ct *Ciphertext, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// EncryptZero generates an encryption of zero under the stored encryption key and writes the result on ct.
// The method accepts only *[Ciphertext] as input.
// The method returns an error if the ct has an unsupported type or if no encryption key is stored
// in the [Encryptor].
//
// The encryption procedure depends on the parameters: If the auxiliary modulus P is defined, the
// encryption of zero is sampled in QP before being rescaled by P; otherwise, it is directly sampled in Q.
// The zero encryption is generated according to the given [Ciphertext] [MetaData].
func (enc Encryptor) EncryptZero(ct interface{}) (err error) { _ = "STUB: not implemented"; return nil }

// EncryptZeroNew generates an encryption of zero under the stored encryption key and returns a newly
// allocated [Ciphertext] containing the result.
// The method returns an error if no encryption key is stored in the [Encryptor].
// The encryption procedure depends on the parameters: If the auxiliary modulus P is defined, the
// encryption of zero is sampled in QP before being rescaled by P; otherwise, it is directly sampled in Q.
func (enc Encryptor) EncryptZeroNew(level int) (ct *Ciphertext) {
	_ = "STUB: not implemented"
	return nil
}

// Sanity check, this error should not happen.

func (enc Encryptor) encryptZeroPk(pk *PublicKey, ct interface{}) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// We sample a RLWE instance (encryption of zero) over the extended ring (ciphertext ring + special prime)

// (#Q + #P) NTT

// ct0 = u*pk0
// ct1 = u*pk1

// 2*(#Q + #P) NTT

// ct0 = (u*pk0 + e0)/P

// ct1 = (u*pk1 + e1)/P

func (enc Encryptor) encryptZeroPkNoP(pk *PublicKey, ct Element[ring.Poly]) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// ct0 = NTT(u*pk0)

// ct1 = NTT(u*pk1)

// c0

// c1

// encryptZeroSk generates an encryption of zero using the stored secret-key and writes the result on ct.
// The method accepts only *rlwe.Ciphertext or *rgsw.Ciphertext as input and will return an error otherwise.
// The zero encryption is generated according to the given Ciphertext MetaData.
func (enc Encryptor) encryptZeroSk(sk *SecretKey, ct interface{}) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// ct = (e, a)

// encryptZeroSkFromC1 computes c0 := c1 * sk + noise according to the ciphertext metadata
func (enc Encryptor) encryptZeroSkFromC1(sk *SecretKey, ct Element[ring.Poly], c1, c0 ring.Poly) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// encryptZeroSkFromC1QP computes c0 := c1 * sk + noise according to the provided levels and ciphertext metadata
func (enc Encryptor) encryptZeroSkFromC1QP(sk *SecretKey, levelQ, levelP int, ctMetaData *CiphertextMetaData, c1, c0 ringqp.Poly) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// ct = (e, 0)

// sk is in Montgomery form by default and c1 can be interpreted as being in the Montgomery domain or not (it is uniformly generated).
// Hence, the domain of the noise e determines whether c0 = c1*sk + e is in the Montgomery form or not.

// (-a*sk + e, a)

// withKeyedUniformSampling returns this encryptor with a keyed prng as its source of randomness for the uniform
// element c1.
// The returned encryptor is not thread safe (sampling will not be deterministic).
func (enc Encryptor) withKeyedUniformSampling(prng *sampling.KeyedPRNG) *Encryptor {
	_ = "STUB: not implemented"
	return nil
}

func (enc Encryptor) WithKey(key EncryptionKey) *Encryptor { _ = "STUB: not implemented"; return nil }

// Sanity check, this error should not happen.

// Sanity check, this error should not happen.

// Sanity check, this error should not happen.

// checkPk checks that a given pk is correct for the parameters.
func (enc Encryptor) checkPk(pk *PublicKey) (err error) { _ = "STUB: not implemented"; return nil }

// checkPk checks that a given pk is correct for the parameters.
func (enc Encryptor) checkSk(sk *SecretKey) (err error) { _ = "STUB: not implemented"; return nil }

func (enc Encryptor) addPtToCt(level int, pt *Plaintext, ct *Ciphertext) {
	_ = "STUB: not implemented"
	return
}
