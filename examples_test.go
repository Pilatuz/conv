package conv_test

import (
	"fmt"
	"net"

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

func ExamplePtrFrom() {
	const a = 12345
	const b = "foo"
	fmt.Println(*conv.PtrFrom(a))
	fmt.Println(*conv.PtrFrom(b))
	// Output:
	// 12345
	// foo
}

func ExampleFromPtrOr() {
	var p *int
	fmt.Println(conv.FromPtrOr(p, 12345))
	// Output:
	// 12345
}

func ExampleFromPtrOrFunc() {
	fn := func() int {
		// ... long init ...
		return 12345
	}
	var p *int
	fmt.Println(conv.FromPtrOrFunc(p, fn))
	// Output:
	// 12345
}

func ExampleOmitEmpty() {
	fmt.Println(conv.OmitEmpty(conv.PtrFrom("")))
	fmt.Println(conv.OmitEmpty(conv.PtrFrom(0)))
	// Output:
	// <nil>
	// <nil>
}

func ExampleFirstNonNil() {
	var a *string
	var b *string
	c := conv.PtrFrom("foo")
	var d *string

	fmt.Println(*conv.FirstNonNil(a, b, c, d))
	// Output:
	// foo
}

func ExampleCoalesce() {
	var a, c int
	b := 123

	fmt.Println(conv.Coalesce(0, a, b, c, 0))
	// Output:
	// 123
}

func ExampleErrorAs() {
	err := &net.ParseError{Type: "typ", Text: "txt"}
	netErr, ok := conv.ErrorsAs[*net.ParseError](err)
	fmt.Println(netErr, ok)
	// Output:
	// invalid typ: txt true
}
