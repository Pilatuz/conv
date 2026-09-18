package conv

// PtrFrom returns a pointer to the value or constant v.
func PtrFrom[T any](v T) *T {
	return &v
}

// FromPtrOr returns the value pointed to by p, or ifNil if p is nil.
func FromPtrOr[T any](p *T, ifNil T) T {
	if p != nil {
		return *p
	}

	return ifNil
}

// FromPtrOrFunc returns the value pointed to by p, or the result of ifNilFn if p is nil.
//
// Unlike [FromPtrOr], the fallback value is initialized lazily:
// ifNilFn is called only when p is nil.
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

// PtrToPtr converts T1 to T2 via pointers using conversion function.
// Nil converted to nil.
func PtrToPtr[T2, T1 any](p1 *T1, convFn func(T1) T2) *T2 {
	if p1 == nil {
		return nil // nil -> nil
	}

	v2 := convFn(*p1)
	return &v2
}

// AnyFromPtr converts a pointer to the any interface.
//
// A nil pointer is converted to a nil interface, so the result
// is never a non-nil interface holding a nil pointer.
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


// Coalesce returns the first non-zero value of vv, or the zero value if there is none.
//
// It is similar to [FirstNonNil] but works with any comparable type, not just pointers.
// Use [CoalesceEx] to distinguish "no non-zero value found" from "the zero value was found".
func Coalesce[T comparable](vv ...T) T {
	out, _ := CoalesceEx(vv...)
	return out // ignore OK status
}

// CoalesceEx returns the first non-zero value of vv.
//
// The ok result reports whether such a value was found.
func CoalesceEx[T comparable](vv ...T) (out T, ok bool) {
	for _, v := range vv {
		if v == out {
			continue // is empty
		}

		return v, true
	}

	return // out, false
}
