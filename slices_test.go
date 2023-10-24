package conv_test

import (
	"testing"

	. "github.com/Pilatuz/conv"
)

// TestSlice unit tests for `Slice` function.
func TestSlice(tt *testing.T) {
	// string
	tt.Run("str", func(t *testing.T) {
		// Slice() gives nil
		if e, a := ([]string)(nil), Slice[string](); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// Slice([]E{}...) gives empty slice
		if e, a := []string{}, Slice([]string{}...); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// Slice(a, b, ...) gives slice [a, b, ...]
		if e, a := []string{"foo", "bar"}, Slice("foo", "bar"); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})

	// integer
	tt.Run("int", func(t *testing.T) {
		// Slice() gives nil
		if e, a := ([]int)(nil), Slice[int](); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// Slice([]E{}...) gives empty slice
		if e, a := []int{}, Slice([]int{}...); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// Slice(a, b, ...) gives slice [a, b, ...]
		if e, a := []int{123, 456}, Slice(123, 456); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})

	// boolean
	tt.Run("bool", func(t *testing.T) {
		// Slice() gives nil
		if e, a := ([]bool)(nil), Slice[bool](); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// Slice([]E{}...) gives empty slice
		if e, a := []bool{}, Slice([]bool{}...); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// Slice(a, b, ...) gives slice [a, b, ...]
		if e, a := []bool{true, false}, Slice(true, false); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})
}

// TestSliceNotNil unit tests for `SliceNotNil` function.
func TestSliceNotNil(tt *testing.T) {
	// string
	tt.Run("str", func(t *testing.T) {
		// SliceNotNil(nil) gives empty slice
		if e, a := []string{}, SliceNotNil[[]string](nil); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// SliceNotNil([]) gives empty slice
		if e, a := []string{}, SliceNotNil([]string{}); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// SliceNotNil([a, b, ...]) gives slice [a, b, ...]
		if e, a := []string{"foo", "bar"}, SliceNotNil([]string{"foo", "bar"}); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})

	// integer
	tt.Run("int", func(t *testing.T) {
		// SliceNotNil(nil) gives empty slice
		if e, a := []int{}, SliceNotNil[[]int](nil); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// SliceNotNil([]) gives empty slice
		if e, a := []int{}, SliceNotNil([]int{}); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// SliceNotNil([a, b, ...]) gives slice [a, b, ...]
		if e, a := []int{123, 456}, SliceNotNil([]int{123, 456}); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})

	// boolean
	tt.Run("bool", func(t *testing.T) {
		// SliceNotNil(nil) gives empty slice
		if e, a := []bool{}, SliceNotNil[[]bool](nil); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// SliceNotNil([]) gives empty slice
		if e, a := []bool{}, SliceNotNil([]bool{}); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// SliceNotNil([a, b, ...]) gives slice [a, b, ...]
		if e, a := []bool{true, false}, SliceNotNil([]bool{true, false}); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})
}

// TestSliceJoin unit tests for `SliceJoin` function.
func TestSliceJoin(tt *testing.T) {
	// string
	tt.Run("str", func(t *testing.T) {
		var A []string
		B := []string{}
		C := []string{"foo", "foo2"}
		D := []string{"bar", "bar2"}

		// SliceJoin() gives empty slice
		if e, a := []string{}, SliceJoin[[]string](); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// SliceJoin(nil, []) gives empty slice
		if e, a := []string{}, SliceJoin(A, B); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// SliceJoin(C, D) gives slice [C,D]
		if e, a := []string{"foo", "foo2", "bar", "bar2"}, SliceJoin(A, C, B, D); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})

	// integer
	tt.Run("int", func(t *testing.T) {
		var A []int
		B := []int{}
		C := []int{1, 2}
		D := []int{3, 4}

		// SliceJoin() gives empty slice
		if e, a := []int{}, SliceJoin[[]int](); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// SliceJoin(nil, []) gives empty slice
		if e, a := []int{}, SliceJoin(A, B); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// SliceJoin(C, D) gives slice [C,D]
		if e, a := []int{1, 2, 3, 4}, SliceJoin(A, C, B, D); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})
}

