package conv

// Slice returns a slice containing the values s of the same type.
func Slice[E any](s ...E) []E {
	return s
}

// SliceNotNil returns a non-nil slice.
//
// If the input slice s is nil then an empty slice is returned instead.
// It is the inverse of [SliceOmitEmpty].
func SliceNotNil[S ~[]E, E any](s S) S {
	if s != nil {
		return s // as is
	}

	return S{} // empty
}

// SliceOmitEmpty returns nil if the input slice s is empty.
//
// Both nil and non-nil slices of zero length are reported as empty.
// It is the inverse of [SliceNotNil].
func SliceOmitEmpty[S ~[]E, E any](s S) S {
	if len(s) == 0 {
		return nil
	}

	return s // as is
}
