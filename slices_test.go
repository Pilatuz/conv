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
		if e, a := ([]string)(nil), Reverse(s1); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// Reverse([]) gives empty slice
		if e, a := []string{}, Reverse(s2); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// Reverse([a, b, ...]) gives slice [..., b, a]
		if e, a := []string{"foo"}, Reverse(s3); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
		if e, a := []string{"bar", "foo"}, Reverse(s4); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
		if e, a := []string{"baz", "bar", "foo"}, Reverse(s5); !sliceEqual(a, e) {
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
		var s1 []string
		s2 := []string{}
		s3 := []string{"foo"}
		s4 := []string{"foo", "bar"}
		s5 := []string{"foo", "bar", "baz"}

		// SliceAnd(nil|empty, nil|empty) gives nil|empty
		if e, a := ([]string)(nil), SliceAnd(s1, s1); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
		if e, a := []string{}, SliceAnd(s1, s2); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
		if e, a := ([]string)(nil), SliceAnd(s2, s1); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
		if e, a := []string{}, SliceAnd(s2, s2); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// SliceAnd(nil|empty, [foo]) gives [foo]
		if e, a := []string{"foo"}, SliceAnd(s1, s3); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
		if e, a := []string{"foo"}, SliceAnd(s2, s3); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
		if e, a := []string{"foo"}, SliceAnd(s3, s1); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
		if e, a := []string{"foo"}, SliceAnd(s3, s2); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// SliceAnd([foo bar], [foo]) gives [foo]
		if e, a := []string{"foo"}, SliceAnd(s4, s3); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
		if e, a := []string{"foo"}, SliceAnd(s3, s4); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// SliceAnd([foo bar baz], [foo]) gives [foo]
		if e, a := []string{"foo"}, SliceAnd(s5, s3); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
		if e, a := []string{"foo"}, SliceAnd(s3, s5); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// SliceAnd([foo bar baz], [foo bar]) gives [foo bar]
		if e, a := []string{"foo", "bar"}, SliceAnd(s4, s5); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
		if e, a := []string{"foo", "bar"}, SliceAnd(s5, s4); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// SliceAnd([foo bar baz], [foo bar baz]) gives [foo bar baz]
		if e, a := []string{"foo", "bar", "baz"}, SliceAnd(s5, s5); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
		if e, a := []string{"foo", "bar"}, SliceAnd(s4, s4); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
		if e, a := []string{"foo"}, SliceAnd(s3, s3); !sliceEqual(a, e) {
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

		// SliceAnd(nil|empty, nil|empty) gives nil|empty
		if e, a := ([]int)(nil), SliceAnd(s1, s1); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
		if e, a := []int{}, SliceAnd(s1, s2); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
		if e, a := ([]int)(nil), SliceAnd(s2, s1); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
		if e, a := []int{}, SliceAnd(s2, s2); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// SliceAnd(nil|empty, [123]) gives [123]
		if e, a := []int{123}, SliceAnd(s1, s3); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
		if e, a := []int{123}, SliceAnd(s2, s3); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
		if e, a := []int{123}, SliceAnd(s3, s1); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
		if e, a := []int{123}, SliceAnd(s3, s2); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// SliceAnd([123 456], [123]) gives [123]
		if e, a := []int{123}, SliceAnd(s4, s3); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
		if e, a := []int{123}, SliceAnd(s3, s4); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// SliceAnd([123 456 789], [123]) gives [123]
		if e, a := []int{123}, SliceAnd(s5, s3); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
		if e, a := []int{123}, SliceAnd(s3, s5); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// SliceAnd([123 456 789], [123 456]) gives [123 456]
		if e, a := []int{123, 456}, SliceAnd(s4, s5); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
		if e, a := []int{123, 456}, SliceAnd(s5, s4); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// SliceAnd([123 456 789], [123 456 789]) gives [123 456 789]
		if e, a := []int{123, 456, 789}, SliceAnd(s5, s5); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
		if e, a := []int{123, 456}, SliceAnd(s4, s4); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
		if e, a := []int{123}, SliceAnd(s3, s3); !sliceEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
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
