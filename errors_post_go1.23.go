//go:build go1.23

package conv

import (
	"iter"
)

// UnwrapAll recursively iterates over all wrapped errors.
//
// UnwrapAll accepts one or more errors and iterates over each error
// and all errors wrapped by it (via Unwrap()).
// The iteration includes the original errors and all nested wrapped errors.
func UnwrapAll(errs ...error) iter.Seq[error] {
	return func(yield func(error) bool) {
		for _, err := range errs {
			if !unwrapAll(err, yield) {
				return
			}
		}
	}
}

// unwrapAll recursively iterates over all wrapped errors.
func unwrapAll(err error, yield func(error) bool) (_continue bool) {
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

// UnwrapAllAs recursively iterates over all wrapped errors of a specific type.
//
// UnwrapAllAs accepts one or more errors and iterates over each error
// of type E that is found by direct type assertion or via the As() method.
// It traverses all wrapped errors recursively.
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
