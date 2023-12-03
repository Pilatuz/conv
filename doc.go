// Package conv contains useful generic tools for trivial Go conversions.
//
// Usually this lib helps to hide some trivial "if" checks or loops.
// For example, if we need to dereference a pointer we use something like:
//
//	var v string
//	if p != nil {
//		v = *p
//	}
//
// but using [FromPtrOr] it could be much shorter:
//
//	v := conv.FromPtrOr(p, "")
//
// Also it's possible to use lazy-initializer function:
//
//	t := conv.FromPtrOrFunc(p, time.Now)
//
// Here t would have a time from pointer p or current time if p is nil.
// Important difference from [FromPtrOr] is that we don't call function
// until we need fallback value, so we can save resources if lazy
// initialization function is pretty heavy.
//
// Next thing is to get pointer to a value. Imagine the following API:
//
//	type API interface {
//		DoLogin(request *Login) error
//	}
//	type Login struct {
//		Username *string // optional
//		Type     *string // optional
//	}
//	const TypeInternal = "internal"
//
// to make a call to that API we need a few temporary variables:
//
//	tmpType := TypeInternal
//	tmpName := "admin"
//	api.DoLogin(&Login{
//		Username: &tmpName,
//		Type:     &tmpType,
//	})
//
// With help of [PtrFrom] we can rewrite it as:
//
//	api.DoLogin(&Login{
//		Username: conv.PtrFrom("admin"),
//		Type:     conv.PtrFrom(TypeInternal),
//	})
//
// There are also a few useful slice helpers.
// You can use [Slice] function to create slice from a set of values:
//
//	list := conv.Slice("foo", "bar")
//
// Here list is a slice of strings containing two elements: "foo" and "bar".
//
// Also take a look at useful [Coalesce] and [ErrorAs] functions.
package conv
