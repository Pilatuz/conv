package conv

import (
	"errors"
)

// ErrorIsAnyOf check if error is ANY of provided targets.
func ErrorIsAnyOf(err error, targets ...error) bool {
	for _, target := range targets {
		if errors.Is(err, target) {
			return true
		}
	}

	return false
}
