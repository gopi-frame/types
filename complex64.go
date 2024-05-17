package types

import (
	"strconv"
)

// Complex64 complex64
type Complex64 complex64

// Equals equals
func (c Complex64) Equals(v any) bool {
	switch value := v.(type) {
	case complex64:
		return complex64(c) == value
	case Complex64:
		return c == value
	default:
		return false
	}
}

func (c Complex64) String() string {
	return strconv.FormatComplex(complex128(c), 'f', 2, 128)
}
