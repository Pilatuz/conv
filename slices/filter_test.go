package slices_test

import (
	"fmt"
	"testing"
	"unicode/utf8"

	"github.com/Pilatuz/conv/slices"
)

// ExampleFilter an example for `Filter` function.
func ExampleFilter() {
	ss := []string{"foo", "barbaz"}
	cond := func(s string) bool {
		return utf8.RuneCountInString(s) > 3
	}
	fmt.Println(slices.Filter(cond, ss))
	// Output:
	// [barbaz]
}

// TestFilter unit tests for `Filter` function
func TestFilter(tt *testing.T) {
	// string
	tt.Run("str", func(t *testing.T) {
		var Nil []string
		Empty := []string{}
		Foo := []string{"foo"}
		FooBar := []string{"foo", "bar"}
		BarFoo := []string{"bar", "foo"}
		FooBarBaz := []string{"foo", "bar", "baz"}

		test := func(input, expected []string) {
			t.Helper()

			isFoo := func(s string) bool {
				return s == "foo"
			}

			if actual := slices.Filter(isFoo, input); !equal(actual, expected) || same(actual, input) {
				t.Errorf("Filter(`%#v`)=`%#v`, expected `%#v`", input, actual, expected)
			}

			input2 := slices.Clone(input)
			if actual := slices.FilterInPlace(isFoo, input2); !equal(actual, expected) {
				t.Errorf("FilterInPlace(`%#v`)=`%#v`, expected `%#v`", input, actual, expected)
			} else if len(actual) > 0 && !same(actual, input2) {
				t.Errorf("FilterInPlace(`%#v`)=`%#v`, expected `%#v`", input, actual, expected)
			}
		}

		test(Nil, Nil)       // Filter(nil) => nil
		test(Empty, Empty)   // Filter([]) => []
		test(Foo, Foo)       // Filter([foo]) => [foo]
		test(FooBar, Foo)    // Filter([foo bar]) => [foo]
		test(BarFoo, Foo)    // Filter([bar foo]) => [foo]
		test(FooBarBaz, Foo) // Filter([foo bar baz]) => [foo]
	})

	// integer
	tt.Run("int", func(t *testing.T) {
		var Nil []int
		Empty := []int{}
		Foo := []int{123}
		FooBar := []int{123, 456}
		BarFoo := []int{456, 123}
		FooBarBaz := []int{123, 456, 789}

		test := func(input, expected []int) {
			t.Helper()

			isFoo := func(s int) bool {
				return s == 123
			}

			if actual := slices.Filter(isFoo, input); !equal(actual, expected) || same(actual, input) {
				t.Errorf("Filter(`%#v`)=`%#v`, expected `%#v`", input, actual, expected)
			}

			input2 := slices.Clone(input)
			if actual := slices.FilterInPlace(isFoo, input2); !equal(actual, expected) {
				t.Errorf("FilterInPlace(`%#v`)=`%#v`, expected `%#v`", input, actual, expected)
			} else if len(actual) > 0 && !same(actual, input2) {
				t.Errorf("FilterInPlace(`%#v`)=`%#v`, expected `%#v`", input, actual, expected)
			}
		}

		test(Nil, Nil)       // Filter(nil) => nil
		test(Empty, Empty)   // Filter([]) => []
		test(Foo, Foo)       // Filter([foo]) => [foo]
		test(FooBar, Foo)    // Filter([foo bar]) => [foo]
		test(BarFoo, Foo)    // Filter([bar foo]) => [foo]
		test(FooBarBaz, Foo) // Filter([foo bar baz]) => [foo]
	})
}
