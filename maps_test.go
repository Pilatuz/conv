package conv_test

import (
	"testing"

	. "github.com/Pilatuz/conv"
)

// TestMapNotNil unit tests for `MapNotNil` function.
func TestMapNotNil(tt *testing.T) {
	// string
	tt.Run("str", func(t *testing.T) {
		// MapNotNil(nil) gives empty map
		if e, a := map[string]string{}, MapNotNil[map[string]string](nil); !mapEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// MapNotNil({}) gives empty map
		if e, a := map[string]string{}, MapNotNil(map[string]string{}); !mapEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// MapNotNil({a, b, ...}) gives map {a, b, ...}
		if e, a := map[string]string{"foo": "1", "bar": "2"}, MapNotNil(map[string]string{"foo": "1", "bar": "2"}); !mapEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})

	// integer
	tt.Run("int", func(t *testing.T) {
		// MapNotNil(nil) gives empty map
		if e, a := map[string]int{}, MapNotNil[map[string]int](nil); !mapEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// MapNotNil({}) gives empty map
		if e, a := map[string]int{}, MapNotNil(map[string]int{}); !mapEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// MapNotNil({a, b, ...}) gives map {a, b, ...}
		if e, a := map[string]int{"foo": 123, "bar": 456}, MapNotNil(map[string]int{"foo": 123, "bar": 456}); !mapEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})

	// boolean
	tt.Run("bool", func(t *testing.T) {
		// MapNotNil(nil) gives empty map
		if e, a := map[string]bool{}, MapNotNil[map[string]bool](nil); !mapEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// MapNotNil({}) gives empty map
		if e, a := map[string]bool{}, MapNotNil(map[string]bool{}); !mapEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// MapNotNil({a, b, ...}) gives map {a, b, ...}
		if e, a := map[string]bool{"foo": true, "bar": false}, MapNotNil(map[string]bool{"foo": true, "bar": false}); !mapEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})
}

// mapEqual checks if two maps are equal.
func mapEqual[M ~map[K]V, K comparable, V comparable](a, b M) bool {
	if a == nil {
		return b == nil
	} else if b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for k := range a {
		if a[k] != b[k] {
			return false
		}
	}

	return true
}
