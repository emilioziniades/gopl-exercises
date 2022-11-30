// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 339.

package sexpr

import (
	"bytes"
	"fmt"
	"reflect"
)

//!+Marshal
// Marshal encodes a Go value in S-expression form.
func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := prettyEncode(&buf, reflect.ValueOf(v), 0); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

//!-Marshal

// prettyEncode writes to buf an S-expression representation of v.
//!+prettyEncode
func prettyEncode(buf *bytes.Buffer, v reflect.Value, depth int) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("nil")

	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(buf, "%d", v.Uint())

	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())

	case reflect.Ptr:
		return prettyEncode(buf, v.Elem(), depth+1)

	case reflect.Array, reflect.Slice: // (value ...)
		buf.WriteByte('(')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.WriteByte(' ')
			}
			if err := prettyEncode(buf, v.Index(i), depth+1); err != nil {
				return err
			}
		}
		buf.WriteByte(')')

	case reflect.Struct: // ((name value) ...)
		buf.WriteString("(")
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				buf.WriteString("\n")
			}
			fmt.Fprintf(buf, "(%s ", v.Type().Field(i).Name)
			if err := prettyEncode(buf, v.Field(i), depth+1); err != nil {
				return err
			}
			buf.WriteString(")")
		}
		buf.WriteByte(')')

	case reflect.Map: // ((key value) ...)
		buf.WriteByte('(')
		for i, key := range v.MapKeys() {
			if i > 0 {
				buf.WriteString("\n")
			}
			fmt.Fprintf(buf, "%*s(", depth, "\t")
			if err := prettyEncode(buf, key, depth+1); err != nil {
				return err
			}
			buf.WriteByte(' ')
			if err := prettyEncode(buf, v.MapIndex(key), depth+1); err != nil {
				return err
			}
			buf.WriteByte(')')
		}
		buf.WriteByte(')')
	case reflect.Bool:
		if v.Bool() {
			buf.WriteByte('t')
		} else {
			buf.WriteString("nil")
		}
	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(buf, "%g", v.Float())
	case reflect.Complex128, reflect.Complex64:
		fmt.Fprintf(buf, "#C(%f %f)", real(v.Complex()), imag(v.Complex()))
	case reflect.Interface:
		fmt.Fprintf(buf, "(%s ", v.Elem().Type())
		if err := prettyEncode(buf, v.Elem(), depth+1); err != nil {
			return err
		}
		buf.WriteByte(')')

	default: // float, complex, chan, func, interface
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}

//!-encode
