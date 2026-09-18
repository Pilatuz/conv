package conv

// SignedInt is a constraint for all signed integer types.
type SignedInt interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// UnsignedInt is a constraint for all unsigned integer types.
type UnsignedInt interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// IntToInt converts an integer of one type to an integer of another type.
//
// The conversion is unchecked: values that do not fit into T2 are
// truncated, exactly as the built-in T2(v1) conversion does.
// It is mostly intended to be used with the [PtrToPtr] function.
func IntToInt[T2, T1 SignedInt | UnsignedInt](v1 T1) T2 {
	return T2(v1) // trivial conversion
}
