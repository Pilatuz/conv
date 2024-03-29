package conv

import (
	"errors"
)

// ErrorAs is a shortcut for errors.As(err, &T).
func ErrorAs[T any](err error) (T, bool) {
	var target T
	ok := errors.As(err, &target)
	return target, ok
}

// ErrorIsAnyOf check if error is ANY of provided targets.
func ErrorIsAnyOf(err error, targets ...error) bool {
	for _, target := range targets {
		if errors.Is(err, target) {
			return true
		}
	}

	return false
}
