package structs

import "sync"

// BufferPool is an interface for all pools of buffers.
type BufferPool[T any] interface {
	Get() T
	Put(T)
}

// SyncPool is a wrapper around [sync.Pool] (it avoids doing type conversion after Get()).
type SyncPool[T any] struct {
	pool *sync.Pool
}

// NewSyncPool creates a new SyncPool.
// The input function f is the function that is used to create new objects if none is available in the pool.
func NewSyncPool[T any](f func() T) *SyncPool[T] { _ = "STUB: not implemented"; return nil }

func NewSyncPoolUint64(N int) *SyncPool[*[]uint64] { _ = "STUB: not implemented"; return nil }

// Get returns a new object of type T from the pool.
func (spool *SyncPool[T]) Get() T { _ = "STUB: not implemented"; return *new(T) }

// Put returns the buff to the pool.
func (spool *SyncPool[T]) Put(buff T) { _ = "STUB: not implemented"; return }
