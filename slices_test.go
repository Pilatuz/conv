package conv_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Pilatuz/conv"
)

func ExampleSlice() {
	a := "foo"
	b := "bar"
	c := "baz"
	// we can combine a few variables of the same type
	// without providing type of the slice, so instead of
	// []string{a, b, c} we write conv.Slice(a, b, c)
	fmt.Println(conv.Slice(a, b, c))
	// Output:
	// [foo bar baz]
}

func ExampleSliceNotNil() {
	var s []string
	s = conv.SliceNotNil(s)
	if s == nil {
		fmt.Println("<nil>")
	} else {
		fmt.Println("[]")
	}
	// Output:
	// []
}

// TestSlice unit tests for `Slice` function.
func TestSlice(tt *testing.T) {
	// string
	tt.Run("str", func(t *testing.T) {
		// Slice() gives nil
		if e, a := ([]string)(nil), conv.Slice[string](); !reflect.DeepEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// Slice([]E{}...) gives empty slice
		if e, a := []string{}, conv.Slice([]string{}...); !reflect.DeepEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// Slice(a, b, ...) gives slice [a, b, ...]
		if e, a := []string{"foo", "bar"}, conv.Slice("foo", "bar"); !reflect.DeepEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})

	// integer
	tt.Run("int", func(t *testing.T) {
		// Slice() gives nil
		if e, a := ([]int)(nil), conv.Slice[int](); !reflect.DeepEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// Slice([]E{}...) gives empty slice
		if e, a := []int{}, conv.Slice([]int{}...); !reflect.DeepEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// Slice(a, b, ...) gives slice [a, b, ...]
		if e, a := []int{123, 456}, conv.Slice(123, 456); !reflect.DeepEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})
}

// TestSliceNotNil trivial unit tests for `SliceNotNil` function.
func TestSliceNotNil(tt *testing.T) {
	// string
	tt.Run("str", func(t *testing.T) {
		// SliceNotNil(nil) gives empty slice
		if e, a := []string{}, conv.SliceNotNil(([]string)(nil)); !reflect.DeepEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// SliceNotNil([]) gives empty slice
		if e, a := []string{}, conv.SliceNotNil([]string{}); !reflect.DeepEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// SliceNotNil([a, b, ...]) gives slice [a, b, ...]
		if e, a := []string{"foo", "bar"}, conv.SliceNotNil(conv.Slice("foo", "bar")); !reflect.DeepEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})

	// integer
	tt.Run("int", func(t *testing.T) {
		// SliceNotNil(nil) gives empty slice
		if e, a := []int{}, conv.SliceNotNil(([]int)(nil)); !reflect.DeepEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// SliceNotNil([]) gives empty slice
		if e, a := []int{}, conv.SliceNotNil([]int{}); !reflect.DeepEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// SliceNotNil([a, b, ...]) gives slice [a, b, ...]
		if e, a := []int{123, 456}, conv.SliceNotNil(conv.Slice(123, 456)); !reflect.DeepEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})
}
