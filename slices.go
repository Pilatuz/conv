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

// SliceToChan converts an array to read-only channel.
func SliceToChan[S ~[]E, E any](s S, chanBufSize int) <-chan E {
	ch := make(chan E, chanBufSize)

	go func() {
		defer close(ch) // close later

		for _, v := range s {
			ch <- v // send to channel
		}
	}()

	return ch
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

	return s[:k] // in-place
}

// Reverse reverses elements.
//
// Returns original slice with elements reversed in-place.
func Reverse[S ~[]E, E any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// SliceTransform transforms each element of slice.
func SliceTransform[S ~[]E, E any, T any](fn func(E) T, s S) []T {
	if s == nil {
		return nil
	}

	out := make([]T, 0, len(s))
	for _, v := range s {
		out = append(out, fn(v))
	}

	return out
}

// me uses element itself as a trivial key.
// Likely this function should be copmletely inlined by compiler.
func me[K comparable](v K) K {
	return v
}

// SliceAnd returns the intersection between two slices.
// I.e. elements presented in both slices.
func SliceAnd[S ~[]E, E comparable](s1 S, s2 S) S {
	return SliceAndBy(me[E], s1, s2)
}

// SliceAndBy returns the intersection between two slices by custom key.
// I.e. elements presented in both slices.
func SliceAndBy[S ~[]E, E any, K comparable](by func(E) K, s1 S, s2 S) S {
	if len(s1) == 0 || len(s2) == 0 {
		return nil
	}

	// all elements seen in s1
	seen := make(map[K]struct{}, len(s1))
	for _, v := range s1 {
		seen[by(v)] = struct{}{}
	}

	var out S // capacity is unknown
	for _, v := range s2 {
		if _, ok := seen[by(v)]; !ok {
			continue // skip it
		}
		out = append(out, v)
	}

	return out
}

// SliceSub returns slice where all s2 elements removed from s1.
// I.e. all elements presented in s1 and missing in s2.
func SliceSub[S ~[]E, E comparable](s1 S, s2 S) S {
	return SliceSubBy(me[E], s1, s2)
}

// SliceSubBy returns slice where all s2 elements removed from s1 by custom key.
// I.e. all elements presented in s1 and missing in s2.
func SliceSubBy[S ~[]E, E any, K comparable](by func(E) K, s1 S, s2 S) S {
	// all elements seen in s1
	seen := make(map[K]struct{}, len(s2))
	for _, v := range s2 {
		seen[by(v)] = struct{}{}
	}

	var out S // capacity is unknown
	for _, v := range s1 {
		if _, ok := seen[by(v)]; ok {
			continue // skip it
		}
		out = append(out, v)
	}

	return out
}

// SliceOr returns the union between two (or more) slices.
// I.e. unique elements presented in at least one slice.
func SliceOr[S ~[]E, E comparable](ss ...S) S {
	return SliceOrBy(me[E], ss...)
}

// SliceOrBy returns the union between two (or more) slices by custom key.
// I.e. unique elements presented in at least one slice.
func SliceOrBy[S ~[]E, E any, K comparable](by func(E) K, ss ...S) S {
	// all elements seen so far
	seen := make(map[K]struct{})

	var out S // capacity is unknown
	for _, s := range ss {
		for _, v := range s {
			k := by(v)
			if _, ok := seen[k]; ok {
				continue // skip it
			}
			seen[k] = struct{}{}
			out = append(out, v)
		}
	}

	return out
}

// SliceDiff returns the difference between two slices.
// Where out1 - elements presented in s1 but missing in s2.
// And out2 - elements presented in s2 but missing in s1.
func SliceDiff[S ~[]E, E comparable](s1 S, s2 S) (out1 S, out2 S) {
	return SliceDiffBy(me[E], s1, s2)
}

// SliceDiff returns the difference between two slices.
// Where out1 - elements presented in s1 but missing in s2.
// And out2 - elements presented in s2 but missing in s1.
func SliceDiffBy[S ~[]E, E any, K comparable](by func(E) K, s1 S, s2 S) (out1 S, out2 S) {
	if len(s1) == 0 || len(s2) == 0 {
		return s1, s2
	}

	// elements seen in s1
	seen1 := make(map[K]struct{}, len(s1))
	for _, v := range s1 {
		seen1[by(v)] = struct{}{}
	}

	// elements seen in s2
	seen2 := make(map[K]struct{}, len(s2))
	for _, v := range s2 {
		seen2[by(v)] = struct{}{}
	}

	// presented in s1, missing in s2
	for _, v := range s1 {
		if _, ok := seen2[by(v)]; ok {
			continue // skip it
		}
		out1 = append(out1, v)
	}

	// presented in s2, missing in s1
	for _, v := range s2 {
		if _, ok := seen1[by(v)]; ok {
			continue
		}
		out2 = append(out2, v)
	}

	return
}
