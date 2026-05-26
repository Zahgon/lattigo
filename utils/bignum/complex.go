package bignum

import (
	"math/big"
)

// Complex is a type for arbitrary precision complex number
type Complex [2]*big.Float

// NewComplex creates a new arbitrary precision complex number
func NewComplex() (c *Complex) { _ = "STUB: not implemented"; return nil }

// ToComplex takes a complex128, float64, int, int64, uint, uint64, *big.Int, *big.Float or *Complex and returns a *Complex set to the given precision.
func ToComplex(value interface{}, prec uint) (cmplx *Complex) {
	_ = "STUB: not implemented"
	return nil
}

// IsInt returns true if both the real and imaginary parts are integers.
func (c Complex) IsInt() bool { _ = "STUB: not implemented"; return false }

func (c Complex) IsReal() bool { _ = "STUB: not implemented"; return false }

func (c *Complex) SetComplex128(x complex128) *Complex { _ = "STUB: not implemented"; return nil }

// Set sets an arbitrary precision complex number
func (c *Complex) Set(a *Complex) *Complex { _ = "STUB: not implemented"; return nil }

func (c *Complex) Prec() uint { _ = "STUB: not implemented"; return 0 }

func (c *Complex) SetPrec(prec uint) *Complex { _ = "STUB: not implemented"; return nil }

// Clone returns a new copy of the target arbitrary precision complex number
func (c *Complex) Clone() *Complex { _ = "STUB: not implemented"; return nil }

// Real returns the real part as a big.Float
func (c *Complex) Real() *big.Float {
	_ = "STUB: not implemented"

	// Imag returns the imaginary part as a big.Float
	return nil
}

func (c *Complex) Imag() *big.Float {
	_ = "STUB: not implemented"

	// Complex128 returns the arbitrary precision complex number as a complex128
	return nil
}

func (c *Complex) Complex128() complex128 { _ = "STUB: not implemented"; return 0 }

// Uint64 returns the real part of the complex number as an uint64.
func (c *Complex) Uint64() (u64 uint64) { _ = "STUB: not implemented"; return 0 }

// Int returns the real part of the complex number as a *big.Int.
func (c *Complex) Int() (bInt *big.Int) { _ = "STUB: not implemented"; return nil }

// Add adds two arbitrary precision complex numbers together
func (c *Complex) Add(a, b *Complex) *Complex { _ = "STUB: not implemented"; return nil }

// Sub subtracts two arbitrary precision complex numbers together
func (c *Complex) Sub(a, b *Complex) *Complex { _ = "STUB: not implemented"; return nil }

// Neg negates a and writes the result on c.
func (c *Complex) Neg(a *Complex) *Complex { _ = "STUB: not implemented"; return nil }

// ComplexMultiplier is a struct for the multiplication or division of two arbitrary precision complex numbers
type ComplexMultiplier struct {
	tmp0 *big.Float
	tmp1 *big.Float
	tmp2 *big.Float
	tmp3 *big.Float
}

// NewComplexMultiplier creates a new ComplexMultiplier
func NewComplexMultiplier() (cEval *ComplexMultiplier) { _ = "STUB: not implemented"; return nil }

// Mul evaluates c = a * b.
func (cEval *ComplexMultiplier) Mul(a, b, c *Complex) { _ = "STUB: not implemented"; return }

// Quo evaluates c = a / b.
func (cEval *ComplexMultiplier) Quo(a, b, c *Complex) { _ = "STUB: not implemented"; return }

// tmp0 = (a[0] * b[0]) + (a[1] * b[1]) real part
// tmp1 = (a[1] * b[0]) - (a[0] * b[0]) imag part
// tmp2 = (b[0] * b[0]) + (b[1] * b[1]) denominator