// TestSliceFromChan unit tests for `SliceFromChan` function.
func TestSliceFromChan(tt *testing.T) {
	// string
	tt.Run("str", func(t *testing.T) {
		// SliceFromChan(nil) gives nil
		if e, a := ([]string)(nil), SliceFromChan[string](nil); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// SliceFromChan
		ch := make(chan string)
		go func() {
			defer close(ch)
			ch <- "foo"
			ch <- "bar"
		}()
		if e, a := []string{"foo", "bar"}, SliceFromChan(ch); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})

	// integer
	tt.Run("int", func(t *testing.T) {
		// SliceFromChan(nil) gives nil
		if e, a := ([]int)(nil), SliceFromChan[int](nil); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// SliceFromChan
		ch := make(chan int)
		go func() {
			defer close(ch)
			ch <- 123
			ch <- 456
		}()
		if e, a := []int{123, 456}, SliceFromChan(ch); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})

	// boolean
	tt.Run("bool", func(t *testing.T) {
		// SliceFromChan(nil) gives nil
		if e, a := ([]bool)(nil), SliceFromChan[bool](nil); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// SliceFromChan
		ch := make(chan bool)
		go func() {
			defer close(ch)
			ch <- true
			ch <- false
		}()
		if e, a := []bool{true, false}, SliceFromChan(ch); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})
}

// TestSliceToChan unit tests for `SliceToChan` function.
func TestSliceToChan(tt *testing.T) {
	// string
	tt.Run("str", func(t *testing.T) {
		// SliceToChan(nil) gives not nil, but empty channel
		var ss []string
		ch := SliceToChan(ss, 0)
		if ch == nil {
			t.Errorf("expected not-nil channel")
		}
		v, ok := <-ch
		if ok || v != "" {
			t.Errorf("expected empty channel, found `%v`", v)
		}

		ss = []string{"foo", "bar"}
		if e, a := ss, SliceFromChan(SliceToChan(ss, 0)); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})

	// integer
	tt.Run("int", func(t *testing.T) {
		// SliceToChan(nil) gives not nil, but empty channel
		var ss []int
		ch := SliceToChan(ss, 0)
		if ch == nil {
			t.Errorf("expected not-nil channel")
		}
		v, ok := <-ch
		if ok || v != 0 {
			t.Errorf("expected empty channel, found `%v`", v)
		}

		ss = []int{123, 456}
		if e, a := ss, SliceFromChan(SliceToChan(ss, 0)); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})

	// boolean
	tt.Run("bool", func(t *testing.T) {
		// SliceToChan(nil) gives not nil, but empty channel
		var ss []bool
		ch := SliceToChan(ss, 0)
		if ch == nil {
			t.Errorf("expected not-nil channel")
		}
		v, ok := <-ch
		if ok || v != false {
			t.Errorf("expected empty channel, found `%v`", v)
		}

		ss = []bool{true, false}
		if e, a := ss, SliceFromChan(SliceToChan(ss, 0)); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})
}

// TestUnique unit tests for `Unique` function.
func TestUnique(tt *testing.T) {
	// string
	tt.Run("str", func(t *testing.T) {
		// Unique(nil) gives nil
		if e, a := ([]string)(nil), Unique[[]string](nil); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// Unique([]) gives empty slice
		if e, a := []string{}, Unique([]string{}); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// Unique([a, a, b, b, ...]) gives slice [a, b, ...]
		if e, a := []string{"foo", "bar"}, Unique([]string{"foo", "foo", "bar", "foo", "bar"}); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})

	// integer
	tt.Run("int", func(t *testing.T) {
		// Unique(nil) gives nil
		if e, a := ([]int)(nil), Unique[[]int](nil); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// Unique([]) gives empty slice
		if e, a := []int{}, Unique([]int{}); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// Unique([a, a, b, b, ...]) gives slice [a, b, ...]
		if e, a := []int{123, 456}, Unique([]int{123, 123, 456, 456, 123}); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})

	// boolean
	tt.Run("bool", func(t *testing.T) {
		// Unique(nil) gives nil
		if e, a := ([]bool)(nil), Unique[[]bool](nil); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// Unique([]) gives empty slice
		if e, a := []bool{}, Unique([]bool{}); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// Unique([a, a, b, b, ...]) gives slice [a, b, ...]
		if e, a := []bool{true, false}, Unique([]bool{true, true, false, false, true}); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})
}

