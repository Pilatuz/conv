package conv_test

import (
	"reflect"
	"testing"

	. "github.com/Pilatuz/conv"
)

// TestSlice unit tests for `Slice` function.
func TestSlice(tt *testing.T) {
	// string
	tt.Run("str", func(t *testing.T) {
		// Slice() gives nil
		if e, a := ([]string)(nil), Slice[string](); !reflect.DeepEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// Slice([]E{}...) gives empty slice
		if e, a := []string{}, Slice([]string{}...); !reflect.DeepEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// Slice(a, b, ...) gives slice [a, b, ...]
		if e, a := []string{"foo", "bar"}, Slice("foo", "bar"); !reflect.DeepEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})

	// integer
	tt.Run("int", func(t *testing.T) {
		// Slice() gives nil
		if e, a := ([]int)(nil), Slice[int](); !reflect.DeepEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// Slice([]E{}...) gives empty slice
		if e, a := []int{}, Slice([]int{}...); !reflect.DeepEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// Slice(a, b, ...) gives slice [a, b, ...]
		if e, a := []int{123, 456}, Slice(123, 456); !reflect.DeepEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})

	// boolean
	tt.Run("bool", func(t *testing.T) {
		// Slice() gives nil
		if e, a := ([]bool)(nil), Slice[bool](); !reflect.DeepEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// Slice([]E{}...) gives empty slice
		if e, a := []bool{}, Slice([]bool{}...); !reflect.DeepEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// Slice(a, b, ...) gives slice [a, b, ...]
		if e, a := []bool{true, false}, Slice(true, false); !reflect.DeepEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})
}

// TestSliceNotNil trivial unit tests for `SliceNotNil` function.
func TestSliceNotNil(tt *testing.T) {
	// string
	tt.Run("str", func(t *testing.T) {
		// SliceNotNil(nil) gives empty slice
		if e, a := []string{}, SliceNotNil(([]string)(nil)); !reflect.DeepEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// SliceNotNil([]) gives empty slice
		if e, a := []string{}, SliceNotNil([]string{}); !reflect.DeepEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// SliceNotNil([a, b, ...]) gives slice [a, b, ...]
		if e, a := []string{"foo", "bar"}, SliceNotNil(Slice("foo", "bar")); !reflect.DeepEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})

	// integer
	tt.Run("int", func(t *testing.T) {
		// SliceNotNil(nil) gives empty slice
		if e, a := []int{}, SliceNotNil(([]int)(nil)); !reflect.DeepEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// SliceNotNil([]) gives empty slice
		if e, a := []int{}, SliceNotNil([]int{}); !reflect.DeepEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// SliceNotNil([a, b, ...]) gives slice [a, b, ...]
		if e, a := []int{123, 456}, SliceNotNil(Slice(123, 456)); !reflect.DeepEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})
}
