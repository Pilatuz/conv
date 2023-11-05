package slices_test

import (
	"testing"

	"github.com/Pilatuz/conv/slices"
)

// TestSetOr unit tests for `SetOr` function.
func TestSetOr(tt *testing.T) {
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
			if a1 := slices.SetOr(a, b); !equal(a1, e1) {
				t.Errorf("SetOr(`%#v`,`%#v`)=`%#v`, expected `%#v`", a, b, a1, e1)
			}
			if a1 := slices.SetOrBy(func(v string) string { return v }, a, b); !equal(a1, e1) {
				t.Errorf("SetOrBy(`%#v`,`%#v`)=`%#v`, expected `%#v`", a, b, a1, e1)
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
			if a1 := slices.SetOr(a, b); !equal(a1, e1) {
				t.Errorf("SetOr(`%#v`,`%#v`)=`%#v`, expected `%#v`", a, b, a1, e1)
			}
			if a1 := slices.SetOrBy(func(v int) int { return v }, a, b); !equal(a1, e1) {
				t.Errorf("SetOrBy(`%#v`,`%#v`)=`%#v`, expected `%#v`", a, b, a1, e1)
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
