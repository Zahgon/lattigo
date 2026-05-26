package rlwe

// ApplyEvaluationKey is a generic method to apply an [EvaluationKey] on a ciphertext.
// An EvaluationKey is a type of public key that is be used during the evaluation of
// a homomorphic circuit to provide additional functionalities, like relinearization
// or rotations.
//
// In a nutshell, an [EvaluationKey] encrypts a secret skIn under a secret skOut and
// enables the public and non interactive re-encryption of any ciphertext encrypted
// under skIn to a new ciphertext encrypted under skOut.
//
// The method will return an error if either ctIn or opOut degree isn't 1.
//
// This method can also be used to switch a ciphertext to one with a different ring degree.
// Note that the parameters of the smaller ring degree must be the same or a subset of the
// moduli Q and P of the one for the larger ring degree.
//
// To do so, it must be provided with the appropriate [EvaluationKey], and have the operands
// matching the target ring degrees.
//
// To switch a ciphertext to a smaller ring degree:
//   - ctIn ring degree must match the evaluator's ring degree.
//   - opOut ring degree must match the smaller ring degree.
//   - evk must have been generated using the key-generator of the large ring degree with as input large-key -> small-key.
//
// To switch a ciphertext to a smaller ring degree:
//   - ctIn ring degree must match the smaller ring degree.
//   - opOut ring degree must match the evaluator's ring degree.
//   - evk must have been generated using the key-generator of the large ring degree with as input small-key -> large-key.
func (eval Evaluator) ApplyEvaluationKey(ctIn *Ciphertext, evk *EvaluationKey, opOut *Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Re-encryption to a larger ring degree.

// Maps to larger ring degree Y = X^{N/n} -> X

// Re-encrypt opOut from the key from small to larger ring degree

// Re-encryption to a smaller ring degree.

// Switches key from large to small degree

// Maps to smaller ring degree X -> Y = X^{N/n}

// Re-encryption to the same ring degree.

func (eval Evaluator) applyEvaluationKey(level int, ctIn *Ciphertext, evk *EvaluationKey, opOut *Ciphertext) {
	_ = "STUB: not implemented"
	return
}

// Relinearize applies the relinearization procedure on ct0 and returns the result in opOut.
// Relinearization is a special procedure required to ensure ciphertext compactness.
// It takes as input a quadratic ciphertext, that decrypts with the key (1, sk, sk^2) and
// outputs a linear ciphertext that decrypts with the key (1, sk).
// In a nutshell, the relinearization re-encrypt the term that decrypts using sk^2 to one
// that decrypts using sk.
// The method will return an error if:
//   - The input ciphertext degree isn't 2.
//   - The corresponding relinearization key to the ciphertext degree
//
// is missing.
func (eval Evaluator) Relinearize(ctIn *Ciphertext, opOut *Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}
