//go:build !go1.26

package conv

import (
	"errors"
)

// ErrorAs is a shortcut for errors.As(err, &target) of type E.
//
// The ok result reports whether a match was found.
func ErrorAs[E error](err error) (E, bool) {
	var target E
	ok := errors.As(err, &target)
	return target, ok
}
