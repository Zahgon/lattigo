package bignum

import (
	"math/big"
)

// Remez implements the optimized multi-interval minimax approximation
// algorithm of Lee et al. (https://eprint.iacr.org/2020/552).
// This is an iterative algorithm that returns the minimax polynomial
// approximation of any function that is smooth over a set of interval
// [a0, b0] U [a1, b1] U ... U [ai, bi].
type Remez struct {
	RemezParameters
	Degree int

	extremePoints      []point
	localExtremePoints []point
	nbExtremePoints    int

	MaxErr, MinErr *big.Float

	Nodes  []point
	Matrix [][]*big.Float
	Vector []*big.Float
	Coeffs []*big.Float
}

type point struct {
	x, y      *big.Float
	slopesign int
}

// RemezParameters is a struct storing the parameters
// required to initialize the Remez algorithm.
type RemezParameters struct {
	// Function is the function to approximate.
	// It has to be smooth in the defined intervals.
	Function func(x *big.Float) (y *big.Float)

	// Basis is the basis to use.
	// Supported basis are: Monomial and Chebyshev
	Basis Basis

	// Intervals is the set of interval [ai, bi] on which to approximate
	// the function. Each interval also define the number of nodes (points)
	// that will be used to approximate the function inside this interval.
	// This allows the user to implement a separate algorithm that allocates
	// an optimal number of nodes per interval.
	Intervals []Interval

	// ScanStep is the size of the default step used to find the extreme points.
	// The smaller this value is, the lower the probability to miss an extreme point is
	// but the longer each iteration will be.
	// A good starting value is 2^{-10}.
	ScanStep *big.Float

	// Prec defines the bit precision of the overall computation.
	Prec uint

	// OptimalScanStep is a boolean to use a dynamic update of the scan step during each
	// iteration.
	OptimalScanStep bool
}

// NewRemez instantiates a new Remez algorithm from the provided parameters.
func NewRemez(p RemezParameters) (r *Remez) { _ = "STUB: not implemented"; return nil }

// Approximate starts the approximation process.
// maxIter: the maximum number of iterations before the approximation process is terminated.
// threshold: the minimum value that (maxErr-minErr)/minErr (the normalized absolute difference
// between the maximum and minimum approximation error over the defined intervals) must take
// before the approximation process is terminated.
func (r *Remez) Approximate(maxIter int, threshold float64) { _ = "STUB: not implemented"; return }

// Solves the linear system and gets the new set of coefficients

// Finds the extreme points of p(x) - f(x) (where the absolute error is max)

// Choose the new nodes based on the set of extreme points

// ShowCoeffs prints the coefficient of the approximate
// prec: the bit precision of the printed values.
func (r *Remez) ShowCoeffs(prec int) { _ = "STUB: not implemented"; return }

// ShowError prints the minimum and maximum error of the approximate
// prec: the bit precision of the printed values.
func (r *Remez) ShowError(prec int) { _ = "STUB: not implemented"; return }

func (r *Remez) initialize() { _ = "STUB: not implemented"; return }

/* #nosec G601 -- Implicit memory aliasing in for loop acknowledged */

/* #nosec G601 -- Implicit memory aliasing in for loop acknowledged */

func (r *Remez) getCoefficients() {
	_ = "STUB: not implemented"

	// Constructs the linear system
	// | 1 x0 x0^2 x0^3 ...  1 | f(x0)
	// | 1 x1 x1^2 x1^3 ... -1 | f(x1)
	// | 1 x2 x2^2 x2^3 ...  1 | f(x2)
	// | 1 x3 x3^2 x3^3 ... -1 | f(x3)
	// |          .            |   .
	// |          .            |   .
	// |          .            |   .
	return
}

/*
	for i := 0; i < r.Degree+2; i++{
		for j := 0; j < r.Degree+2; j++{
			fmt.Printf("%v\n", r.Matrix[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
*/

// Solves the linear system

// Updates the new [x0, x1, ..., xi]

func (r *Remez) findExtremePoints() { _ = "STUB: not implemented"; return }

// e = p(x) - f(x) over [a, b]

// show error message

// ChooseNewNodes implements Algorithm 3 of High-Precision Bootstrapping
// of RNS-CKKS Homomorphic Encryption Using Optimal Minimax Polynomial
// Approximation and Inverse Sine Function (https://eprint.iacr.org/2020/552).
// This is an optimized Go reimplementation of Remez::choosemaxs at
// https://github.com/snu-ccl/FHE-MP-CNN/blob/main-3.6.6/cnn_ckks/common/Remez.cpp
func (r *Remez) chooseNewNodes() {
	_ = "STUB: not implemented"

	// Allocates the list of new nodes
	return
}

// Retrieve the list of extrem points

// Resets max and min error

//=========================
//========= PART 1 ========
//=========================

// Line 1 to 8 of Algorithm 3

// The first part of the algorithm is to remove
// consecutive extreme points with the same slope sign,
// which will ensure that new linear system has a
// solution by the Haar condition.

// Stores consecutive extreme points with the same slope sign
// It is unlikely that more that two consecutive extreme points
// will have the same slope sign.

// To find the maximum value between extreme points that have the
// same slope sign.

// Tracks the total number of extreme points iterated on

// If idxAdjSameSlope is empty then adds the next point

// If the slope of two consecutive extreme points is not alternating in sign
// then adds the point index to the temporary array

