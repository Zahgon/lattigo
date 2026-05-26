package ckks

import (
	"github.com/tuneinsight/lattigo/v6/utils/bignum"
)

const (
	minVecLenForLoopUnrolling = 16
)

// SpecialIFFTDouble performs the CKKS special inverse FFT transform in place.
func SpecialIFFTDouble(values []complex128, N, M int, rotGroup []int, roots []complex128) {
	_ = "STUB: not implemented"

	// Sanity check
	return
}

/* #nosec G115 -- previous check ensures N is greater than 0 */

/* #nosec G115 -- previous check ensures M is greater than 0 */

// SpecialFFTDouble performs the CKKS special FFT transform in place.
func SpecialFFTDouble(values []complex128, N, M int, rotGroup []int, roots []complex128) {
	_ = "STUB: not implemented"

	// Sanity check
	return
}

/* #nosec G115 -- previous check ensures N is greater than 0 */

/* #nosec G115 -- previous check ensures M is greater than 0 */

// SpecialFFTArbitrary evaluates the decoding matrix on a slice of ring.Complex values.
func SpecialFFTArbitrary(values []*bignum.Complex, N, M int, rotGroup []int, roots []*bignum.Complex) {
	_ = "STUB: not implemented"

	// Sanity check
	return
}

/* #nosec G115 -- previous check ensures N is greater than 0 */

/* #nosec G115 -- previous check ensures M is greater than 0 */

// SpecialIFFTArbitrary evaluates the encoding matrix on a slice of ring.Complex values.
func SpecialIFFTArbitrary(values []*bignum.Complex, N, M int, rotGroup []int, roots []*bignum.Complex) {
	_ = "STUB: not implemented"

	// Sanity check
	return
}

/* #nosec G115 -- previous check ensures N is greater than 0 */

/* #nosec G115 -- previous check ensures M is greater than 0 */

// SpecialFFTDoubleUL8 performs the CKKS special FFT transform in place with unrolled loops of size 8.
func SpecialFFTDoubleUL8(values []complex128, N, M int, rotGroup []int, roots []complex128) {
	_ = "STUB: not implemented"

	// Sanity check
	return
}

// Sanity check

/* #nosec G115 -- previous check ensures N is greater than 0 */

/* #nosec G115 -- previous check ensures M is greater than 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(values)%8 != 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(values)%8 != 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(values)%8 != 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(values)%8 != 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(values)%8 != 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(values)%8 != 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(values)%8 != 0 */

// SpecialiFFTDoubleUnrolled8 performs the CKKS special inverse FFT transform in place with unrolled loops of size 8.
func SpecialiFFTDoubleUnrolled8(values []complex128, N, M int, rotGroup []int, roots []complex128) {
	_ = "STUB: not implemented"

	// Sanity check
	return
}

// Sanity check

/* #nosec G115 -- previous check ensures N is greater than 0 */

/* #nosec G115 -- previous check ensures M is greater than 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(values)%8 != 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(values)%8 != 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(rotGroup)%8 != 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(rotGroup)%8 != 0  */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(rotGroup)%8 != 0  */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(rotGroup)%8 != 0  */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(rotGroup)%8 != 0  */

// divideComplex128SliceUnrolled8 divides the entries in values by scaleVal in place.
func divideComplex128SliceUnrolled8(values []complex128, scaleVal complex128) {
	_ = "STUB: not implemented"
	return
}

/* #nosec G103 -- behavior and consequences well understood */