// TestReverse unit tests for `Reverse` function.
func TestReverse(tt *testing.T) {
	// string
	tt.Run("str", func(t *testing.T) {
		var s1 []string
		s2 := []string{}
		s3 := []string{"foo"}
		s4 := []string{"foo", "bar"}
		s5 := []string{"foo", "bar", "baz"}

		// Reverse(nil) gives nil
		Reverse(s1)
		if e, a := ([]string)(nil), s1; !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// Reverse([]) gives empty slice
		Reverse(s2)
		if e, a := []string{}, s2; !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// Reverse([a, b, ...]) gives slice [..., b, a]
		Reverse(s3)
		Reverse(s4)
		Reverse(s5)
		if e, a := []string{"foo"}, s3; !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
		if e, a := []string{"bar", "foo"}, s4; !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
		if e, a := []string{"baz", "bar", "foo"}, s5; !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})

	// integer
	tt.Run("int", func(t *testing.T) {
		var s1 []int
		s2 := []int{}
		s3 := []int{123}
		s4 := []int{123, 456}
		s5 := []int{123, 456, 789}

		// Reverse(nil) gives nil
		Reverse(s1)
		if e, a := ([]int)(nil), s1; !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// Reverse([]) gives empty slice
		Reverse(s2)
		if e, a := []int{}, s2; !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// Reverse([a, b, ...]) gives slice [..., b, a]
		Reverse(s3)
		Reverse(s4)
		Reverse(s5)
		if e, a := []int{123}, s3; !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
		if e, a := []int{456, 123}, s4; !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
		if e, a := []int{789, 456, 123}, s5; !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})

	// boolean
	tt.Run("bool", func(t *testing.T) {
		var s1 []bool
		s2 := []bool{}
		s3 := []bool{true}
		s4 := []bool{true, false}
		s5 := []bool{true, true, false}

		// Reverse(nil) gives nil
		Reverse(s1)
		if e, a := ([]bool)(nil), s1; !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// Reverse([]) gives empty slice
		Reverse(s2)
		if e, a := []bool{}, s2; !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// Reverse([a, b, ...]) gives slice [..., b, a]
		Reverse(s3)
		Reverse(s4)
		Reverse(s5)
		if e, a := []bool{true}, s3; !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
		if e, a := []bool{false, true}, s4; !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
		if e, a := []bool{false, true, true}, s5; !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})
}

