package conv

// Slice initializes slice with some values s of the same type.
func Slice[E any](s ...E) []E {
	return s
}

// SliceNotNil gets non-nil slice.
//
// If input slice s is nil then empty slice is returned instead.
func SliceNotNil[S ~[]E, E any](s S) S {
	if s != nil {
		return s // as is
	}

	return S{} // empty
}

// SliceJoin joins multiple slices.
func SliceJoin[S ~[]E, E any](ss ...S) S {
	// output's slice capacity
	var n int
	for _, s := range ss {
		n += len(s)
	}

	// join the slices
	out := make(S, 0, n)
	for _, s := range ss {
		out = append(out, s...)
	}

	return out
}

// SliceFromChan drains the channel ch into array.
//
// It collects all the values in channel ch until the channel is closed.
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

// Unique gets only unique elements.
//
// Returns original slice with duplicates removed in-place.
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

// Reverse reverses elements.
//
// Returns original slice with elements reversed in-place.
func Reverse[S ~[]E, E any](s S) S {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s // the same
}
