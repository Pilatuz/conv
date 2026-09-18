package conv

// MapNotNil returns a non-nil map.
//
// If the input map m is nil then an empty map is returned instead.
// It is the inverse of [MapOmitEmpty].
func MapNotNil[M ~map[K]V, K comparable, V any](m M) M {
	if m != nil {
		return m // as is
	}

	return M{} // empty
}

// MapOmitEmpty returns nil if the input map m is empty.
//
// Both nil and non-nil maps of zero length are reported as empty.
// It is the inverse of [MapNotNil].
func MapOmitEmpty[M ~map[K]V, K comparable, V any](m M) M {
	if len(m) == 0 {
		return nil
	}

	return m // as is
}
