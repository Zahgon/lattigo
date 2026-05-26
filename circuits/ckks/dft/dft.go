// Package dft implements a homomorphic DFT circuit for the CKKS scheme.
package dft

import (
	"math/big"

	ltcommon "github.com/tuneinsight/lattigo/v6/circuits/ckks/lintrans"
	"github.com/tuneinsight/lattigo/v6/core/rlwe"
	"github.com/tuneinsight/lattigo/v6/schemes/ckks"
	"github.com/tuneinsight/lattigo/v6/utils/bignum"
)

// Type is a type used to distinguish between different discrete Fourier transformations.
type Type int

// HomomorphicEncode (IDFT) and HomomorphicDecode (DFT) are two available linear transformations for homomorphic encoding and decoding.
const (
	HomomorphicEncode = Type(0) // Homomorphic Encoding (IDFT)
	HomomorphicDecode = Type(1) // Homomorphic Decoding (DFT)
)

// Format is a type used to distinguish between the
// different input/output formats of the Homomorphic DFT.
type Format int

const (
	// Standard: designates the regular DFT.
	// Example: [a+bi, c+di] -> DFT([a+bi, c+di])
	Standard = Format(0)
	// SplitRealAndImag: HomomorphicEncode will return the real and
	// imaginary part into separate ciphertexts, both as real vectors.
	// Example: [a+bi, c+di] -> DFT([a, c]) and DFT([b, d])
	SplitRealAndImag = Format(1)
	// RepackImagAsReal: behaves the same as SplitRealAndImag except that
	// if the ciphertext is sparsely packed (at most N/4 slots), HomomorphicEncode
	// will repacks the real part into the left N/2 slots and the imaginary part
	// into the right N/2 slots. HomomorphicDecode must be specified with the same
	// format for correctness.
	// Example: [a+bi, 0, c+di, 0] -> [DFT([a, b]), DFT([b, d])]
	RepackImagAsReal = Format(2)
)

// Matrix is a struct storing the factorized IDFT, DFT matrices, which are
// used to homomorphically encode and decode a ciphertext respectively.
type Matrix struct {
	MatrixLiteral
	Matrices []ltcommon.LinearTransformation
}

// MatrixLiteral is a struct storing the parameters to generate the factorized DFT/IDFT matrices.
// This struct has mandatory and optional fields.
//
// Mandatory:
//   - Type: HomomorphicEncode (a.k.a. CoeffsToSlots) or HomomorphicDecode (a.k.a. SlotsToCoeffs)
//   - LogSlots: log2(slots)
//   - LevelQ: starting level of the linear transformation
//   - LevelP: number of auxiliary primes used during the automorphisms. User must ensure that this
//     value is the same as the one used to generate the Galois keys.
//   - Levels: depth of the linear transform (i.e. the degree of factorization of the encoding matrix)
//
// Optional:
//   - Format: which post-processing (if any) to apply to the DFT.
//   - Scaling: constant by which the matrix is multiplied
//   - BitReversed: if true, then applies the transformation bit-reversed and expects bit-reversed inputs
//   - LogBSGSRatio: log2 of the ratio between the inner and outer loop of the baby-step giant-step algorithm
type MatrixLiteral struct {
	// Mandatory
	Type     Type
	LogSlots int
	LevelQ   int
	LevelP   int
	Levels   []int
	// Optional
	Format       Format     // Default: standard.
	Scaling      *big.Float // Default 1.0.
	BitReversed  bool       // Default: False.
	LogBSGSRatio int        // Default: 0.
}

// Depth returns the number of levels allocated to the linear transform.
// If actual == true then returns the number of moduli consumed, else
// returns the factorization depth.
func (d MatrixLiteral) Depth(actual bool) (depth int) { _ = "STUB: not implemented"; return 0 }

// GaloisElements returns the list of rotations performed during the CoeffsToSlot operation.
func (d MatrixLiteral) GaloisElements(params ckks.Parameters) (galEls []uint64) {
	_ = "STUB: not implemented"
	return nil
}

// Coeffs to Slots rotations

