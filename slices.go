package conv

import (
	impl "github.com/Pilatuz/conv/slices"
)

// Slice initializes slice with some values s of the same type.
func Slice[E any](s ...E) []E {
	return impl.New(s...)
}

// SliceNotNil gets non-nil slice.
//
// If input slice s is nil then empty slice is returned instead.
func SliceNotNil[S ~[]E, E any](s S) S {
	return impl.NotNil(s)
}

// SliceJoin joins multiple slices.
func SliceJoin[S ~[]E, E any](ss ...S) S {
	return impl.Join(ss...)
}

// SliceFromChan drains the channel ch into array.
//
// It collects all the values in channel ch until the channel is closed.
func SliceFromChan[E any](ch <-chan E) []E {
	return impl.FromChan(ch)
}

// SliceToChan converts an array to read-only channel.
func SliceToChan[S ~[]E, E any](s S, chanBufSize int) <-chan E {
	return impl.ToChan(s, chanBufSize)
}

// Unique gets only unique elements.
//
// Returns original slice with duplicates removed in-place.
func Unique[S ~[]E, E comparable](s S) S {
	return impl.Unique(s)
}

// Reverse reverses elements.
//
// Returns original slice with elements reversed in-place.
func Reverse[S ~[]E, E any](s S) {
	impl.Reverse(s)
}

// SliceTransform transforms each element of slice.
func SliceTransform[S ~[]E, E any, T any](fn func(E) T, s S) []T {
	return impl.Transform(fn, s)
}

// SliceAnd returns the intersection between two slices.
// I.e. elements presented in both slices.
func SliceAnd[S ~[]E, E comparable](s1 S, s2 S) S {
	return impl.And(s1, s2)
}

// SliceAndBy returns the intersection between two slices by custom key.
// I.e. elements presented in both slices.
func SliceAndBy[S ~[]E, E any, K comparable](by func(E) K, s1 S, s2 S) S {
	return impl.AndBy(by, s1, s2)
}

// SliceSub returns slice where all s2 elements removed from s1.
// I.e. all elements presented in s1 and missing in s2.
func SliceSub[S ~[]E, E comparable](s1 S, s2 S) S {
	return impl.Sub(s1, s2)
}

// SliceSubBy returns slice where all s2 elements removed from s1 by custom key.
// I.e. all elements presented in s1 and missing in s2.
func SliceSubBy[S ~[]E, E any, K comparable](by func(E) K, s1 S, s2 S) S {
	return impl.SubBy(by, s1, s2)
}

// SliceOr returns the union between two (or more) slices.
// I.e. unique elements presented in at least one slice.
func SliceOr[S ~[]E, E comparable](ss ...S) S {
	return impl.Or(ss...)
}

// SliceOrBy returns the union between two (or more) slices by custom key.
// I.e. unique elements presented in at least one slice.
func SliceOrBy[S ~[]E, E any, K comparable](by func(E) K, ss ...S) S {
	return impl.OrBy(by, ss...)
}

// SliceDiff returns the difference between two slices.
// Where out1 - elements presented in s1 but missing in s2.
// And out2 - elements presented in s2 but missing in s1.
func SliceDiff[S ~[]E, E comparable](s1 S, s2 S) (out1 S, out2 S) {
	return impl.Diff(s1, s2)
}

// SliceDiff returns the difference between two slices.
// Where out1 - elements presented in s1 but missing in s2.
// And out2 - elements presented in s2 but missing in s1.
func SliceDiffBy[S ~[]E, E any, K comparable](by func(E) K, s1 S, s2 S) (out1 S, out2 S) {
	return impl.DiffBy(by, s1, s2)
}
