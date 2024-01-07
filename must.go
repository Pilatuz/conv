package conv

// Must checks there is no error.
// Panics otherwise!
func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}

	return v
}

// MustOK checks condition is OK (true).
// Panics otherwise!
func MustOK[T any](v T, ok bool) T {
	if !ok {
		panic("not OK")
	}

	return v
}