// MarshalBinary returns a JSON representation of the the target [MatrixLiteral] on a slice of bytes.
// See `Marshal` from the `encoding/json` package.
func (d MatrixLiteral) MarshalBinary() (data []byte, err error) {
	_ = "STUB: not implemented"
	return nil,

		// UnmarshalBinary reads a JSON representation on the target [MatrixLiteral] struct.
		// See `Unmarshal` from the `encoding/json` package.
		nil
}

func (d *MatrixLiteral) UnmarshalBinary(data []byte) error { _ = "STUB: not implemented"; return nil }

// Evaluator is an evaluator providing an API for homomorphic DFT.
// All fields of this struct are public, enabling custom instantiations.
type Evaluator struct {
	*ckks.Evaluator
	LTEvaluator *ltcommon.Evaluator
	parameters  ckks.Parameters
	pool        *rlwe.BufferPool
}

// NewEvaluator instantiates a new [Evaluator] from a [ckks.Evaluator].
func NewEvaluator(params ckks.Parameters, eval *ckks.Evaluator) *Evaluator {
	_ = "STUB: not implemented"
	return nil
}

// NewMatrixFromLiteral generates the factorized DFT/IDFT matrices for the homomorphic encoding/decoding.
func NewMatrixFromLiteral(params ckks.Parameters, d MatrixLiteral, encoder *ckks.Encoder) (Matrix, error) {
	_ = "STUB: not implemented"
	return *new(Matrix), nil
}

// CoeffsToSlots vectors

