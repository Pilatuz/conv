package conv

import (
	"errors"
)

// ErrorIsAnyOf reports whether err matches any of the provided targets.
//
// Each target is checked with [errors.Is], so wrapped errors are matched too.
// It returns false if no targets are provided.
func ErrorIsAnyOf(err error, targets ...error) bool {
	for _, target := range targets {
		if errors.Is(err, target) {
			return true
		}
	}

	return false
}
