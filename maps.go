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
