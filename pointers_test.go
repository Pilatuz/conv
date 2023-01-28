package conv_test

import (
	"testing"

	. "github.com/Pilatuz/conv"
)

// TestPtrFrom unit tests for `PtrFrom` function.
func TestPtrFrom(tt *testing.T) {
	// string
	tt.Run("str", func(t *testing.T) {
		p := PtrFrom("foo")
		if p == nil {
			t.Error("expected non-nil value, found nil")
		} else if *p != "foo" {
			t.Errorf("expected `%v`, found `%v`", "foo", *p)
		}
	})

	// integer
	tt.Run("int", func(t *testing.T) {
		p := PtrFrom(123)
		if p == nil {
			t.Error("expected non-nil value, found nil")
		} else if *p != 123 {
			t.Errorf("expected `%v`, found `%v`", 123, *p)
		}
	})

	// boolean
	tt.Run("bool", func(t *testing.T) {
		p := PtrFrom(true)
		if p == nil {
			t.Error("expected non-nil value, found nil")
		} else if *p != true {
			t.Errorf("expected `%v`, found `%v`", true, *p)
		}
	})
}

// TestFromPtrOr unit tests for `FromPtrOr` function.
func TestFromPtrOr(tt *testing.T) {
	// string
	tt.Run("str", func(t *testing.T) {
		var p1 *string
		if e, a := "ups", FromPtrOr(p1, "ups"); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		p2 := "foo"
		if e, a := "foo", FromPtrOr(&p2, "ups"); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})

	// integer
	tt.Run("int", func(t *testing.T) {
		var p1 *int
		if e, a := 123, FromPtrOr(p1, 123); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		p2 := 456
		if e, a := 456, FromPtrOr(&p2, 123); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})

	// boolean
	tt.Run("bool", func(t *testing.T) {
		var p1 *bool
		if e, a := true, FromPtrOr(p1, true); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		p2 := true
		if e, a := true, FromPtrOr(&p2, false); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})
}

// TestFromPtrOrFunc unit tests for `FromPtrOrFunc` function.
func TestFromPtrOrFunc(tt *testing.T) {
	// string
	tt.Run("str", func(t *testing.T) {
		fn := func() string { return "ups" }

		var p1 *string
		if e, a := "ups", FromPtrOrFunc(p1, fn); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		p2 := "foo"
		if e, a := "foo", FromPtrOrFunc(&p2, fn); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})

	// integer
	tt.Run("int", func(t *testing.T) {
		fn := func() int { return 123 }

		var p1 *int
		if e, a := 123, FromPtrOrFunc(p1, fn); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		p2 := 456
		if e, a := 456, FromPtrOrFunc(&p2, fn); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})

	// boolean
	tt.Run("bool", func(t *testing.T) {
		fn := func() bool { return false }

		var p1 *bool
		if e, a := false, FromPtrOrFunc(p1, fn); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		p2 := true
		if e, a := true, FromPtrOrFunc(&p2, fn); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})
}

// TestOmitEmpty unit tests for `OmitEmpty` function.
func TestOmitEmpty(tt *testing.T) {
	// string
	tt.Run("str", func(t *testing.T) {
		var p1 *string
		if e, a := (*string)(nil), OmitEmpty(p1); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		var p2 string
		if e, a := (*string)(nil), OmitEmpty(&p2); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		p3 := "foo"
		if e, a := &p3, OmitEmpty(&p3); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})

	// integer
	tt.Run("int", func(t *testing.T) {
		var p1 *int
		if e, a := (*int)(nil), OmitEmpty(p1); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		var p2 int
		if e, a := (*int)(nil), OmitEmpty(&p2); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		p3 := 123
		if e, a := &p3, OmitEmpty(&p3); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})

	// boolean
	tt.Run("bool", func(t *testing.T) {
		var p1 *bool
		if e, a := (*bool)(nil), OmitEmpty(p1); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		var p2 bool
		if e, a := (*bool)(nil), OmitEmpty(&p2); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		p3 := true
		if e, a := &p3, OmitEmpty(&p3); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})
}

// TestFirstNonNil unit tests for `FirstNonNil` function.
func TestFirstNonNil(tt *testing.T) {
	// string
	tt.Run("str", func(t *testing.T) {
		var p1 *string
		var p2 string
		p3 := "foo"

		if e, a := (*string)(nil), FirstNonNil(p1, nil, p1); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		if e, a := &p2, FirstNonNil(p1, &p2, &p3, p1); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		if e, a := &p3, FirstNonNil(p1, &p3, &p2, p1); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})

	// integer
	tt.Run("int", func(t *testing.T) {
		var p1 *int
		var p2 int
		p3 := 123

		if e, a := (*int)(nil), FirstNonNil(p1, nil, p1); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		if e, a := &p2, FirstNonNil(p1, &p2, &p3, p1); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		if e, a := &p3, FirstNonNil(p1, &p3, &p2, p1); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})

	// boolean
	tt.Run("bool", func(t *testing.T) {
		var p1 *bool
		var p2 bool
		p3 := true

		if e, a := (*bool)(nil), FirstNonNil(p1, nil, p1); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		if e, a := &p2, FirstNonNil(p1, &p2, &p3, p1); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		if e, a := &p3, FirstNonNil(p1, &p3, &p2, p1); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})
}
