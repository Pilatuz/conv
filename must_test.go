package conv_test

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"testing"

	. "github.com/Pilatuz/conv"
)

func ExampleMust() {
	i := Must(strconv.Atoi("123"))
	fmt.Println(i)
	// Output: 123
}

// TestMust unit tests for [Must] function.
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

func ExampleMustOK() {
	err := &net.ParseError{}
	w := MustOK(ErrorAs[net.Error](err))
	fmt.Println(w.Timeout())
	// Output: false
}

// TestMustOK unit tests for [MustOK] function.
func TestMustOK(tt *testing.T) {
	tt.Run("good_str", func(t *testing.T) {
		fn := func() (string, bool) {
			return "foo", true
		}

		if e, a := "foo", MustOK(fn()); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})

	tt.Run("good_int", func(t *testing.T) {
		fn := func() (int, bool) {
			return 123, true
		}

		if e, a := 123, MustOK(fn()); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})

	tt.Run("bad_str", func(t *testing.T) {
		fn := func() (string, bool) {
			return "foo", false
		}

		defer func() {
			if r := recover(); r != nil {
				if _, ok := r.(string); !ok {
					t.Errorf("unexpected panic type `%v`", r)
				}
			}
		}()

		MustOK(fn()) // panics!
		t.Error("no expected panic")
	})

	tt.Run("bad_int", func(t *testing.T) {
		fn := func() (int, bool) {
			return 123, false
		}

		defer func() {
			if r := recover(); r != nil {
				if _, ok := r.(string); !ok {
					t.Errorf("unexpected panic type `%v`", r)
				}
			}
		}()

		MustOK(fn()) // panics!
		t.Error("no expected panic")
	})
}
