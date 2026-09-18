//go:build go1.23

package conv_test

import (
	"reflect"
	"slices"
	"testing"

	"github.com/Pilatuz/conv"
)

// TestAllNotNil unit tests for the [AllNotNil] function.
func TestAllNotNil(tt *testing.T) {
	// string
	tt.Run("str", func(t *testing.T) {
		var p1 *string
		var p2 string
		p3 := "foo"

		if e, a := ([]*string)(nil), slices.Collect(conv.AllNotNil(p1, nil, p1)); !reflect.DeepEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		if e, a := []*string{&p2, &p3}, slices.Collect(conv.AllNotNil(p1, &p2, &p3, p1)); !reflect.DeepEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		if e, a := []*string{&p3, &p2}, slices.Collect(conv.AllNotNil(p1, &p3, &p2, p1)); !reflect.DeepEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		var a []*string
		for p := range conv.AllNotNil(nil, &p2, &p3) {
			a = append(a, p)
			break // on first iteration
		}
		if e := []*string{&p2}; !reflect.DeepEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})

	// integer
	tt.Run("int", func(t *testing.T) {
		var p1 *int
		var p2 int
		p3 := 123

		if e, a := ([]*int)(nil), slices.Collect(conv.AllNotNil(p1, nil, p1)); !reflect.DeepEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		if e, a := []*int{&p2, &p3}, slices.Collect(conv.AllNotNil(p1, &p2, &p3, p1)); !reflect.DeepEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		if e, a := []*int{&p3, &p2}, slices.Collect(conv.AllNotNil(p1, &p3, &p2, p1)); !reflect.DeepEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}

		var a []*int
		for p := range conv.AllNotNil(nil, &p2, &p3) {
			a = append(a, p)
			break // on first iteration
		}
		if e := []*int{&p2}; !reflect.DeepEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})
}

// TestAllNotNil2 unit tests for the [AllNotNil2] function.
func TestAllNotNil2(tt *testing.T) {
	// string
	tt.Run("str", func(t *testing.T) {
		var p1 *string
		var p2 string
		p3 := "foo"

		a := make(map[int]*string)
		for i, p := range conv.AllNotNil2(nil, p1, &p2, &p3) {
			a[i] = p
			break // on first iteration
		}
		if e := map[int]*string{2: &p2}; !reflect.DeepEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})

	// integer
	tt.Run("int", func(t *testing.T) {
		var p1 *int
		var p2 int
		p3 := 123

		a := make(map[int]*int)
		for i, p := range conv.AllNotNil2(nil, p1, &p2, &p3) {
			a[i] = p
			break // on first iteration
		}
		if e := map[int]*int{2: &p2}; !reflect.DeepEqual(a, e) {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})
}
