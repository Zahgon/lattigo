package ringqp

import (
	"github.com/tuneinsight/lattigo/v6/ring"
	"github.com/tuneinsight/lattigo/v6/utils/structs"
)

// BufferPool represents a pool of polys that can be used (concurrently) to instantiate temporary polynomials in RingQP.
type BufferPool struct {
	*ring.BufferPool
	PoolP *ring.BufferPool
}

// NewPool returns a new pool given a RingQP, and optionally a pool to draw the backing arrays from.
func NewPool(ringqp *Ring, pools ...structs.BufferPool[*[]uint64]) *BufferPool {
	_ = "STUB: not implemented"
	// If no backing pool is given, we create one here.
	return nil
}

// AtLevel returns a new pool from which polynomials at the given levels can be drawn.
// The method accepts up to two arguments:
// Zero level: the polyomials are returned at level 0.
// One level: the polynomials in RingQ (resp. RingP) are returned at the given level (resp. level 0).
// Two levels: the polynomials in RingQ (resp. RingP) are returned at levels[0] (resp. levels[1]).
func (p BufferPool) AtLevel(levels ...int) *BufferPool { _ = "STUB: not implemented"; return nil }

// GetBuffPolyQP returns a new [Poly], built from backing []uint64 arrays obtained from a pool.
// After use, the [Poly] should be recycled using the [BufferPool.RecycleBuffPolyQP] method.
func (p BufferPool) GetBuffPolyQP() *Poly { _ = "STUB: not implemented"; return nil }

// RecycleBuffPolyQP takes a reference to a [Poly] and recycles its backing []uint64 arrays
// (i.e. they are returned to a pool). The input [Poly] must not be used after calling this method.
func (p BufferPool) RecycleBuffPolyQP(poly *Poly) { _ = "STUB: not implemented"; return }
