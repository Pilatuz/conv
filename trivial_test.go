package conv_test

import (
	"testing"

	"github.com/Pilatuz/conv"
)

// TestIntToInt unit tests for `IntToInt` function.
func TestIntToInt(tt *testing.T) {
	// int -> *
	tt.Run("int", func(t *testing.T) {
		const IN int = 123

		if e, a := int8(IN), conv.IntToInt[int8](IN); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
		if e, a := int16(IN), conv.IntToInt[int16](IN); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
		if e, a := int32(IN), conv.IntToInt[int32](IN); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
		if e, a := int64(IN), conv.IntToInt[int64](IN); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})

	// uint -> *
	tt.Run("uint", func(t *testing.T) {
		const IN uint = 123

		if e, a := uint8(IN), conv.IntToInt[uint8](IN); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
		if e, a := uint16(IN), conv.IntToInt[uint16](IN); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
		if e, a := uint32(IN), conv.IntToInt[uint32](IN); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
		if e, a := uint64(IN), conv.IntToInt[uint64](IN); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})

	// int8 -> *
	tt.Run("uint", func(t *testing.T) {
		const IN int8 = 123

		if e, a := int(IN), conv.IntToInt[int](IN); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
		if e, a := uint16(IN), conv.IntToInt[uint16](IN); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
		if e, a := int32(IN), conv.IntToInt[int32](IN); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
		if e, a := uint64(IN), conv.IntToInt[uint64](IN); a != e {
			t.Errorf("expected `%v`, found `%v`", e, a)
		}
	})
}
