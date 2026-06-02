//go:build go1.23

package conv

import (
	"iter"
)

// AllNotNil returns all not-nil elements.
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

// AllNotNil2 returns all not-nil elements with original indices.
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
