package conv_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Pilatuz/conv"
)

func ExampleMapNotNil() {
	var m map[string]string
	m = conv.MapNotNil(m)
	if m == nil {
		fmt.Println("<nil>")
	} else {
		fmt.Println("{}")
	}
	// Output:
	// {}
}

func ExampleMapOmitEmpty() {
	m := conv.MapOmitEmpty(map[string]string{})
	if m == nil {
		fmt.Println("<nil>")
	} else {
		fmt.Println("{}")
	}
	// Output:
	// <nil>
}

// TestMapNotNil unit tests for the [MapNotNil] function.
func TestMapNotNil(tt *testing.T) {
	// string
	tt.Run("str", func(t *testing.T) {
		// MapNotNil(nil) gives empty map
		if e, a := map[string]string{}, conv.MapNotNil[map[string]string](nil); !reflect.DeepEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// MapNotNil({}) gives empty map
		if e, a := map[string]string{}, conv.MapNotNil(map[string]string{}); !reflect.DeepEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// MapNotNil({a, b, ...}) gives map {a, b, ...}
		if e, a := map[string]string{"foo": "1", "bar": "2"}, conv.MapNotNil(map[string]string{"foo": "1", "bar": "2"}); !reflect.DeepEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})

	// integer
	tt.Run("int", func(t *testing.T) {
		// MapNotNil(nil) gives empty map
		if e, a := map[string]int{}, conv.MapNotNil[map[string]int](nil); !reflect.DeepEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// MapNotNil({}) gives empty map
		if e, a := map[string]int{}, conv.MapNotNil(map[string]int{}); !reflect.DeepEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// MapNotNil({a, b, ...}) gives map {a, b, ...}
		if e, a := map[string]int{"foo": 123, "bar": 456}, conv.MapNotNil(map[string]int{"foo": 123, "bar": 456}); !reflect.DeepEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})

	// boolean
	tt.Run("bool", func(t *testing.T) {
		// MapNotNil(nil) gives empty map
		if e, a := map[string]bool{}, conv.MapNotNil[map[string]bool](nil); !reflect.DeepEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// MapNotNil({}) gives empty map
		if e, a := map[string]bool{}, conv.MapNotNil(map[string]bool{}); !reflect.DeepEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// MapNotNil({a, b, ...}) gives map {a, b, ...}
		if e, a := map[string]bool{"foo": true, "bar": false}, conv.MapNotNil(map[string]bool{"foo": true, "bar": false}); !reflect.DeepEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})
}

// TestMapOmitEmpty unit tests for the [MapOmitEmpty] function.
func TestMapOmitEmpty(tt *testing.T) {
	// string
	tt.Run("str", func(t *testing.T) {
		// MapOmitEmpty(nil) gives nil
		if e, a := (map[string]string)(nil), conv.MapOmitEmpty[map[string]string](nil); !reflect.DeepEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// MapOmitEmpty({}) gives nil
		if e, a := (map[string]string)(nil), conv.MapOmitEmpty(map[string]string{}); !reflect.DeepEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// MapOmitEmpty({a, b, ...}) gives map {a, b, ...}
		if e, a := map[string]string{"foo": "1", "bar": "2"}, conv.MapOmitEmpty(map[string]string{"foo": "1", "bar": "2"}); !reflect.DeepEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})

	// integer
	tt.Run("int", func(t *testing.T) {
		// MapOmitEmpty(nil) gives nil
		if e, a := (map[string]int)(nil), conv.MapOmitEmpty[map[string]int](nil); !reflect.DeepEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// MapOmitEmpty({}) gives nil
		if e, a := (map[string]int)(nil), conv.MapOmitEmpty(map[string]int{}); !reflect.DeepEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// MapOmitEmpty({a, b, ...}) gives map {a, b, ...}
		if e, a := map[string]int{"foo": 123, "bar": 456}, conv.MapOmitEmpty(map[string]int{"foo": 123, "bar": 456}); !reflect.DeepEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})

	// boolean
	tt.Run("bool", func(t *testing.T) {
		// MapOmitEmpty(nil) gives nil
		if e, a := (map[string]bool)(nil), conv.MapOmitEmpty[map[string]bool](nil); !reflect.DeepEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// MapOmitEmpty({}) gives nil
		if e, a := (map[string]bool)(nil), conv.MapOmitEmpty(map[string]bool{}); !reflect.DeepEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		// MapOmitEmpty({a, b, ...}) gives map {a, b, ...}
		if e, a := map[string]bool{"foo": true, "bar": false}, conv.MapOmitEmpty(map[string]bool{"foo": true, "bar": false}); !reflect.DeepEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})
}
