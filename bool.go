package types

import "strconv"

// Bool bool
type Bool bool

// Equals equals
func (b Bool) Equals(v any) bool {
	switch value := v.(type) {
	case Bool:
		return b == value
	case bool:
		return bool(b) == value
	default:
		return false
	}
}

func (b Bool) String() string {
	return strconv.FormatBool(bool(b))
}