// If the next point has alternating sign, then iterates over all the index in the temporary array
// with extreme points whose slope is of the same sign and looks for the one with the maximum
// absolute value

// Adds to the new nodes the extreme points whose absolute value is the largest
// between all consecutive extreme points with the same slope sign

// The above loop might terminate without flushing the array of extreme points
// with the same slope sign, the second part of the loop is called one last time.

//=========================
//========= PART 2 ========
//=========================

// Lines 11 to 24 of Algorithm 3

// Choosing the new nodes if the set of alternating extreme points
// is larger than degree+2.

// Loops run as long as the number of extreme points is not equal to deg+2 (the dimension of the linear system)

// If the number of remaining extreme points is one more than the number needed
// then we can remove only one point

// Removes the largest one between the first and the last

// If the number of remaining extreme points is two more than the number needed
// then we can remove two points.

// Finds the minimum index of the sum of two adjacent points

// If the index is the last, then remove the first and last points

// Else remove the two consecutive points

// If the number of remaining extreme points is more four over the number needed
// then remove up to two points, prioritizing the first and last points.

// Finds the minimum index of the sum of two adjacent points

// If the first element is included in the smallest sum, then removes it

// If the last element is included in the smallest sum, then removes it

// Else removes the two consecutive points adding to the smallest sum

// Assigns the new points to the nodes and computes the min and max error

// Deep copy

// we must evaluate, because Y was the error Function)
// should have alternating sign

// findLocalExtrempointsWithSlope finds local extrema/minima of a function.
// It starts by scanning the interval with a pre-defined window size, until it finds that the function is concave or convex
// in this window. Then it uses a binary search to find the local maximum/minimum in this window. The process is repeated
// until the entire interval has been scanned.
// This is an optimized Go re-implementation of the method find_extreme that can be found at
// https://github.com/snu-ccl/FHE-MP-CNN/blob/main-3.6.6/cnn_ckks/common/MinicompFunc.cpp
func (r *Remez) findLocalExtrempointsWithSlope(fErr func(*big.Float) (y *big.Float), interval Interval) []point {
	_ = "STUB: not implemented"
	return nil
}

// start + 10*scan/pow(10,i)

// end - 10*scan/pow(10,i)

// a < scanRight && scanRight < b

// Breaks when the scan window gets out of the interval

// Positive and negative slope (concave)

// Negative and positive slope (convex)

// findLocalMaximum finds the local maximum of a function that is concave in a given window.
func findLocalMaximum(fErr func(x *big.Float) (y *big.Float), start, end *big.Float, prec uint, p *point) {
	_ = "STUB: not implemented"
	return
}

/* #nosec G115 -- prec is a bit-size */

// Obtains the sign of the err Function in the interval (normalized and zeroed)
// 0: [0.00, 0.25]
// 1: [0.25, 0.50]
// 2: [0.50, 0.75]
// 3: [0.75, 1.00]

// Look for a sign change between the 4 intervals.
// Since we are here in a concave Function, we look
// for the point in the interval where the sign of the
// err Function changes.

// Sign change occurs between [0, 0.5]

// Reduces the windowEnd from 1 to 0.5

// Divides the scan step by half

// Sign change occurs between [0.25, 0.75]

// Increases windowStart from 0 to 0.25

// Decreases windowEnd from 1 to 0.75

// Divides the scan step by half

// Sign change occurs between [0.5, 1.0]

// Increases windowStart fro 0 to 0.5

// Divides the scan step by half

// findLocalMaximum finds the local maximum of a function that is convex in a given window.
func findLocalMinimum(fErr func(x *big.Float) (y *big.Float), start, end *big.Float, prec uint, p *point) {
	_ = "STUB: not implemented"
	return
}

/* #nosec G115 -- prec is a bit-size */

// Obtains the sign of the err Function in the interval (normalized and zeroed)
// 0: [0.00, 0.25]
// 1: [0.25, 0.50]
// 2: [0.50, 0.75]
// 3: [0.75, 1.00]

// Look for a sign change between the 4 intervals.
// Since we are here in a convex Function, we look
// for the point in the interval where the sign of the
// err Function changes.

// Sign change occurs between [0, 0.5]

// Reduces the windowEnd from 1 to 0.5

// Divides the scan step by half

// Sign change occurs between [0.25, 0.75]

// Increases windowStart from 0 to 0.25

// Decreases windowEnd from 1 to 0.75

// Divides the scan step by half

// Sign change occurs between [0.5, 1.0]

// Increases windowStart fro 0 to 0.5

// Divides the scan step by half

// slopes takes a window, divides it into four intervals and computes the sign of the slope of the error function in each sub-interval.
func slopes(fErr func(x *big.Float) (y *big.Float), searchStart, searchEnd, searchquarter *big.Float) (searchslopeLeft, searchslopeRight, searchInc3, searchInc4 int) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0
}

// [start, start + sc]

//[start + sc, start + 2*sc]

// [start + 2*sc, enc-sc]

// [end-sc, end]

func (r *Remez) eval(x *big.Float) (y *big.Float) { _ = "STUB: not implemented"; return nil }

// solves for y the system matrix * y = vector using Gaussian elimination.
func solveLinearSystemInPlace(matrix [][]*big.Float, vector []*big.Float) {
	_ = "STUB: not implemented"
	return
}
