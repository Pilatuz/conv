package conv

// PtrFrom gets pointer from any value or constant v.
func PtrFrom[T any](v T) *T {
	return &v
}

// FromPtrOr dereferences pointer p or use fallback value ifNil if pointer is nil.
func FromPtrOr[T any](p *T, ifNil T) T {
	if p != nil {
		return *p
	}

	return ifNil
}

// FromPtrOrFunc dereferences pointer p or use fallback function ifNilFn call result if pointer p is nil.
//
// The main difference from [FromPtrOr] - the fallback value is lazy initialized here.
func FromPtrOrFunc[T any](p *T, ifNilFn func() T) T {
	if p != nil {
		return *p
	}

	return ifNilFn()
}

// OmitEmpty returns nil pointer if value *p is empty (or default).
//
// Is used to get nil pointer instead of empty string or zero integer.
func OmitEmpty[T comparable](p *T) *T {
	if p != nil {
		var EMPTY T
		if *p == EMPTY {
			return nil
		}
	}

	return p // as is
}

// PtrToPtr converts T2 to T1 via pointers using conversion function.
// Nil converted to nil.
func PtrToPtr[T1, T2 any](convFn func(T2) T1, p2 *T2) *T1 {
	if p2 == nil {
		return nil // nil -> nil
	}

	v1 := convFn(*p2)
	return &v1
}

// AnyFromPtr converts a pointer to any interface.
// Nil pointer converted to nil.
func AnyFromPtr[T any](p *T) any {
	if p == nil {
		return nil
	}

	return p
}

// FirstNonNil gets first non-nil pointer.
// It works similar to SQL COALESCE function.
func FirstNonNil[T any](pp ...*T) *T {
	return Coalesce(pp...)
}

// Coalesce gets first non-empty value.
// Similar to [FirstNonNil] but works with any types, not just pointers.
func Coalesce[T comparable](vv ...T) T {
	out, _ := CoalesceEx(vv...)
	return out // ignore OK status
}

// CoalesceEx gets first non-empty value with OK status.
func CoalesceEx[T comparable](vv ...T) (out T, ok bool) {
	for _, v := range vv {
		if v == out {
			continue // is empty
		}

		return v, true
	}

	return // out, false
}
