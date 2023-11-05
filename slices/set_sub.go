package slices

// SetSub returns set where all s2 elements removed from s1.
// I.e. all elements presented in s1 and missing in s2.
func SetSub[S ~[]E, E comparable](s1 S, s2 S) S {
	return SetSubBy(me[E], s1, s2)
}

// SetSubBy returns set where all s2 elements removed from s1 by custom key.
// I.e. all elements presented in s1 and missing in s2.
func SetSubBy[S ~[]E, E any, K comparable](byFn func(E) K, s1 S, s2 S) S {
	// all elements seen in s1
	seen := make(map[K]struct{}, len(s2))
	for _, v := range s2 {
		seen[byFn(v)] = struct{}{}
	}

	var out S // capacity is unknown
	for _, v := range s1 {
		if _, ok := seen[byFn(v)]; ok {
			continue // skip it
		}
		out = append(out, v)
	}

	return out
}
