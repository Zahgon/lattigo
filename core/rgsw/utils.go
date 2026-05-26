package rgsw

import (
	"github.com/tuneinsight/lattigo/v6/core/rlwe"
	"github.com/tuneinsight/lattigo/v6/ring"
)

// NoiseRGSWCiphertext returns the log2 of the standard deviation of the noise of each component of the RGSW ciphertext.
// pt must be in the NTT and Montgomery domain
func NoiseRGSWCiphertext(ct *Ciphertext, pt ring.Poly, sk *rlwe.SecretKey, params rlwe.Parameters) (float64, float64) {
	_ = "STUB: not implemented"
	return 0, 0
}
