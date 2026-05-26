package blindrot

import (
	"github.com/tuneinsight/lattigo/v6/core/rgsw"
	"github.com/tuneinsight/lattigo/v6/core/rlwe"
)

const (
	// Parameter w of Algorithm 3 in https://eprint.iacr.org/2022/198
	windowSize = 10
)

// BlindRotationEvaluationKeySet is a interface implementing methods
// to load the blind rotation keys (RGSW) and automorphism keys
// (via the [rlwe.EvaluationKeySet] interface).
// Implementation of this interface must be safe for concurrent use.
type BlindRotationEvaluationKeySet interface {

	// GetBlindRotationKey should return RGSW(X^{s[i]})
	GetBlindRotationKey(i int) (brk *rgsw.Ciphertext, err error)

	// GetEvaluationKeySet should return an rlwe.EvaluationKeySet
	// providing access to all the required automorphism keys.
	GetEvaluationKeySet() (evk rlwe.EvaluationKeySet, err error)
}

// MemBlindRotationEvaluationKeySet is a basic in-memory implementation of the [BlindRotationEvaluationKeySet] interface.
type MemBlindRotationEvaluationKeySet struct {
	BlindRotationKeys []*rgsw.Ciphertext
	AutomorphismKeys  []*rlwe.GaloisKey
}

func (evk MemBlindRotationEvaluationKeySet) GetBlindRotationKey(i int) (*rgsw.Ciphertext, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (evk MemBlindRotationEvaluationKeySet) GetEvaluationKeySet() (rlwe.EvaluationKeySet, error) {
	_ = "STUB: not implemented"
	return *new(rlwe.EvaluationKeySet), nil
}

// GenEvaluationKeyNew generates a new Blind Rotation evaluation key
func GenEvaluationKeyNew(paramsRLWE rlwe.ParameterProvider, skRLWE *rlwe.SecretKey, paramsLWE rlwe.ParameterProvider, skLWE *rlwe.SecretKey, evkParams ...rlwe.EvaluationKeyParameters) (key MemBlindRotationEvaluationKeySet) {
	_ = "STUB: not implemented"
	return *new(MemBlindRotationEvaluationKeySet)
}

// Sanity check, this error should never happen unless this algorithm
// has been improperly modified to provides invalid inputs.
