package rlwe

import (
	"github.com/tuneinsight/lattigo/v6/ring"
	"github.com/tuneinsight/lattigo/v6/ring/ringqp"
)

// KeyGenerator is a structure that stores the elements required to create new keys,
// as well as a memory buffer for intermediate values.
type KeyGenerator struct {
	*Encryptor
	pool *BufferPool
}

// NewKeyGenerator creates a new KeyGenerator, from which the secret and public keys, as well as [EvaluationKey].
func NewKeyGenerator(params ParameterProvider) *KeyGenerator { _ = "STUB: not implemented"; return nil }

// GenSecretKeyNew generates a new [SecretKey].
// Distribution is set according to [rlwe.Parameters.HammingWeight].
func (kgen KeyGenerator) GenSecretKeyNew() (sk *SecretKey) { _ = "STUB: not implemented"; return nil }

// GenSecretKey generates a [SecretKey].
// Distribution is set according to [rlwe.Parameters.HammingWeight].
func (kgen KeyGenerator) GenSecretKey(sk *SecretKey) { _ = "STUB: not implemented"; return }

// GenSecretKeyWithHammingWeightNew generates a new [SecretKey] with exactly hw non-zero coefficients.
func (kgen *KeyGenerator) GenSecretKeyWithHammingWeightNew(hw int) (sk *SecretKey) {
	_ = "STUB: not implemented"
	return nil
}

// GenSecretKeyWithHammingWeight generates a [SecretKey] with exactly hw non-zero coefficients.
func (kgen KeyGenerator) GenSecretKeyWithHammingWeight(hw int, sk *SecretKey) {
	_ = "STUB: not implemented"
	return
}

// Sanity check, this error should not happen.

func (kgen KeyGenerator) genSecretKeyFromSampler(sampler ring.Sampler, sk *SecretKey) {
	_ = "STUB: not implemented"
	return
}

// GenPublicKeyNew generates a new [PublicKey] from the provided [SecretKey].
func (kgen KeyGenerator) GenPublicKeyNew(sk *SecretKey) (pk *PublicKey) {
	_ = "STUB: not implemented"
	return nil
}

// GenPublicKey generates a [PublicKey] from the provided [SecretKey].
func (kgen KeyGenerator) GenPublicKey(sk *SecretKey, pk *PublicKey) {
	_ = "STUB: not implemented"
	return
}

// Sanity check, this error should not happen.

// GenKeyPairNew generates a new [SecretKey] and a corresponding [PublicKey].
// Distribution of the [SecretKey] set according to [rlwe.Parameters.HammingWeight].
func (kgen KeyGenerator) GenKeyPairNew() (sk *SecretKey, pk *PublicKey) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GenRelinearizationKeyNew generates a new [EvaluationKey] that will be used to relinearize [Ciphertexts] during multiplication.
func (kgen KeyGenerator) GenRelinearizationKeyNew(sk *SecretKey, evkParams ...EvaluationKeyParameters) (rlk *RelinearizationKey) {
	_ = "STUB: not implemented"
	return nil
}

// GenRelinearizationKey generates an [EvaluationKey] that will be used to relinearize [Ciphertexts] during multiplication.
func (kgen KeyGenerator) GenRelinearizationKey(sk *SecretKey, rlk *RelinearizationKey) {
	_ = "STUB: not implemented"
	return
}

// GenGaloisKeyNew generates a new [GaloisKey], enabling the automorphism X^{i} -> X^{i * galEl}.
func (kgen KeyGenerator) GenGaloisKeyNew(galEl uint64, sk *SecretKey, evkParams ...EvaluationKeyParameters) (gk *GaloisKey) {
	_ = "STUB: not implemented"
	return nil
}

// GenGaloisKey generates a GaloisKey, enabling the automorphism X^{i} -> X^{i * galEl}.
func (kgen KeyGenerator) GenGaloisKey(galEl uint64, sk *SecretKey, gk *GaloisKey) {
	_ = "STUB: not implemented"
	return
}

// We encrypt [-a * pi_{k^-1}(sk) + sk, a]
// This enables to first apply the gadget product, re-encrypting
// a ciphetext from sk to pi_{k^-1}(sk) and then we apply pi_{k}
// on the ciphertext.

// Sanity check, this error should not happen unless the
// evaluator's buffer thave been improperly tempered with.

// GenGaloisKeys generates the [GaloisKey] objects for all galois elements in galEls, and stores
// the resulting key for galois element i in gks[i].
// The galEls and gks parameters must have the same length.
func (kgen KeyGenerator) GenGaloisKeys(galEls []uint64, sk *SecretKey, gks []*GaloisKey) {
	_ = "STUB: not implemented"

	// Sanity check
	return
}

