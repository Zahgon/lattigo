package bignum

import (
	"math/big"
)

// MonomialEval evaluates y = sum x^i * poly[i].
func MonomialEval(x *big.Float, poly []*big.Float) (y *big.Float) {
	_ = "STUB: not implemented"
	return nil
}

// ChebyshevEval evaluates y = sum Ti(x) * poly[i], where T0(x) = 1, T1(x) = (2x-a-b)/(b-a) and T{i+j}(x) = 2TiTj(x)- T|i-j|(x).
func ChebyshevEval(x *big.Float, poly []*big.Float, inter Interval) (y *big.Float) {
	_ = "STUB: not implemented"
	return nil
}

// u = (2*x - (a+b))/(b-a)
