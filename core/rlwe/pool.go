package rlwe

import (
	"github.com/tuneinsight/lattigo/v6/ring/ringqp"
	"github.com/tuneinsight/lattigo/v6/utils/structs"
)

// BufferPool represents a pool of different objects (plaintexts, ciphertexts, polys) that can be used to instantiate temporary buffers.
type BufferPool struct {
	*ringqp.BufferPool
}

// NewPool returns a new pool given a RingQP, and optionally a pool to draw the backing arrays from.
func NewPool(rqp *ringqp.Ring, pools ...structs.BufferPool[*[]uint64]) *BufferPool {
	_ = "STUB: not implemented"
	// If no backing pool is given, we create one here.
	return nil
}

// AtLevel returns a new pool from which objects from polynomials at the given levels can be drawn.
// The method accepts up to two arguments:
// Zero level: the objects returned are built from polynomials at level 0.
// One level: the objects returned are built from polynomials in RingQ (resp. RingP) at the given level (resp. level 0).
// Two levels: the objects returned are built from polynomials in RingQ (resp. RingP) at levels[0] (resp. levels[1]).
func (pool BufferPool) AtLevel(levels ...int) *BufferPool { _ = "STUB: not implemented"; return nil }

// GetBuffCt returns a ciphertext that can be used as a buffer for intermediate computations.
// After use, the ciphertext should be recycled with [BufferPool.RecycleBuffCt].
// The optional dimensions specify the degree and level of the ciphertext (default to 2, pool.GetLevel()).
func (pool *BufferPool) GetBuffCt(dimensions ...int) *Ciphertext {
	_ = "STUB: not implemented"
	return nil
}

// sanity check: should not happen

// RecycleBuffCt recycles a temporary ciphertext (i.e. returns its backing uint64 arrays to the pool).
// The input ciphertext must not be used after calling this method.
func (pool *BufferPool) RecycleBuffCt(ct *Ciphertext) { _ = "STUB: not implemented"; return }

// GetBuffPt returns a plaintext that can be used as a buffer for intermediate computations.
// After use, the plaintext should be recycled with [BufferPool.RecycleBuffPt].
// The optional argument specifies the level of the returned plaintext (default to pool.GetLevel()).
func (pool *BufferPool) GetBuffPt(level ...int) *Plaintext { _ = "STUB: not implemented"; return nil }

// sanity check: should not happen

// RecycleBuffPt recycles a temporary plaintext (i.e. returns its backing uint64 arrays to the pool).
// The input plaintext must not be used after calling this method.
func (pool *BufferPool) RecycleBuffPt(pt *Plaintext) { _ = "STUB: not implemented"; return }

// GetBuffDecompQP returns buffers of polys to be used for RNS decomposition.
// After use, the array of buffers must be recycled with [BufferPool.RecycleBuffDecompQP].
func (pool *BufferPool) GetBuffDecompQP(params Parameters, levelQ, levelP int) []ringqp.Poly {
	_ = "STUB: not implemented"
	return nil
}

// RecycleBuffDecompQP recycles a temporary array of polys used for decomposition.
func (pool *BufferPool) RecycleBuffDecompQP(decomp []ringqp.Poly) {
	_ = "STUB: not implemented"
	return
}
