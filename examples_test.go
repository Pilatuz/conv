package conv_test

import (
	"fmt"
	"net"
	"sync"

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

func ExampleSliceNotNil() {
	var s []string
	s = conv.SliceNotNil(s)
	if s == nil {
		fmt.Println("<nil>")
	} else {
		fmt.Println("[]")
	}
	// Output:
	// []
}

func ExampleSliceJoin() {
	s := conv.SliceJoin(
		conv.Slice("foo", "bar"),
		conv.Slice("baz"))
	fmt.Println(s)
	// Output:
	// [foo bar baz]
}

func ExampleSliceFromChan() {
	wg := &sync.WaitGroup{}
	errCh := make(chan error)

	wg.Add(3)
	go func() {
		defer wg.Done()
		errCh <- nil
	}()
	go func() {
		defer wg.Done()
		errCh <- nil
	}()
	go func() {
		defer wg.Done()
		errCh <- nil
	}()

	go func() {
		wg.Wait()
		close(errCh)
	}()

	res := conv.SliceFromChan(errCh)
	fmt.Println(res)
	// Output:
	// [<nil> <nil> <nil>]
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

func ExampleUnique() {
	s := conv.Slice(4, 1, 2, 1, 2, 3, 2, 1)
	fmt.Println(conv.Unique(s))
	// Output:
	// [4 1 2 3]
}

func ExampleReverse() {
	s := conv.Slice(1, 2, 3, 4)
	conv.Reverse(s)
	fmt.Println(s)
	// Output:
	// [4 3 2 1]
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

func ExampleErrorAs() {
	err := &net.ParseError{Type: "typ", Text: "txt"}
	netErr, ok := conv.ErrorsAs[*net.ParseError](err)
	fmt.Println(netErr, ok)
	// Output:
	// invalid typ: txt true
}
