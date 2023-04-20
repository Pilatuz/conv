package conv

// MapNotNil gets non-nil map.
//
// If input map m is nil then empty map is returned instead.
func MapNotNil[M ~map[K]V, K comparable, V any](m M) M {
	if m != nil {
		return m // as is
	}

	return M{} // empty
}
