package ckks

import (
	"testing"

	"github.com/tuneinsight/lattigo/v6/core/rlwe"
	"github.com/tuneinsight/lattigo/v6/utils/bignum"
)

// PrecisionStats is a struct storing statistic about the precision of a CKKS plaintext
type PrecisionStats struct {
	MINLog2Prec Stats
	MAXLog2Prec Stats
	AVGLog2Prec Stats
	MEDLog2Prec Stats
	STDLog2Prec Stats

	MINLog2Err Stats
	MAXLog2Err Stats
	AVGLog2Err Stats
	MEDLog2Err Stats
	STDLog2Err Stats

	Log2Scale float64

	RealDist, ImagDist, L2Dist []struct {
		Prec  float64
		Count int
	}

	cdfResol int
}

// Stats is a struct storing the real, imaginary and L2 norm (modulus)
// about the precision of a complex value.
type Stats struct {
	Real, Imag, L2 float64
}

func (prec PrecisionStats) String() string { _ = "STUB: not implemented"; return "" }

// GetPrecisionStats generates a [PrecisionStats] struct from the reference values and the decrypted values
// vWant.(type) must be either []complex128 or []float64
// element.(type) must be either *Plaintext, *Ciphertext, []complex128 or []float64. If not *Ciphertext, then decryptor can be nil.
func GetPrecisionStats(params Parameters, encoder *Encoder, decryptor *rlwe.Decryptor, want, have interface{}, logprec float64, computeDCF bool) (prec PrecisionStats) {
	_ = "STUB: not implemented"
	return *new(PrecisionStats)
}

func VerifyTestVectors(params Parameters, encoder *Encoder, decryptor *rlwe.Decryptor, valuesWant, valuesHave interface{}, log2MinPrec int, logprec float64, printPrecisionStats bool, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

// Z[X]/(X^{N} + 1)

// Z[X + X^1]/(X^{2N} + 1)

func getPrecisionStats(params Parameters, encoder *Encoder, decryptor *rlwe.Decryptor, want, have interface{}, logprec float64, computeDCF bool) (prec PrecisionStats) {
	_ = "STUB: not implemented"
	return *new(PrecisionStats)
}

func getRawVectors(params Parameters, encoder *Encoder, decryptor *rlwe.Decryptor, want, have interface{}, logprec float64) (valuesWant, valuesHave []*bignum.Complex) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Sanity check, this error should never happen.

// Sanity check, this error should never happen.

func (prec *PrecisionStats) calcCDF(precs []float64, res []struct {
	Prec  float64
	Count int
}) {
	_ = "STUB: not implemented"
	return
}

func calcmedian(values []float64) (median float64) { _ = "STUB: not implemented"; return 0 }