// GenGaloisKeysNew generates the [GaloisKey] objects for all galois elements in galEls, and
// returns the resulting keys in a newly allocated []*[GaloisKey].
func (kgen KeyGenerator) GenGaloisKeysNew(galEls []uint64, sk *SecretKey, evkParams ...EvaluationKeyParameters) (gks []*GaloisKey) {
	_ = "STUB: not implemented"
	return nil
}

// GenEvaluationKeysForRingSwapNew generates the necessary evaluation keys to switch from a standard ring to to a conjugate invariant ring and vice-versa.
func (kgen KeyGenerator) GenEvaluationKeysForRingSwapNew(skStd, skConjugateInvariant *SecretKey, evkParams ...EvaluationKeyParameters) (stdToci, ciToStd *EvaluationKey) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GenEvaluationKeyNew generates a new [EvaluationKey], that will re-encrypt a [Ciphertext] encrypted under the input key into the output key.
// If the ringDegree(skOutput) > ringDegree(skInput),  generates [-a*SkOut + w*P*skIn_{Y^{N/n}} + e, a] in X^{N}.
// If the ringDegree(skOutput) < ringDegree(skInput),  generates [-a*skOut_{Y^{N/n}} + w*P*skIn + e_{N}, a_{N}] in X^{N}.
// Else generates [-a*skOut + w*P*skIn + e, a] in X^{N}.
// The output [EvaluationKey] is always given in max(N, n).
// If evkParams is void, the output [EvaluationKey] will be in the moduli given by the parameters stored in the [KeyGenerator] kgen.
// When re-encrypting a [Ciphertext] from Y^{N/n} to X^{N}, the Ciphertext must first be mapped to X^{N}
// using [SwitchCiphertextRingDegreeNTT](ctSmallDim, nil, ctLargeDim).
// When re-encrypting a [Ciphertext] from X^{N} to Y^{N/n}, the output of the re-encryption is in still X^{N} and
// must be mapped Y^{N/n} using [SwitchCiphertextRingDegreeNTT](ctLargeDim, ringQLargeDim, ctSmallDim).
func (kgen KeyGenerator) GenEvaluationKeyNew(skInput, skOutput *SecretKey, evkParams ...EvaluationKeyParameters) (evk *EvaluationKey) {
	_ = "STUB: not implemented"
	return nil
}

// GenEvaluationKey generates an [EvaluationKey], that will re-encrypt a [Ciphertext] encrypted under the input key into the output key.
// If the ringDegree(skOutput) > ringDegree(skInput),  generates [-a*SkOut + w*P*skIn_{Y^{N/n}} + e, a] in X^{N}.
// If the ringDegree(skOutput) < ringDegree(skInput),  generates [-a*skOut_{Y^{N/n}} + w*P*skIn + e_{N}, a_{N}] in X^{N}.
// Else generates [-a*skOut + w*P*skIn + e, a] in X^{N}.
// The output [EvaluationKey] is always given in max(N, n) and in the moduli of the output [EvaluationKey].
// When re-encrypting a [Ciphertext] from Y^{N/n} to X^{N}, the [Ciphertext] must first be mapped to X^{N}
// using [SwitchCiphertextRingDegreeNTT](ctSmallDim, nil, ctLargeDim).
// When re-encrypting a [Ciphertext] from X^{N} to Y^{N/n}, the output of the re-encryption is in still X^{N} and
// must be mapped Y^{N/n} using [SwitchCiphertextRingDegreeNTT](ctLargeDim, ringQLargeDim, ctSmallDim).
func (kgen KeyGenerator) GenEvaluationKey(skInput, skOutput *SecretKey, evk *EvaluationKey) {
	_ = "STUB: not implemented"
	return
}

// Maps the smaller key to the largest with Y = X^{N/n}.

// Extends the modulus P of skOutput to the one of skInput

// Maps the smaller key to the largest dimension with Y = X^{N/n}.

func (kgen KeyGenerator) genEvaluationKey(skIn ring.Poly, skOut ringqp.Poly, evk *EvaluationKey) {
	_ = "STUB: not implemented"
	return
}

// For a compressed evaluation key, a seed is created and stored in the EvaluationKey struct
// struct while an uncompressed key uses an ephemeral seed.

// Samples an encryption of zero for each element of the EvaluationKey.

// evk[i][j] = (-a*sk + e, a) if the degree of degree of the GadgetCiphertext is 1
// evk[i][j] = (-a*sk + e) if the degree is 0

// Adds the plaintext (input-key) to the EvaluationKey.

// Sanity check, this error should not happen.