// CoeffsToSlotsNew applies the homomorphic encoding and returns the result on new ciphertexts.
// Homomorphically encodes a complex vector vReal + i*vImag.
// Given n = current number of slots and N/2 max number of slots (half the ring degree):
// If the packing is sparse (n < N/2), then returns ctReal = Ecd(vReal || vImag) and ctImag = nil.
// If the packing is dense (n == N/2), then returns ctReal = Ecd(vReal) and ctImag = Ecd(vImag).
func (eval *Evaluator) CoeffsToSlotsNew(ctIn *rlwe.Ciphertext, ctsMatrices Matrix) (ctReal, ctImag *rlwe.Ciphertext, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// CoeffsToSlots applies the homomorphic encoding and returns the results on the provided ciphertexts.
// Homomorphically encodes a complex vector vReal + i*vImag of size n on a real vector of size 2n.
// If the packing is sparse (n < N/2), then returns ctReal = Ecd(vReal || vImag) and ctImag = nil.
// If the packing is dense (n == N/2), then returns ctReal = Ecd(vReal) and ctImag = Ecd(vImag).
func (eval *Evaluator) CoeffsToSlots(ctIn *rlwe.Ciphertext, ctsMatrices Matrix, ctReal, ctImag *rlwe.Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Imag part

// Real part

// If repacking, then ct0 and ct1 right n/2 slots are zero.

// SlotsToCoeffsNew applies the homomorphic decoding and returns the result on a new ciphertext.
// Homomorphically decodes a real vector of size 2n on a complex vector vReal + i*vImag of size n.
// If the packing is sparse (n < N/2) then ctReal = Ecd(vReal || vImag) and ctImag = nil.
// If the packing is dense (n == N/2), then ctReal = Ecd(vReal) and ctImag = Ecd(vImag).
func (eval *Evaluator) SlotsToCoeffsNew(ctReal, ctImag *rlwe.Ciphertext, stcMatrices Matrix) (opOut *rlwe.Ciphertext, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SlotsToCoeffs applies the homomorphic decoding and returns the result on the provided ciphertext.
// Homomorphically decodes a real vector of size 2n on a complex vector vReal + i*vImag of size n.
// If the packing is sparse (n < N/2) then ctReal = Ecd(vReal || vImag) and ctImag = nil.
// If the packing is dense (n == N/2), then ctReal = Ecd(vReal) and ctImag = Ecd(vImag).
func (eval *Evaluator) SlotsToCoeffs(ctReal, ctImag *rlwe.Ciphertext, stcMatrices Matrix, opOut *rlwe.Ciphertext) (err error) {
	_ = "STUB: not implemented"
	// If full packing, the repacking can be done directly using ct0 and ct1.
	return nil
}

// dft evaluates homorphically the iDFT/DFT [Matrix] on ctIn and stores the result in opOut.
func (eval *Evaluator) dft(ctIn *rlwe.Ciphertext, mat Matrix, opOut *rlwe.Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Encoding matrices are a special case of `fractal` linear transform
// that doesn't change the underlying plaintext polynomial Y = X^{N/n}
// of the input ciphertext.

func fftPlainVec(logN, dslots int, roots []*bignum.Complex, pow5 []int) (a, b, c [][]*bignum.Complex) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func ifftPlainVec(logN, dslots int, roots []*bignum.Complex, pow5 []int) (a, b, c [][]*bignum.Complex) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func addMatrixRotToList(pVec map[int]bool, rotations []int, N1, slots int, repack bool) []int {
	_ = "STUB: not implemented"
	return nil
}

// Sparse repacking, occurring during the first DFT matrix of the CoeffsToSlots.

// Other cases

func (d MatrixLiteral) computeBootstrappingDFTIndexMap(logN int) (rotationMap []map[int]bool) {
	_ = "STUB: not implemented"
	return nil
}

// We compute the chain of merge in order or reverse order depending if its DFT or InvDFT because
// the way the levels are collapsed has an impact on the total number of rotations and keys to be
// stored. Ex. instead of using 255 + 64 plaintext vectors, we can use 127 + 128 plaintext vectors
// by reversing the order of the merging.

// Special initial matrix for the repacking before Decode

// Merges this special initial matrix with the first layer of Decode DFT

// Continues the merging with the next layers if the total depth requires it.

// First layer of the i-th level of the DFT

// Merges the layer with the next levels of the DFT if the total depth requires it.

func genWfftIndexMap(logL, level int, ltType Type, bitreversed bool) (vectors map[int]bool) {
	_ = "STUB: not implemented"
	return nil
}

func genWfftRepackIndexMap(logL, level int) (vectors map[int]bool) {
	_ = "STUB: not implemented"
	return nil
}

func nextLevelfftIndexMap(vec map[int]bool, logL, N, nextLevel int, ltType Type, bitreversed bool) (newVec map[int]bool) {
	_ = "STUB: not implemented"
	return nil
}

// GenMatrices returns the ordered list of factors of the non-zero diagonals of the IDFT (encoding) or DFT (decoding) matrix.
func (d MatrixLiteral) GenMatrices(LogN int, prec uint) (plainVector []ltcommon.Diagonals[*bignum.Complex]) {
	_ = "STUB: not implemented"
	return nil
}

// We compute the chain of merge in order or reverse order depending if its DFT or InvDFT because
// the way the levels are collapsed has an impact on the total number of rotations and keys to be
// stored. Ex. instead of using 255 + 64 plaintext vectors, we can use 127 + 128 plaintext vectors
// by reversing the order of the merging.

// Special initial matrix for the repacking before DFT

// Merges this special initial matrix with the first layer of DFT

// Continues the merging with the next layers if the total depth requires it.

// First layer of the i-th level of the DFT

// Merges the layer with the next levels of the DFT if the total depth requires it.

// Repacking after the IDFT (we multiply the last matrix with the vector [1, 1, ..., 1, 1, 0, 0, ..., 0, 0]).

// If DFT matrix, rescale by 1/N

// Real/Imag extraction 1/2 factor

// Spreads the scale across the matrices

func genFFTDiagMatrix(logL, fftLevel int, a, b, c []*bignum.Complex, ltType Type, bitreversed bool) (vectors map[int][]*bignum.Complex) {
	_ = "STUB: not implemented"
	return nil
}

func genRepackMatrix(logL int, prec uint, bitreversed bool) (vectors map[int][]*bignum.Complex) {
	_ = "STUB: not implemented"
	return nil
}

func multiplyFFTMatrixWithNextFFTLevel(vec map[int][]*bignum.Complex, logL, N, nextLevel int, a, b, c []*bignum.Complex, ltType Type, bitreversed bool) (newVec map[int][]*bignum.Complex) {
	_ = "STUB: not implemented"
	return nil
}

func addToDiagMatrix(diagMat map[int][]*bignum.Complex, index int, vec []*bignum.Complex) {
	_ = "STUB: not implemented"
	return
}

func rotateAndMulNew(a []*bignum.Complex, k int, b []*bignum.Complex) (c []*bignum.Complex) {
	_ = "STUB: not implemented"
	return nil
}

func add(a, b, c []*bignum.Complex) { _ = "STUB: not implemented"; return }
