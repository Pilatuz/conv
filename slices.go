package conv

import (
	impl "github.com/Pilatuz/conv/slices"
)

// Slice initializes slice with some values s of the same type.
func Slice[E any](s ...E) []E {
	return s
}

// SliceNotNil gets non-nil slice.
//
// If input slice s is nil then empty slice is returned instead.
func SliceNotNil[S ~[]E, E any](s S) S {
	return impl.NotNil(s)
}
