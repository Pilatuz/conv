package conv

// Slice initializes slice with some values.
func Slice[E any](s ...E) []E {
	return s
}

// SliceNotNil gets non-`nil` slice.
func SliceNotNil[S ~[]E, E any](s S) S {
	if s != nil {
		return s // as is
	}

	return S{} // empty
}

// SliceFromChan drains the channel into array.
func SliceFromChan[E any](ch <-chan E) []E {
	if ch == nil {
		return nil
	}

	var out []E // capacity is unknown
	for v := range ch {
		out = append(out, v)
	}

	return out
}

// Unique gets only unique elements (in-place).
func Unique[S ~[]E, E comparable](s S) S {
	if len(s) <= 1 {
		return s // as is
	}

	seen := make(map[E]struct{}, len(s))

	k := 0 // write position, total output elements
	for _, v := range s {
		if _, ok := seen[v]; ok {
			continue
		}

		seen[v] = struct{}{}
		s[k] = v
		k++
	}

	return s[:k]
}

// Reverse reverses elements (in-place).
func Reverse[S ~[]E, E any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
