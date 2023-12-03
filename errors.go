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
