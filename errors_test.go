package conv_test

import (
	"context"
	"fmt"
	"net"
	"testing"

	"github.com/Pilatuz/conv"
)

func ExampleErrorAs() {
	err := &net.ParseError{Type: "typ", Text: "txt"}
	netErr, ok := conv.ErrorAs[*net.ParseError](err)
	fmt.Println(netErr, ok)
	// Output:
	// invalid typ: txt true
}

// TestErrorAs unit tests for [ErrorAs] function.
func TestErrorAs(tt *testing.T) {
	tt.Run("not_ok", func(t *testing.T) {
		err := context.Canceled
		if netErr, ok := conv.ErrorAs[net.Error](err); ok || netErr != nil {
			t.Errorf("expected no error conversion, found `%v`", netErr)
		}
	})

	tt.Run("ok", func(t *testing.T) {
		err := &net.ParseError{Type: "typ", Text: "txt"}
		if netErr, ok := conv.ErrorAs[net.Error](err); !ok || netErr == nil {
			t.Errorf("expected error converted")
		}
	})
}

// TestErrorIsAnyOf unit tets for [ErrorIsAnyOf] function.
func TestErrorIsAnyOf(tt *testing.T) {
	tt.Run("empty", func(t *testing.T) {
		if conv.ErrorIsAnyOf(context.Canceled) {
			t.Errorf("expected no match for empty targets")
		}
		if conv.ErrorIsAnyOf(fmt.Errorf("foo")) {
			t.Errorf("expected no match for empty targets")
		}
	})

	tt.Run("match", func(t *testing.T) {
		if conv.ErrorIsAnyOf(fmt.Errorf("foo"), context.Canceled) {
			t.Errorf("expected no match")
		}
		if !conv.ErrorIsAnyOf(fmt.Errorf("foo %w", context.Canceled), context.DeadlineExceeded, context.Canceled) {
			t.Errorf("expected match")
		}
	})
}