// TestSliceAnd unit tests for `SliceAnd` function.
func TestSliceAnd(tt *testing.T) {
	// string
	tt.Run("str", func(t *testing.T) {
		var Nil []string
		Empty := []string{}
		Foo := []string{"foo"}
		Bar := []string{"bar"}
		FooBar := []string{"foo", "bar"}
		FooBarBaz := []string{"foo", "bar", "baz"}

		test := func(a, b []string, e1 []string) {
			t.Helper()
			if a1 := SliceAnd(a, b); !sliceEqual(a1, e1) {
				t.Errorf("SliceAnd(`%#v`,`%#v`)=`%#v`, expected `%#v`", a, b, a1, e1)
			}
		}

		test(Nil, Nil, Nil)     // nil && nil => nil
		test(Nil, Empty, Nil)   // nil && [] => nil
		test(Empty, Nil, Nil)   // [] && nil => nil
		test(Empty, Empty, Nil) // [] && [] => nil
		test(Nil, Foo, Nil)     // nil && [foo] => nil
		test(Empty, Foo, Nil)   // [] && [foo] => nil
		test(Foo, Nil, Nil)     // [foo] && nil => nil
		test(Foo, Empty, Nil)   // [foo] && [] => nil

		test(Foo, Bar, Nil) // [foo] && [bar] => nil
		test(Bar, Foo, Nil) // [foo] && [bar] => nil

		test(FooBar, Foo, Foo)    // [foo bar] && [foo] => [foo]
		test(Foo, FooBar, Foo)    // [foo] && [foo bar] => [foo]
		test(FooBarBaz, Foo, Foo) // [foo bar baz] && [foo] => [foo]
		test(Foo, FooBarBaz, Foo) // [foo] && [foo bar baz] => [foo]

		test(FooBar, FooBarBaz, FooBar)       // [foo bar baz] && [foo bar] => [foo bar]
		test(FooBarBaz, FooBar, FooBar)       // [foo bar] && [foo bar baz] => [foo bar]
		test(FooBarBaz, FooBarBaz, FooBarBaz) // [foo bar baz] && [foo bar baz] => [foo bar baz]
		test(FooBar, FooBar, FooBar)          // [foo bar] && [foo bar] => [foo bar]
		test(Foo, Foo, Foo)                   // [foo] && [foo] => [foo]
	})

	// integer
	tt.Run("int", func(t *testing.T) {
		var Nil []int
		Empty := []int{}
		Foo := []int{123}
		Bar := []int{456}
		FooBar := []int{123, 456}
		FooBarBaz := []int{123, 456, 789}

		test := func(a, b []int, e1 []int) {
			t.Helper()
			if a1 := SliceAnd(a, b); !sliceEqual(a1, e1) {
				t.Errorf("SliceAnd(`%#v`,`%#v`)=`%#v`, expected `%#v`", a, b, a1, e1)
			}
		}

		test(Nil, Nil, Nil)     // nil && nil => nil
		test(Nil, Empty, Nil)   // nil && [] => nil
		test(Empty, Nil, Nil)   // [] && nil => nil
		test(Empty, Empty, Nil) // [] && [] => nil
		test(Nil, Foo, Nil)     // nil && [foo] => nil
		test(Empty, Foo, Nil)   // [] && [foo] => nil
		test(Foo, Nil, Nil)     // [foo] && nil => nil
		test(Foo, Empty, Nil)   // [foo] && [] => nil

		test(Foo, Bar, Nil) // [foo] && [bar] => nil
		test(Bar, Foo, Nil) // [foo] && [bar] => nil

		test(FooBar, Foo, Foo)    // [foo bar] && [foo] => [foo]
		test(Foo, FooBar, Foo)    // [foo] && [foo bar] => [foo]
		test(FooBarBaz, Foo, Foo) // [foo bar baz] && [foo] => [foo]
		test(Foo, FooBarBaz, Foo) // [foo] && [foo bar baz] => [foo]

		test(FooBar, FooBarBaz, FooBar)       // [foo bar] && [foo bar baz] => [foo bar]
		test(FooBarBaz, FooBar, FooBar)       // [foo bar baz] && [foo bar] => [foo bar]
		test(FooBarBaz, FooBarBaz, FooBarBaz) // [foo bar baz] && [foo bar baz] => [foo bar baz]
		test(FooBar, FooBar, FooBar)          // [foo bar] && [foo bar] => [foo bar]
		test(Foo, Foo, Foo)                   // [foo] && [foo] => [foo]
	})
}

