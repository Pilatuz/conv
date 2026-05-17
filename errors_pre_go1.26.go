//go:build !go1.26

package conv

import (
	"errors"
)

// ErrorAs is a shortcut for errors.As(err, &E).
func ErrorAs[E error](err error) (E, bool) {
	var target E
	ok := errors.As(err, &target)
	return target, ok
}
