package utils

import (
	"golang.org/x/exp/constraints"
)

// Alias1D returns true if x and y share the same base array.
// Taken from http://golang.org/src/pkg/math/big/nat.go#L340 .
func Alias1D[V any](x, y []V) bool { _ = "STUB: not implemented"; return false }

// Alias2D returns true if x and y share the same base array.
// Taken from http://golang.org/src/pkg/math/big/nat.go#L340 .
func Alias2D[V any](x, y [][]V) bool { _ = "STUB: not implemented"; return false }

// GetKeys returns the keys of the input map.
// Order is not guaranteed.
func GetKeys[K constraints.Ordered, V any](m map[K]V) (keys []K) {
	_ = "STUB: not implemented"
	return nil
}

// GetSortedKeys returns the sorted keys of a map.
func GetSortedKeys[K constraints.Ordered, V any](m map[K]V) (keys []K) {
	_ = "STUB: not implemented"
	return nil
}

// GetDistincts returns the list of distinct elements in v.
func GetDistincts[V comparable](v []V) (vd []V) { _ = "STUB: not implemented"; return nil }

// SortSlice sorts a slice in place.
func SortSlice[T constraints.Ordered](s []T) { _ = "STUB: not implemented"; return }

// RotateSlice returns a new slice corresponding to s rotated by k positions to the left.
func RotateSlice[V any](s []V, k int) []V { _ = "STUB: not implemented"; return nil }

// RotateSliceAllocFree rotates slice s by k positions to the left and writes the result in sout.
// without allocating new memory.
func RotateSliceAllocFree[V any](s []V, k int, sout []V) { _ = "STUB: not implemented"; return }

// checks if the two slice share the same backing array

// RotateSliceInPlace rotates slice s in place by k positions to the left.
func RotateSliceInPlace[V any](s []V, k int) { _ = "STUB: not implemented"; return }

// RotateSlotsNew returns a new slice where the two half of the
// original slice are rotated each by k positions independently.
func RotateSlotsNew[V any](s []V, k int) (r []V) { _ = "STUB: not implemented"; return nil }

// BitReverseInPlaceSlice applies an in-place bit-reverse permutation on the input slice.
func BitReverseInPlaceSlice[V any](slice []V, N int) { _ = "STUB: not implemented"; return }
