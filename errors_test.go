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
