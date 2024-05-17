package types

import (
	"strconv"
)

// Complex128 complex128
type Complex128 complex128

// Equals equals
func (c Complex128) Equals(v any) bool {
	switch value := v.(type) {
	case complex128:
		return complex128(c) == value
	case Complex128:
		return c == value
	default:
		return false
	}
}

func (c Complex128) String() string {
	return strconv.FormatComplex(complex128(c), 'f', 2, 128)
}
