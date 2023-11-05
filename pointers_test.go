package conv_test

import (
	"fmt"
	"testing"

	"github.com/Pilatuz/conv"
)

func ExamplePtrFrom() {
	const a = 12345
	const b = "foo"
	fmt.Println(*conv.PtrFrom(a))
	fmt.Println(*conv.PtrFrom(b))
	// Output:
	// 12345
	// foo
}

func ExampleFromPtrOr() {
	var p *int
	fmt.Println(conv.FromPtrOr(p, 12345))
	// Output:
	// 12345
}

func ExampleFromPtrOrFunc() {
	fn := func() int {
		// ... long init ...
		return 12345
	}
	var p *int
	fmt.Println(conv.FromPtrOrFunc(p, fn))
	// Output:
	// 12345
}

func ExampleOmitEmpty() {
	fmt.Println(conv.OmitEmpty(conv.PtrFrom("")))
	fmt.Println(conv.OmitEmpty(conv.PtrFrom(0)))
	// Output:
	// <nil>
	// <nil>
}

func ExampleFirstNonNil() {
	var a *string
	var b *string
	c := conv.PtrFrom("foo")
	var d *string

	fmt.Println(*conv.FirstNonNil(a, b, c, d))
	// Output:
	// foo
}

func ExampleCoalesce() {
	var a, c int
	b := 123

	fmt.Println(conv.Coalesce(0, a, b, c, 0))
	// Output:
	// 123
}

// TestPtrFrom unit tests for `PtrFrom` function.
func TestPtrFrom(tt *testing.T) {
	// string
	tt.Run("str", func(t *testing.T) {
		p := conv.PtrFrom("foo")
		if p == nil {
			t.Error("expected non-nil value, found nil")
		} else if *p != "foo" {
			t.Errorf("expected `%v`, found `%v`", "foo", *p)
		}
	})

	// integer
	tt.Run("int", func(t *testing.T) {
		p := conv.PtrFrom(123)
		if p == nil {
			t.Error("expected non-nil value, found nil")
		} else if *p != 123 {
			t.Errorf("expected `%v`, found `%v`", 123, *p)
		}
	})

	// boolean
	tt.Run("bool", func(t *testing.T) {
		p := conv.PtrFrom(true)
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
		if e, a := "ups", conv.FromPtrOr(p1, "ups"); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		p2 := "foo"
		if e, a := "foo", conv.FromPtrOr(&p2, "ups"); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})

	// integer
	tt.Run("int", func(t *testing.T) {
		var p1 *int
		if e, a := 123, conv.FromPtrOr(p1, 123); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		p2 := 456
		if e, a := 456, conv.FromPtrOr(&p2, 123); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})

	// boolean
	tt.Run("bool", func(t *testing.T) {
		var p1 *bool
		if e, a := true, conv.FromPtrOr(p1, true); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		p2 := true
		if e, a := true, conv.FromPtrOr(&p2, false); a != e {
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
		if e, a := "ups", conv.FromPtrOrFunc(p1, fn); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		p2 := "foo"
		if e, a := "foo", conv.FromPtrOrFunc(&p2, fn); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})

	// integer
	tt.Run("int", func(t *testing.T) {
		fn := func() int { return 123 }

		var p1 *int
		if e, a := 123, conv.FromPtrOrFunc(p1, fn); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		p2 := 456
		if e, a := 456, conv.FromPtrOrFunc(&p2, fn); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})

	// boolean
	tt.Run("bool", func(t *testing.T) {
		fn := func() bool { return false }

		var p1 *bool
		if e, a := false, conv.FromPtrOrFunc(p1, fn); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		p2 := true
		if e, a := true, conv.FromPtrOrFunc(&p2, fn); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})
}

// TestOmitEmpty unit tests for `OmitEmpty` function.
func TestOmitEmpty(tt *testing.T) {
	// string
	tt.Run("str", func(t *testing.T) {
		var p1 *string
		if e, a := (*string)(nil), conv.OmitEmpty(p1); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		var p2 string
		if e, a := (*string)(nil), conv.OmitEmpty(&p2); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		p3 := "foo"
		if e, a := &p3, conv.OmitEmpty(&p3); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})

	// integer
	tt.Run("int", func(t *testing.T) {
		var p1 *int
		if e, a := (*int)(nil), conv.OmitEmpty(p1); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		var p2 int
		if e, a := (*int)(nil), conv.OmitEmpty(&p2); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		p3 := 123
		if e, a := &p3, conv.OmitEmpty(&p3); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})

	// boolean
	tt.Run("bool", func(t *testing.T) {
		var p1 *bool
		if e, a := (*bool)(nil), conv.OmitEmpty(p1); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		var p2 bool
		if e, a := (*bool)(nil), conv.OmitEmpty(&p2); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		p3 := true
		if e, a := &p3, conv.OmitEmpty(&p3); a != e {
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

		if e, a := (*string)(nil), conv.FirstNonNil(p1, nil, p1); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		if e, a := &p2, conv.FirstNonNil(p1, &p2, &p3, p1); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		if e, a := &p3, conv.FirstNonNil(p1, &p3, &p2, p1); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})

	// integer
	tt.Run("int", func(t *testing.T) {
		var p1 *int
		var p2 int
		p3 := 123

		if e, a := (*int)(nil), conv.FirstNonNil(p1, nil, p1); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		if e, a := &p2, conv.FirstNonNil(p1, &p2, &p3, p1); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		if e, a := &p3, conv.FirstNonNil(p1, &p3, &p2, p1); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})

	// boolean
	tt.Run("bool", func(t *testing.T) {
		var p1 *bool
		var p2 bool
		p3 := true

		if e, a := (*bool)(nil), conv.FirstNonNil(p1, nil, p1); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		if e, a := &p2, conv.FirstNonNil(p1, &p2, &p3, p1); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		if e, a := &p3, conv.FirstNonNil(p1, &p3, &p2, p1); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})
}

// TestCoalesce unit tests for `Coalesce` function.
func TestCoalesce(tt *testing.T) {
	// string
	tt.Run("str", func(t *testing.T) {
		var p1 *string
		var p2 string
		p3 := "foo"

		if e, a := (*string)(nil), conv.Coalesce(p1, nil, p1); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		if e, a := &p2, conv.Coalesce(p1, &p2, &p3, p1); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		if e, a := &p3, conv.Coalesce(p1, &p3, &p2, p1); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		if e, a := p3, conv.Coalesce("", p2, p3, ""); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})

	// integer
	tt.Run("int", func(t *testing.T) {
		var p1 *int
		var p2 int
		p3 := 123

		if e, a := (*int)(nil), conv.Coalesce(p1, nil, p1); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		if e, a := &p2, conv.Coalesce(p1, &p2, &p3, p1); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		if e, a := &p3, conv.Coalesce(p1, &p3, &p2, p1); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		if e, a := p3, conv.Coalesce(0, p2, p3, 0); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})

	// boolean
	tt.Run("bool", func(t *testing.T) {
		var p1 *bool
		var p2 bool
		p3 := true

		if e, a := (*bool)(nil), conv.Coalesce(p1, nil, p1); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		if e, a := &p2, conv.Coalesce(p1, &p2, &p3, p1); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		if e, a := &p3, conv.Coalesce(p1, &p3, &p2, p1); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		if e, a := p3, conv.Coalesce(false, p2, p3, false); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})
}