// TestSliceOr unit tests for `SliceOr` function.
func TestSliceOr(tt *testing.T) {
	// string
	tt.Run("str", func(t *testing.T) {
		var Nil []string
		Empty := []string{}
		Foo := []string{"foo"}
		Bar := []string{"bar"}
		FooBar := []string{"foo", "bar"}
		BarFoo := []string{"bar", "foo"}
		FooBarBaz := []string{"foo", "bar", "baz"}

		test := func(a, b []string, e1 []string) {
			t.Helper()
			if a1 := SliceOr(a, b); !sliceEqual(a1, e1) {
				t.Errorf("SliceOr(`%#v`,`%#v`)=`%#v`, expected `%#v`", a, b, a1, e1)
			}
		}

		test(Nil, Nil, Nil)     // nil || nil => nil
		test(Nil, Empty, Nil)   // nil || [] => nil
		test(Empty, Nil, Nil)   // [] || nil => nil
		test(Empty, Empty, Nil) // [] || [] => nil
		test(Nil, Foo, Foo)     // nil || [foo] => [foo]
		test(Empty, Foo, Foo)   // [] && [foo] => [foo]
		test(Foo, Nil, Foo)     // [foo] && nil => [foo]
		test(Foo, Empty, Foo)   // [foo] && [] => [foo]

		test(Foo, Bar, FooBar) // [foo] || [bar] => [foo bar]
		test(Bar, Foo, BarFoo) // [foo] || [bar] => [bar foo]

		test(FooBar, Foo, FooBar)       // [foo bar] || [foo] => [foo bar]
		test(Foo, FooBar, FooBar)       // [foo] || [foo bar] => [foo bar]
		test(FooBarBaz, Foo, FooBarBaz) // [foo bar baz] || [foo] => [foo bar baz]
		test(Foo, FooBarBaz, FooBarBaz) // [foo] || [foo bar baz] => [foo bar baz]

		test(FooBar, FooBarBaz, FooBarBaz)    // [foo bar] || [foo bar baz] => [foo bar baz]
		test(FooBarBaz, FooBar, FooBarBaz)    // [foo bar baz] || [foo bar] => [foo bar baz]
		test(FooBarBaz, FooBarBaz, FooBarBaz) // [foo bar baz] || [foo bar baz] => [foo bar baz]
		test(FooBar, FooBar, FooBar)          // [foo bar] || [foo bar] => [foo bar]
		test(Foo, Foo, Foo)                   // [foo] || [foo] => [foo]
	})

	// integer
	tt.Run("int", func(t *testing.T) {
		var Nil []int
		Empty := []int{}
		Foo := []int{123}
		Bar := []int{456}
		FooBar := []int{123, 456}
		BarFoo := []int{456, 123}
		FooBarBaz := []int{123, 456, 789}

		test := func(a, b []int, e1 []int) {
			t.Helper()
			if a1 := SliceOr(a, b); !sliceEqual(a1, e1) {
				t.Errorf("SliceOr(`%#v`,`%#v`)=`%#v`, expected `%#v`", a, b, a1, e1)
			}
		}

		test(Nil, Nil, Nil)     // nil || nil => nil
		test(Nil, Empty, Nil)   // nil || [] => nil
		test(Empty, Nil, Nil)   // [] || nil => nil
		test(Empty, Empty, Nil) // [] || [] => nil
		test(Nil, Foo, Foo)     // nil || [foo] => [foo]
		test(Empty, Foo, Foo)   // [] && [foo] => [foo]
		test(Foo, Nil, Foo)     // [foo] && nil => [foo]
		test(Foo, Empty, Foo)   // [foo] && [] => [foo]

		test(Foo, Bar, FooBar) // [foo] || [bar] => [foo bar]
		test(Bar, Foo, BarFoo) // [foo] || [bar] => [bar foo]

		test(FooBar, Foo, FooBar)       // [foo bar] || [foo] => [foo bar]
		test(Foo, FooBar, FooBar)       // [foo] || [foo bar] => [foo bar]
		test(FooBarBaz, Foo, FooBarBaz) // [foo bar baz] || [foo] => [foo bar baz]
		test(Foo, FooBarBaz, FooBarBaz) // [foo] || [foo bar baz] => [foo bar baz]

		test(FooBar, FooBarBaz, FooBarBaz)    // [foo bar] || [foo bar baz] => [foo bar baz]
		test(FooBarBaz, FooBar, FooBarBaz)    // [foo bar baz] || [foo bar] => [foo bar baz]
		test(FooBarBaz, FooBarBaz, FooBarBaz) // [foo bar baz] || [foo bar baz] => [foo bar baz]
		test(FooBar, FooBar, FooBar)          // [foo bar] || [foo bar] => [foo bar]
		test(Foo, Foo, Foo)                   // [foo] || [foo] => [foo]
	})
}

