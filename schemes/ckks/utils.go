package ckks

import (
	"math/big"

	"github.com/tuneinsight/lattigo/v6/core/rlwe"
	"github.com/tuneinsight/lattigo/v6/ring"
	"github.com/tuneinsight/lattigo/v6/utils/bignum"
)

// GetRootsBigComplex returns the roots e^{2*pi*i/m *j} for 0 <= j <= NthRoot
// with prec bits of precision.
func GetRootsBigComplex(NthRoot int, prec uint) (roots []*bignum.Complex) {
	_ = "STUB: not implemented"
	return nil
}

// GetRootsComplex128 returns the roots e^{2*pi*i/m *j} for 0 <= j <= NthRoot.
func GetRootsComplex128(NthRoot int) (roots []complex128) { _ = "STUB: not implemented"; return nil }

// StandardDeviation computes the scaled standard deviation of the input vector.
func StandardDeviation(vec interface{}, scale rlwe.Scale) (std float64) {
	_ = "STUB: not implemented"
	return 0
}

// We assume that the error is centered around zero

// Complex128ToFixedPointCRT encodes a vector of complex128 on a CRT polynomial.
// The real part is put in a left N/2 coefficient and the imaginary in the right N/2 coefficients.
func Complex128ToFixedPointCRT(r *ring.Ring, values []complex128, scale float64, coeffs [][]uint64) {
	_ = "STUB: not implemented"
	return
}

// Float64ToFixedPointCRT encodes a vector of floats on a CRT polynomial.
func Float64ToFixedPointCRT(r *ring.Ring, values []float64, scale float64, coeffs [][]uint64) {
	_ = "STUB: not implemented"
	return
}

// SingleFloat64ToFixedPointCRT encodes a single float64 on a CRT polynomialon in the i-th coefficient.
func SingleFloat64ToFixedPointCRT(r *ring.Ring, i int, value float64, scale float64, coeffs [][]uint64) {
	_ = "STUB: not implemented"
	return
}

func ComplexArbitraryToFixedPointCRT(r *ring.Ring, values []*bignum.Complex, scale *big.Float, coeffs [][]uint64) {
	_ = "STUB: not implemented"
	return
}

func BigFloatToFixedPointCRT(r *ring.Ring, values []*big.Float, scale *big.Float, coeffs [][]uint64) {
	_ = "STUB: not implemented"
	return
}
