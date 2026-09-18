//go:build go1.23

package conv

import (
	"iter"
)

// AllNotNil returns an iterator over all the non-nil pointers of pp.
//
// Use [AllNotNil2] if the original indices are also needed.
func AllNotNil[T any](pp ...*T) iter.Seq[*T] {
	return func(yield func(*T) bool) {
		for _, p := range pp {
			if p == nil {
				continue
			}

			if !yield(p) {
				return
			}
		}
	}
}

// AllNotNil2 returns an iterator over all the non-nil pointers of pp
// along with their original indices.
func AllNotNil2[T any](pp ...*T) iter.Seq2[int, *T] {
	return func(yield func(int, *T) bool) {
		for i, p := range pp {
			if p == nil {
				continue
			}

			if !yield(i, p) {
				return
			}
		}
	}
}
