package conv_test

import (
	"context"
	"fmt"
	"net"
	"testing"

	"github.com/Pilatuz/conv"
)

func ExampleErrorsAs() {
	err := &net.ParseError{Type: "typ", Text: "txt"}
	netErr, ok := conv.ErrorsAs[*net.ParseError](err)
	fmt.Println(netErr, ok)
	// Output:
	// invalid typ: txt true
}

// TestErrorsAs unit tests for `ErrorsAs` function.
func TestErrorsAs(tt *testing.T) {
	tt.Run("not_ok", func(t *testing.T) {
		err := context.Canceled
		if netErr, ok := conv.ErrorsAs[net.Error](err); ok || netErr != nil {
			tt.Errorf("expected no error conversion, found `%v`", netErr)
		}
	})

	tt.Run("ok", func(t *testing.T) {
		err := &net.ParseError{Type: "typ", Text: "txt"}
		if netErr, ok := conv.ErrorsAs[net.Error](err); !ok || netErr == nil {
			tt.Errorf("expected error converted")
		}
	})
}
