package rlwe

import (
	"github.com/tuneinsight/lattigo/v6/ring/ringqp"
)

// Automorphism computes phi(ct), where phi is the map X -> X^galEl. The method requires
// that the corresponding RotationKey has been added to the [Evaluator]. The method will
// return an error if either ctIn or opOut degree is not equal to 1.
func (eval Evaluator) Automorphism(ctIn *Ciphertext, galEl uint64, opOut *Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// AutomorphismHoisted is similar to [Evaluator.Automorphism], except that it takes as input ctIn and c1DecompQP, where c1DecompQP is the RNS
// decomposition of its element of degree 1. This decomposition can be obtained with [Evaluator.DecomposeNTT].
// The method requires that the corresponding RotationKey has been added to the [Evaluator].
// The method will return an error if either ctIn or opOut degree is not equal to 1.
func (eval Evaluator) AutomorphismHoisted(level int, ctIn *Ciphertext, c1DecompQP []ringqp.Poly, galEl uint64, opOut *Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// AutomorphismHoistedLazy is similar to [Evaluator.AutomorphismHoisted], except that it returns a ciphertext modulo QP and scaled by P.
// The method requires that the corresponding RotationKey has been added to the [Evaluator].
// Result NTT domain is returned according to the NTT flag of ctQP.
func (eval Evaluator) AutomorphismHoistedLazy(levelQ int, ctIn *Ciphertext, c1DecompQP []ringqp.Poly, galEl uint64, ctQP *Element[ringqp.Poly]) (err error) {
	_ = "STUB: not implemented"
	return nil
}