// TestSliceDiff unit tests for `SliceDiff` function.
func TestSliceDiff(tt *testing.T) {
	// string
	tt.Run("str", func(t *testing.T) {
		var Nil []string
		Empty := []string{}
		Foo := []string{"foo"}
		Bar := []string{"bar"}
		Baz := []string{"baz"}
		FooBar := []string{"foo", "bar"}
		BarBaz := []string{"bar", "baz"}
		FooBarBaz := []string{"foo", "bar", "baz"}

		test := func(a, b []string, e1, e2 []string) {
			t.Helper()
			if a1, a2 := SliceDiff(a, b); !sliceEqual(a1, e1) || !sliceEqual(a2, e2) {
				t.Errorf("SliceDiff(`%#v`,`%#v`)=(`%#v`,`%#v`), expected (`%#v`,`%#v`)", a, b, a1, a2, e1, e2)
			}
		}

		test(Nil, Nil, Nil, Nil)         // diff(nil, nil) => nil, nil
		test(Nil, Empty, Nil, Empty)     // diff(nil, [])  => nil, []
		test(Empty, Nil, Empty, Nil)     // diff([], nil)  => [], nil
		test(Empty, Empty, Empty, Empty) // diff([], [])   => [], []

		test(Nil, Foo, Nil, Foo)     // diff(nil, [foo]) => nil, [foo]
		test(Empty, Foo, Empty, Foo) // diff([], [foo])  => [], [foo]
		test(Foo, Nil, Foo, Nil)     // diff([foo], nil) => [foo], nil
		test(Foo, Empty, Foo, Empty) // diff([foo], [])  => [foo], []

		test(Bar, Foo, Bar, Foo) // diff([bar], [foo]) => [bar], [foo]
		test(Foo, Bar, Foo, Bar) // diff([foo], [bar]) => [foo], [bar]
		test(Baz, Foo, Baz, Foo) // diff([baz], [foo]) => [baz], [foo]
		test(Foo, Baz, Foo, Baz) // diff([foo], [baz]) => [foo], [baz]

		test(FooBar, Foo, Bar, Nil)       // diff([foo bar], [foo]) => [bar], nil
		test(Foo, FooBar, Nil, Bar)       // diff([foo], [foo bar]) => nil, [bar]
		test(FooBarBaz, Foo, BarBaz, Nil) // diff([foo bar baz], [foo]) => [bar baz], nil
		test(Foo, FooBarBaz, Nil, BarBaz) // diff([foo], [foo bar baz]) => nil, [bar baz]

		test(FooBar, BarBaz, Foo, Baz)    // diff([foo bar], [bar baz]) => [foo], [baz]
		test(BarBaz, FooBar, Baz, Foo)    // diff([bar baz], [foo bar]) => [baz], [foo]
		test(FooBarBaz, FooBar, Baz, Nil) // diff([foo bar baz], [foo bar]) => [baz], nil
		test(FooBar, FooBarBaz, Nil, Baz) // diff([foo bar], [foo bar baz]) => nil, [baz]

		test(FooBarBaz, FooBarBaz, Nil, Nil) // diff([foo bar baz], [foo bar baz]) => nil, nil
		test(FooBar, FooBar, Nil, Nil)       // diff([foo bar], [foo bar]) => nil, nil
		test(BarBaz, BarBaz, Nil, Nil)       // diff([bar baz], [bar baz]) => nil, nil
		test(Foo, Foo, Nil, Nil)             // diff([foo], [foo]) => nil, nil
	})

	// integer
	tt.Run("int", func(t *testing.T) {
		var Nil []int
		Empty := []int{}
		Foo := []int{123}
		Bar := []int{456}
		Baz := []int{789}
		FooBar := []int{123, 456}
		BarBaz := []int{456, 789}
		FooBarBaz := []int{123, 456, 789}

		test := func(a, b []int, e1, e2 []int) {
			t.Helper()
			if a1, a2 := SliceDiff(a, b); !sliceEqual(a1, e1) || !sliceEqual(a2, e2) {
				t.Errorf("SliceDiff(`%#v`,`%#v`)=(`%#v`,`%#v`), expected (`%#v`,`%#v`)", a, b, a1, a2, e1, e2)
			}
		}

		test(Nil, Nil, Nil, Nil)         // diff(nil, nil) => nil, nil
		test(Nil, Empty, Nil, Empty)     // diff(nil, [])  => nil, []
		test(Empty, Nil, Empty, Nil)     // diff([], nil)  => [], nil
		test(Empty, Empty, Empty, Empty) // diff([], [])   => [], []

		test(Nil, Foo, Nil, Foo)     // diff(nil, [foo]) => nil, [foo]
		test(Empty, Foo, Empty, Foo) // diff([], [foo])  => [], [foo]
		test(Foo, Nil, Foo, Nil)     // diff([foo], nil) => [foo], nil
		test(Foo, Empty, Foo, Empty) // diff([foo], [])  => [foo], []

		test(Bar, Foo, Bar, Foo) // diff([bar], [foo]) => [bar], [foo]
		test(Foo, Bar, Foo, Bar) // diff([foo], [bar]) => [foo], [bar]
		test(Baz, Foo, Baz, Foo) // diff([baz], [foo]) => [baz], [foo]
		test(Foo, Baz, Foo, Baz) // diff([foo], [baz]) => [foo], [baz]

		test(FooBar, Foo, Bar, Nil)       // diff([foo bar], [foo]) => [bar], nil
		test(Foo, FooBar, Nil, Bar)       // diff([foo], [foo bar]) => nil, [bar]
		test(FooBarBaz, Foo, BarBaz, Nil) // diff([foo bar baz], [foo]) => [bar baz], nil
		test(Foo, FooBarBaz, Nil, BarBaz) // diff([foo], [foo bar baz]) => nil, [bar baz]

		test(FooBar, BarBaz, Foo, Baz)    // diff([foo bar], [bar baz]) => [foo], [baz]
		test(BarBaz, FooBar, Baz, Foo)    // diff([bar baz], [foo bar]) => [baz], [foo]
		test(FooBarBaz, FooBar, Baz, Nil) // diff([foo bar baz], [foo bar]) => [baz], nil
		test(FooBar, FooBarBaz, Nil, Baz) // diff([foo bar], [foo bar baz]) => nil, [baz]

		test(FooBarBaz, FooBarBaz, Nil, Nil) // diff([foo bar baz], [foo bar baz]) => nil, nil
		test(FooBar, FooBar, Nil, Nil)       // diff([foo bar], [foo bar]) => nil, nil
		test(BarBaz, BarBaz, Nil, Nil)       // diff([bar baz], [bar baz]) => nil, nil
		test(Foo, Foo, Nil, Nil)             // diff([foo], [foo]) => nil, nil
	})
}

// TestSliceEqual unit tests for `sliceEqual` function.
func TestSliceEqual(tt *testing.T) {
	// string
	tt.Run("str", func(t *testing.T) {
		// nil == nil
		if e, a := true, sliceEqual(([]string)(nil), ([]string)(nil)); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// nil != []
		if e, a := false, sliceEqual(([]string)(nil), []string{}); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// [] != nil
		if e, a := false, sliceEqual([]string{}, ([]string)(nil)); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// [] == []
		if e, a := true, sliceEqual([]string{}, []string{}); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// [] != [foo]
		if e, a := false, sliceEqual([]string{}, []string{"foo"}); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// [foo] != []
		if e, a := false, sliceEqual([]string{"foo"}, []string{}); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// [foo] == [foo]
		if e, a := true, sliceEqual([]string{"foo"}, []string{"foo"}); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// [foo] != [bar]
		if e, a := false, sliceEqual([]string{"foo"}, []string{"bar"}); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})
}

// sliceEqual checks if two slices are equal.
func sliceEqual[S ~[]E, E comparable](a, b S) bool {
	if a == nil {
		return b == nil
	} else if b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
