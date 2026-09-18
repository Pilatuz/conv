//go:build go1.23

package conv

import (
	"iter"
)

// UnwrapAll returns an iterator over errs and all the errors they wrap.
//
// Each error is unwrapped recursively via its Unwrap() error
// or Unwrap() []error method. The iteration includes the original
// errors themselves and is done in depth-first pre-order.
// Nil errors are skipped.
func UnwrapAll(errs ...error) iter.Seq[error] {
	return func(yield func(error) bool) {
		for _, err := range errs {
			if !unwrapAll(err, yield) {
				return
			}
		}
	}
}

// unwrapAll yields err and all the errors it wraps.
//
// It returns false if the consumer has stopped the iteration.
func unwrapAll(err error, yield func(error) bool) bool {
	if err == nil {
		return true
	}

	if !yield(err) {
		return false
	}

	switch x := err.(type) {
	case interface{ Unwrap() error }:
		if !unwrapAll(x.Unwrap(), yield) {
			return false
		}

	case interface{ Unwrap() []error }:
		for _, err := range x.Unwrap() {
			if !unwrapAll(err, yield) {
				return false
			}
		}
	}

	return true
}

// UnwrapAllAs returns an iterator over all the errors of type E
// found in errs and in the errors they wrap.
//
// It walks the same error tree as [UnwrapAll] and reports each error
// that is either of type E directly or convertible to E via its
// As(any) bool method.
//
// Note that an error satisfying both conditions is reported twice.
func UnwrapAllAs[E error](errs ...error) iter.Seq[E] {
	return func(yield func(E) bool) {
		for err := range UnwrapAll(errs...) {
			if e, ok := err.(E); ok {
				if !yield(e) {
					return
				}
			}

			type AsI interface{ As(any) bool }
			if x, ok := err.(AsI); ok {
				var e E
				if x.As(&e) {
					if !yield(e) {
						return
					}
				}
			}
		}
	}
}
