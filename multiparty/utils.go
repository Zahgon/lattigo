package multiparty

import (
	"github.com/tuneinsight/lattigo/v6/core/rlwe"
)

// NoiseRelinearizationKey returns the standard deviation of the noise of each individual elements in the collective RelinearizationKey.
func NoiseRelinearizationKey(params rlwe.Parameters, nbParties int) (std float64) {
	_ = "STUB: not implemented"

	// rlk noise = [s*e0 + u*e1 + e2 + e3]
	//
	// s  = sum(s_i)
	// u  = sum(u_i)
	// e0 = sum(e_i0)
	// e1 = sum(e_i1)
	// e2 = sum(e_i2)
	// e3 = sum(e_i3)
	return 0
}

// var(sk) and var(u)
// var(e0), var(e1), var(e2), var(e3)

// var([s*e0 + u*e1 + e2 + e3]) = H*e + H*e + e + e = e(2H+2) = 2e(H+1)

// NoiseEvaluationKey returns the standard deviation of the noise of each individual elements in a collective EvaluationKey.
func NoiseEvaluationKey(params rlwe.Parameters, nbParties int) (std float64) {
	_ = "STUB: not implemented"
	return 0
}

// NoiseGaloisKey returns the standard deviation of the noise of each individual elements in a collective GaloisKey.
func NoiseGaloisKey(params rlwe.Parameters, nbParties int) (std float64) {
	_ = "STUB: not implemented"
	return 0
}

// NoiseKeySwitch returns the standard deviation of the noise of a ciphertext after the KeySwitch protocol
func NoiseKeySwitch(params rlwe.Parameters, nbParties int, noisect, noiseflood float64) (std float64) {
	_ = "STUB: not implemented"
	// #Parties * (noiseflood + noiseFreshSK) + noise ct
	return 0
}

func NoisePublicKeySwitch(params rlwe.Parameters, nbParties int, noisect, noiseflood float64) (std float64) {
	_ = "STUB: not implemented"
	// #Parties * (var(freshZeroPK) + var(noiseFlood)) + noise ct
	return 0
}

func noiseDecryptWithSmudging(nbParties int, noisect, noisefresh, noiseflood float64) (std float64) {
	_ = "STUB: not implemented"
	return 0
}
