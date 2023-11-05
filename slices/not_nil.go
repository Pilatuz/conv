package slices

// NotNil gets non-nil slice.
//
// If input slice s is nil then empty slice is returned instead.
func NotNil[S ~[]E, E any](s S) S {
	if s != nil {
		return s // as is
	}

	return S{} // empty
}
