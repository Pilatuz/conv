package conv

import "errors"

// ErrorsAs is a shortcut for errors.As(err, &&T).
func ErrorsAs[T any](err error) (T, bool) {
	var target T
	ok := errors.As(err, &target)
	return target, ok
}
