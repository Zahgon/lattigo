package bignum

import (
	"math/big"
)

// ChebyshevApproximation computes a Chebyshev approximation of the input function, for the range [-a, b] of degree degree.
// f.(type) can be either :
//   - func(Complex128)Complex128
//   - func(float64)float64
//   - func(*big.Float)*big.Float
//   - func(*Complex)*Complex
//
// The reference precision is taken from the values stored in the Interval struct.
func ChebyshevApproximation(f interface{}, interval Interval) (pol Polynomial) {
	_ = "STUB: not implemented"
	return *new(Polynomial)
}

func chebyshevNodes(n int, interval Interval) (nodes []*big.Float) {
	_ = "STUB: not implemented"
	return nil
}

func chebyCoeffs(nodes []*big.Float, fi []*Complex, interval Interval) (coeffs []*Complex) {
	_ = "STUB: not implemented"
	return nil
}

func chebyshevBasisInPlace(deg int, x *big.Float, inter Interval, poly []*big.Float) {
	_ = "STUB: not implemented"
	return
}

// u = (2*x - (a+b))/(b-a)
