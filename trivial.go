package conv

// SignedInt a signed integer.
type SignedInt interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// UnsignedInt an unsigned integer.
type UnsignedInt interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// IntToInt converts integers of various types.
// Mostly should be used with [PtrToPtr] function.
func IntToInt[T2, T1 SignedInt | UnsignedInt](v1 T1) T2 {
	return T2(v1) // trivial conversion
}
