package rlwe

import (
	"github.com/tuneinsight/lattigo/v6/ring"
	"github.com/tuneinsight/lattigo/v6/ring/ringqp"
)

// GadgetProduct evaluates poly x Gadget -> RLWE where
//
// ct = [<decomp(cx), gadget[0]>, <decomp(cx), gadget[1]>] mod Q
//
// Expects the flag IsNTT of ct to correctly reflect the domain of cx.
func (eval Evaluator) GadgetProduct(levelQ int, cx ring.Poly, gadgetCt *GadgetCiphertext, ct *Ciphertext) {
	_ = "STUB: not implemented"
	return
}

// ModDown takes ctQP (mod QP) and returns ct = (ctQP/P) (mod Q).
func (eval Evaluator) ModDown(levelQ, levelP int, ctQP *Element[ringqp.Poly], ct *Ciphertext) {
	_ = "STUB: not implemented"
	return
}

// NTT -> NTT

// NTT -> INTT

// INTT -> NTT

// INTT -> INTT

// NTT -> NTT

// NTT -> INTT

// INTT -> NTT

// INTT -> INTT

// GadgetProductLazy evaluates poly x Gadget -> RLWE where
//
// ctQP = [<decomp(cx), gadget[0]>, <decomp(cx), gadget[1]>] mod QP
//
// Expects the flag IsNTT of ctQP to correctly reflect the domain of cx.
//
// Result NTT domain is returned according to the NTT flag of ctQP.
//
// The method will return an error if ctQP.Level() < gadgetCt.Level().
func (eval Evaluator) GadgetProductLazy(levelQ int, cx ring.Poly, gadgetCt *GadgetCiphertext, ctQP *Element[ringqp.Poly]) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (eval Evaluator) gadgetProductMultiplePLazy(levelQ int, cx ring.Poly, gadgetCt *GadgetCiphertext, ctQP *Element[ringqp.Poly]) {
	_ = "STUB: not implemented"
	return
}

// Re-encryption with CRT decomposition for the Qi

func (eval Evaluator) gadgetProductSinglePAndBitDecompLazy(levelQ int, cx ring.Poly, gadgetCt *GadgetCiphertext, ctQP *Element[ringqp.Poly]) {
	_ = "STUB: not implemented"
	return
}

// Re-encryption with CRT decomposition for the Qi

// Only centers the coefficients if the mask is 0
// As centering doesn't help reduce the noise if
// the power of two decomposition is applied on top
// of the RNS decomposition

// GadgetProductHoisted applies the key-switch to the decomposed polynomial c2 mod QP (BuffQPDecompQP)
// and divides the result by P, reducing the basis from QP to Q.
//
// ct = [<BuffQPDecompQP,  gadgetCt[0]) mod Q
//
// BuffQPDecompQP is expected to be in the NTT domain.
//
// Result NTT domain is returned according to the NTT flag of ct.
func (eval Evaluator) GadgetProductHoisted(levelQ int, BuffQPDecompQP []ringqp.Poly, gadgetCt *GadgetCiphertext, ct *Ciphertext) {
	_ = "STUB: not implemented"
	return
}

// GadgetProductHoistedLazy applies the gadget product to the decomposed polynomial c2 mod QP (BuffQPDecompQ and BuffQPDecompP)
//
// BuffQP2 = dot(BuffQPDecompQ||BuffQPDecompP * gadgetCt[0]) mod QP
// BuffQP3 = dot(BuffQPDecompQ||BuffQPDecompP * gadgetCt[1]) mod QP
//
// BuffQPDecompQP is expected to be in the NTT domain.
//
// Result NTT domain is returned according to the NTT flag of ctQP.
//
// The method will return an error if ctQP.Level() < gadgetCt.Level().
func (eval Evaluator) GadgetProductHoistedLazy(levelQ int, BuffQPDecompQP []ringqp.Poly, gadgetCt *GadgetCiphertext, ctQP *Element[ringqp.Poly]) (err error) {
	_ = "STUB: not implemented"

	// Sanity check for invalid parameters.
	return nil
}

func (eval Evaluator) gadgetProductMultiplePLazyHoisted(levelQ int, BuffQPDecompQP []ringqp.Poly, gadgetCt *GadgetCiphertext, ct *Element[ringqp.Poly]) {
	_ = "STUB: not implemented"
	return
}

// Key switching with CRT decomposition for the Qi

// DecomposeNTT applies the full RNS basis decomposition on c2.
// Expects the IsNTT flag of c2 to correctly reflect the domain of c2.
// BuffQPDecompQ and BuffQPDecompQ are vectors of polynomials (mod Q and mod P) that store the
// special RNS decomposition of c2 (in the NTT domain)
func (eval Evaluator) DecomposeNTT(levelQ, levelP, nbPi int, c2 ring.Poly, c2IsNTT bool, decompQP []ringqp.Poly) {
	_ = "STUB: not implemented"
	return
}

// DecomposeSingleNTT takes the input polynomial c2 (c2NTT and c2InvNTT, respectively in the NTT and out of the NTT domain)
// modulo the RNS basis, and returns the result on c2QiQ and c2QiP, the receiver polynomials respectively mod Q and mod P (in the NTT domain)
func (eval Evaluator) DecomposeSingleNTT(levelQ, levelP, nbPi, BaseRNSDecompositionVectorSize int, c2NTT, c2InvNTT, c2QiQ, c2QiP ring.Poly) {
	_ = "STUB: not implemented"
	return
}

// c2_qi = cx mod qi mod qi

// c2QiP = c2 mod qi mod pj

/*
type DecompositionBuffer [][]ringqp.Poly

func (eval Evaluator) ALlocateDecompositionBuffer(levelQ, levelP, Pow2Base int) (DecompositionBuffer){

	decompQP := make([][]ringqp.Poly, BaseRNSDecompositionVectorSize)
	for i := 0; i < BaseRNSDecompositionVectorSize; i++ {

		for j := 0; j < BaseTwoDecompositionVectorSize; j++{
			DecompositionBuffer[i][j] = ringQP.NewPoly()
		}
	}

	return decompQPs
}
*/
