package conv

// Must returns v if err is nil.
//
// It panics with err otherwise.
func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}

	return v
}

// MustOK returns v if ok is true.
//
// It panics otherwise.
func MustOK[T any](v T, ok bool) T {
	if !ok {
		panic("not OK")
	}

	return v
}
