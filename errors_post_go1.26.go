//go:build go1.26

package conv

import (
	"errors"
)

// ErrorAs is a shortcut for errors.As(err, &E).
//
//go:fix inline
func ErrorAs[E error](err error) (E, bool) {
	return errors.AsType[E](err)
}
