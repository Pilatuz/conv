package conv

// PtrFrom gets pointer from any value.
func PtrFrom[T any](v T) *T {
	return &v
}

// FromPtrOr dereferences pointer or use fallback value if pointer is `nil`.
func FromPtrOr[T any](p *T, ifNil T) T {
	if p != nil {
		return *p
	}

	return ifNil
}

// FromPtrOrFunc dereferences pointer or use fallback function call result if pointer is `nil`.
func FromPtrOrFunc[T any](p *T, ifNilFn func() T) T {
	if p != nil {
		return *p
	}

	return ifNilFn()
}

// OmitEmpty returns `nil` pointer if value is empty (default).
func OmitEmpty[T comparable](p *T) *T {
	if p != nil {
		var EMPTY T
		if *p == EMPTY {
			return nil
		}
	}

	return p // as is
}

// FirstNonNil gets first non-`nil` pointer.
func FirstNonNil[T any](pp ...*T) *T {
	for _, p := range pp {
		if p != nil {
			return p // found
		}
	}

	return nil // all `nil`
}
