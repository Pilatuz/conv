//go:build go1.23

package conv_test

import (
	"errors"
	"net"
	"slices"
	"testing"

	"github.com/Pilatuz/conv"
)

// eq reports whether two error slices are equal.
func eq[E error](a, b []E) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if error(a[i]) != error(b[i]) {
			return false
		}
	}

	return true
}

// TestUnwrapAllErrors unit tests for the [UnwrapAll] and [UnwrapAllAs] functions.
func TestUnwrapAllErrors(t *testing.T) {
	var Nil error
	err1 := errors.New("err1")
	err2 := errors.New("err2")

	t.Run("UnwrapAll.empty", func(t *testing.T) {
		got := slices.Collect(conv.UnwrapAll())
		if expected := []error{}; !eq(expected, got) {
			t.Errorf("UnwrapAll() = %v, expected %v", got, expected)
		}
	})

	t.Run("UnwrapAll.nil", func(t *testing.T) {
		got := slices.Collect(conv.UnwrapAll(Nil, Nil))
		if expected := []error{}; !eq(expected, got) {
			t.Errorf("UnwrapAll(nil) = %v, expected %v", got, expected)
		}
	})

	t.Run("UnwrapAll.single_error", func(t *testing.T) {
		err := errors.New("single")
		got := slices.Collect(conv.UnwrapAll(err))
		if expected := []error{err}; !eq(expected, got) {
			t.Errorf("UnwrapAll(single) = %v, expected %v", got, expected)
		}
	})

	t.Run("UnwrapAll.1", func(t *testing.T) {
		got := slices.Collect(conv.UnwrapAll(Nil, err1, Nil))
		if expected := []error{err1}; !eq(expected, got) {
			t.Errorf("UnwrapAll(err1) = %v, expected %v", got, expected)
		}
	})

	t.Run("UnwrapAll.2", func(t *testing.T) {
		got := slices.Collect(conv.UnwrapAll(Nil, err2, Nil))
		if expected := []error{err2}; !eq(expected, got) {
			t.Errorf("UnwrapAll(err2) = %v, expected %v", got, expected)
		}
	})

	t.Run("UnwrapAll.12", func(t *testing.T) {
		err := errors.Join(err1, err2)
		got := slices.Collect(conv.UnwrapAll(err))
		if expected := []error{err, err1, err2}; !eq(expected, got) {
			t.Errorf("UnwrapAll([err1, err2]) = %v, expected %v", got, expected)
		}
	})

	t.Run("UnwrapAll.1a", func(t *testing.T) {
		err := &net.DNSError{UnwrapErr: err1}
		got := slices.Collect(conv.UnwrapAll(err))
		if expected := []error{err, err1}; !eq(expected, got) {
			t.Errorf("UnwrapAll(DNSError{err1}) = %v, expected %v", got, expected)
		}
	})

	t.Run("UnwrapAll.12a", func(t *testing.T) {
		err1a := &net.DNSError{UnwrapErr: err1}
		err2a := &net.DNSError{UnwrapErr: err2}
		err := errors.Join(err1a, err2a)
		got := slices.Collect(conv.UnwrapAll(err))
		if expected := []error{err, err1a, err1, err2a, err2}; !eq(expected, got) {
			t.Errorf("UnwrapAll([DNSError{err1}, DNSError{err2}]) = %v, expected %v", got, expected)
		}
	})

	t.Run("UnwrapAll.12b", func(t *testing.T) {
		err0 := errors.Join(err1, err2)
		err1a := &net.DNSError{UnwrapErr: err0}
		err2a := &net.DNSError{}
		err := errors.Join(err1a, err2a)
		got := slices.Collect(conv.UnwrapAll(err))
		if expected := []error{err, err1a, err0, err1, err2, err2a}; !eq(expected, got) {
			t.Errorf("UnwrapAll([DNSError{[err1,err2]}, DNSError{}]) = %v, expected %v", got, expected)
		}
	})

	t.Run("UnwrapAll.12b.break", func(t *testing.T) {
		err0 := errors.Join(err1, err2)
		err1a := &net.DNSError{UnwrapErr: err0}
		err2a := &net.DNSError{}
		err := errors.Join(err1a, err2a)

		expected := []error{err, err1a, err0, err1, err2, err2a}

		for n := range len(expected) {
			var got []error
			for e := range conv.UnwrapAll(Nil, err, Nil) {
				got = append(got, e)
				if len(got) > n {
					break
				}
			}
			if expected := expected[:n+1]; !eq(expected, got) {
				t.Errorf("UnwrapAll([...], break=%d) = %v, expected %v", n, got, expected)
			}
		}
	})

	t.Run("UnwrapAllAs.12b", func(t *testing.T) {
		err0 := errors.Join(err1, err2)
		err1a := &net.DNSError{UnwrapErr: err0}
		err2a := &net.DNSError{}
		err := errors.Join(err1a, err2a)
		got := slices.Collect(conv.UnwrapAllAs[*net.DNSError](err))
		if expected := []*net.DNSError{err1a, err2a}; !eq(expected, got) {
			t.Errorf("UnwrapAll([DNSError{[err1,err2]}, DNSError{}]) = %v, expected %v", got, expected)
		}
	})

	t.Run("UnwrapAllAs.12c", func(t *testing.T) {
		err1a := Error1{msg: "foo"}
		err2a := Error2{msg: "bar"}
		err := errors.Join(err1a, err2a)
		got := slices.Collect(conv.UnwrapAllAs[Error1](err))
		if expected := []Error1{err1a}; !eq(expected, got) {
			t.Errorf("UnwrapAll([Error1, Error2]) = %v, expected %v", got, expected)
		}
	})

	t.Run("UnwrapAllAs.12d", func(t *testing.T) {
		err1a := Error1{msg: "foo"}
		err2a := Error2{msg: "bar"}
		err := errors.Join(err1a, err2a)
		got := slices.Collect(conv.UnwrapAllAs[Error2](err))
		if expected := []Error2{{msg: "foo"}, err2a}; !eq(expected, got) {
			t.Errorf("UnwrapAll([Error1, Error2]) = %v, expected %v", got, expected)
		}
	})

	t.Run("UnwrapAllAs.12d.break", func(t *testing.T) {
		err1a := Error1{msg: "foo"}
		err2a := Error2{msg: "bar"}
		err3a := Error2{msg: "baz"}
		err := errors.Join(err1a, err2a, err3a)

		expected := []Error2{{msg: "foo"}, err2a, err3a}

		for n := range len(expected) {
			var got []Error2
			for e := range conv.UnwrapAllAs[Error2](Nil, err, Nil) {
				got = append(got, e)
				if len(got) > n {
					break
				}
			}
			if expected := expected[:n+1]; !eq(expected, got) {
				t.Errorf("UnwrapAllAs([...], break=%d) = %v, expected %v", n, got, expected)
			}
		}
	})

	t.Run("UnwrapAll.deeply_nested", func(t *testing.T) {
		err1 := errors.New("err1")
		err2 := errors.New("err2")
		err3 := errors.New("err3")
		inner := errors.Join(err1, err2)
		middle := &net.DNSError{UnwrapErr: inner}
		outer := errors.Join(middle, err3)
		got := slices.Collect(conv.UnwrapAll(outer))
		if expected := []error{outer, middle, inner, err1, err2, err3}; !eq(expected, got) {
			t.Errorf("UnwrapAll(deeply_nested) = %v, expected %v", got, expected)
		}
	})

	t.Run("UnwrapAllAs.empty", func(t *testing.T) {
		got := slices.Collect(conv.UnwrapAllAs[*net.DNSError]())
		if expected := []*net.DNSError{}; !eq(expected, got) {
			t.Errorf("UnwrapAllAs() = %v, expected %v", got, expected)
		}
	})

	t.Run("UnwrapAllAs.no_matches", func(t *testing.T) {
		err1 := errors.New("err1")
		err2 := errors.New("err2")
		err := errors.Join(err1, err2)
		got := slices.Collect(conv.UnwrapAllAs[*net.DNSError](err))
		if expected := []*net.DNSError{}; !eq(expected, got) {
			t.Errorf("UnwrapAllAs(no_matches) = %v, expected %v", got, expected)
		}
	})

	t.Run("UnwrapAllAs.nil_input", func(t *testing.T) {
		got := slices.Collect(conv.UnwrapAllAs[*net.DNSError](nil, nil))
		if expected := []*net.DNSError{}; !eq(expected, got) {
			t.Errorf("UnwrapAllAs(nil, nil) = %v, expected %v", got, expected)
		}
	})

	t.Run("UnwrapAllAs.mixed_types", func(t *testing.T) {
		err1 := errors.New("err1")
		dnsErr1 := &net.DNSError{UnwrapErr: err1}
		dnsErr2 := &net.DNSError{}
		opErr := &net.OpError{Err: dnsErr1}
		err := errors.Join(opErr, dnsErr2)
		got := slices.Collect(conv.UnwrapAllAs[*net.DNSError](err))
		if expected := []*net.DNSError{dnsErr1, dnsErr2}; !eq(expected, got) {
			t.Errorf("UnwrapAllAs(mixed_types) = %v, expected %v", got, expected)
		}
	})

	t.Run("UnwrapAllAs.interface_type", func(t *testing.T) {
		err1a := Error1{msg: "foo"}
		err2a := Error2{msg: "bar"}
		err := errors.Join(err1a, err2a)
		// Error1.As converts to Error2, so UnwrapAllAs[Error2] should find both
		got := slices.Collect(conv.UnwrapAllAs[Error2](err))
		if expected := []Error2{{msg: "foo"}, err2a}; !eq(expected, got) {
			t.Errorf("UnwrapAllAs(interface_type) = %v, expected %v", got, expected)
		}
	})
}

type Error1 struct {
	msg string
}

func (e Error1) Error() string {
	return e.msg
}

func (e Error1) As(target any) bool {
	if p, ok := target.(*Error2); ok {
		p.msg = e.msg
		return true
	}
	return false
}

type Error2 struct {
	msg string
}

func (e Error2) Error() string {
	return e.msg
}

func (Error2) As(target any) bool {
	return false
}
