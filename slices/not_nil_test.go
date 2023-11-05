package slices_test

import (
	"fmt"
	"testing"

	"github.com/Pilatuz/conv/slices"
)

// ExampleNotNil an example for `NotNil` function.
func ExampleNotNil() {
	var s []string
	s = slices.NotNil(s)
	if s == nil {
		fmt.Println("<nil>")
	} else {
		fmt.Println("[]")
	}
	// Output:
	// []
}

// TestNotNil unit tests for `NotNil` function.
func TestNotNil(tt *testing.T) {
	// string
	tt.Run("str", func(t *testing.T) {
		var Nil []string
		Empty := []string{}
		Foo := []string{"foo"}
		FooBar := []string{"foo", "bar"}

		test := func(a, expected []string) {
			t.Helper()

			if actual := slices.NotNil(a); !equal(actual, expected) {
				t.Errorf("NotNil(`%#v`)=`%#v`, expected `%#v`", a, actual, expected)
			}
		}

		test(Nil, Empty)     // NotNil(nil) => []
		test(Empty, Empty)   // NotNil([]) => []
		test(Foo, Foo)       // NotNil([foo]) => [foo]
		test(FooBar, FooBar) // NotNil([foo bar]) => [foo bar]
	})

	// integer
	tt.Run("int", func(t *testing.T) {
		var Nil []int
		Empty := []int{}
		Foo := []int{123}
		FooBar := []int{123, 456}

		test := func(a, expected []int) {
			t.Helper()

			if actual := slices.NotNil(a); !equal(actual, expected) {
				t.Errorf("NotNil(`%#v`)=`%#v`, expected `%#v`", a, actual, expected)
			}
		}

		test(Nil, Empty)     // NotNil(nil) => []
		test(Empty, Empty)   // NotNil([]) => []
		test(Foo, Foo)       // NotNil([foo]) => [foo]
		test(FooBar, FooBar) // NotNil([foo bar]) => [foo bar]
	})

	// boolean
	tt.Run("bool", func(t *testing.T) {
		var Nil []bool
		Empty := []bool{}
		Foo := []bool{true}
		FooBar := []bool{true, false}

		test := func(a, expected []bool) {
			t.Helper()

			if actual := slices.NotNil(a); !equal(actual, expected) {
				t.Errorf("NotNil(`%#v`)=`%#v`, expected `%#v`", a, actual, expected)
			}
		}

		test(Nil, Empty)     // NotNil(nil) => []
		test(Empty, Empty)   // NotNil([]) => []
		test(Foo, Foo)       // NotNil([foo]) => [foo]
		test(FooBar, FooBar) // NotNil([foo bar]) => [foo bar]
	})
}
