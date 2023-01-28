package conv_test

import (
	"errors"
	"testing"

	. "github.com/Pilatuz/conv"
)

// TestMust unit tests for `Must` function.
func TestMust(tt *testing.T) {
	tt.Run("good_str", func(t *testing.T) {
		fn := func() (string, error) {
			return "foo", nil
		}

		if e, a := "foo", Must(fn()); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})

	tt.Run("good_int", func(t *testing.T) {
		fn := func() (int, error) {
			return 123, nil
		}

		if e, a := 123, Must(fn()); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})

	tt.Run("bad_str", func(t *testing.T) {
		anErr := errors.New("ups")
		fn := func() (string, error) {
			return "foo", anErr
		}

		defer func() {
			if r := recover(); r != nil {
				if !errors.Is(r.(error), anErr) {
					t.Errorf("unexpected panic type `%v`", r)
				}
			}
		}()

		Must(fn()) // panics!
		t.Error("no expected panic")
	})

	tt.Run("bad_int", func(t *testing.T) {
		anErr := errors.New("ups")
		fn := func() (int, error) {
			return 123, anErr
		}

		defer func() {
			if r := recover(); r != nil {
				if !errors.Is(r.(error), anErr) {
					t.Errorf("unexpected panic type `%v`", r)
				}
			}
		}()

		Must(fn()) // panics!
		t.Error("no expected panic")
	})
}
