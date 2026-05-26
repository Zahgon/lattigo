package ckks

import (
	"math/big"

	"github.com/tuneinsight/lattigo/v6/ring"
	"github.com/tuneinsight/lattigo/v6/utils/bignum"
)

func bigComplexToRNSScalar(r *ring.Ring, scale *big.Float, cmplx *bignum.Complex) (RNSReal, RNSImag ring.RNSScalar) {
	_ = "STUB: not implemented"
	return *new(ring.RNSScalar), *new(ring.RNSScalar)
}

// Divides x by n, returns a float.
func scaleDown(coeff *big.Int, n float64) (x float64) { _ = "STUB: not implemented"; return 0 }
