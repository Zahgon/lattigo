package rlwe

import (
	"math/big"

	"github.com/tuneinsight/lattigo/v6/ring"
)

// NoisePublicKey returns the log2 of the standard deviation of the input [PublicKey] with respect to the given [Secret] and parameters.
func NoisePublicKey(pk *PublicKey, sk *SecretKey, params Parameters) float64 {
	_ = "STUB: not implemented"
	return 0
}

// [-as + e] + [as]

// NoiseRelinearizationKey the log2 of the standard deviation of the noise of the input [RelinearizationKey] with respect to the given [Secret] and paramters.
func NoiseRelinearizationKey(rlk *RelinearizationKey, sk *SecretKey, params Parameters) float64 {
	_ = "STUB: not implemented"
	return 0
}

// NoiseGaloisKey the log2 of the standard deviation of the noise of the input [GaloisKey] key with respect to the given [SecretKey] and paramters.
func NoiseGaloisKey(gk *GaloisKey, sk *SecretKey, params Parameters) float64 {
	_ = "STUB: not implemented"
	return 0
}

// NoiseGadgetCiphertext returns the log2 of the standard deviation of the noise of the input [GadgetCiphertext] with respect to the given [Plaintext], [SecretKey] and [Parameters].
// The polynomial pt is expected to be in the NTT and Montgomery domain.
func NoiseGadgetCiphertext(gct *GadgetCiphertext, pt ring.Poly, sk *SecretKey, params Parameters) float64 {
	_ = "STUB: not implemented"
	return 0
}

// required else the check becomes very complicated

// Decrypts
// [-asIn + w*P*sOut + e, a] + [asIn]

// Sums all bases together (equivalent to multiplying with CRT decomposition of 1)
// sum([1]_w * [RNS*PW2*P*sOut + e]) = PWw*P*sOut + sum(e)
// RNS decomp

// PW2 decomp

// sOut * P

// P*s^i + sum(e) - P*s^i = sum(e)

// Checks that the error is below the bound
// Worst error bound is N * floor(6*sigma) * #Keys

// sOut * P * PW2

// NoiseEvaluationKey the log2 of the standard deviation of the noise of the input [GaloisKey] with respect to the given [SecretKey] and [Parameters].
func NoiseEvaluationKey(evk *EvaluationKey, skIn, skOut *SecretKey, params Parameters) float64 {
	_ = "STUB: not implemented"
	return 0
}

// Norm returns the log2 of the standard deviation, minimum and maximum absolute norm of
// the decrypted [Ciphertext], before the decoding (i.e. including the error).
func Norm(ct *Ciphertext, dec *Decryptor) (std, min, max float64) {
	_ = "STUB: not implemented"
	return 0, 0, 0
}

func NormStats(vec []*big.Int) (float64, float64, float64) {
	_ = "STUB: not implemented"
	return 0, 0, 0
}

// NTTSparseAndMontgomery takes a polynomial Z[Y] outside of the NTT domain and maps it to a polynomial Z[X] in the NTT domain where Y = X^(gap).
// This method is used to accelerate the NTT of polynomials that encode sparse polynomials.
func NTTSparseAndMontgomery(r *ring.Ring, metadata *MetaData, pol ring.Poly) {
	_ = "STUB: not implemented"
	return
}

// NTT in dimension n but with roots of N
// This is a small hack to perform at reduced cost an NTT of dimension N on a vector in Y = X^{N/n}, i.e. sparse polynomials.

// Maps NTT in dimension n to NTT in dimension N

// ExtendBasisSmallNormAndCenterNTTMontgomery extends a small-norm polynomial polQ in R_Q to a polynomial
// polP in R_P.
// This method can be used to extend from Q0 to QL.
// Input and output are in the NTT and Montgomery domain.
func ExtendBasisSmallNormAndCenterNTTMontgomery(rQ, rP *ring.Ring, polQ, buff, polP ring.Poly) {
	_ = "STUB: not implemented"
	return
}

// Switches Q[0] out of the NTT and Montgomery domain.

// Reconstruct P from Q
