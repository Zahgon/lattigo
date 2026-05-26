package factorization

import (
	"math/big"
)

// Weierstrass is an elliptic curve y^2 = x^3 + ax + b mod N.
type Weierstrass struct {
	A, B, N *big.Int
}

// Point represents an elliptic curve point in standard coordinates.
type Point struct {
	X, Y *big.Int
}

// Add adds two Weierstrass points together with respect
// to the underlying Weierstrass curve.
// This method does not check if the points lie on
// the underlying curve.
func (w *Weierstrass) Add(P, Q Point) Point { _ = "STUB: not implemented"; return *new(Point) }

// slope

// S = (yQ-yP)/(xQ-xP)

// S = (3*(xP^2) + a)/(2*yP)

// s^2 - xP - xQ

// s*(xP-xR)-yP

// NewRandomWeierstrassCurve generates a new random Weierstrass curve modulo N,
// along with a random point that lies on the curve.
func NewRandomWeierstrassCurve(N *big.Int) (Weierstrass, Point) {
	_ = "STUB: not implemented"
	return *new(Weierstrass), *new(Point)
}

// Select random values for A, xG and yG

// Deduces B from Y^2 = X^3 + A * X + B evaluated at point (xG, yG)

// B = yG^2 - xG*(xG^2 - A)

// Checks that 4A^3 + 27B^2 != 0
