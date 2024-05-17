package types

import (
	"math"
	"strconv"
)

// Float32 float32
type Float32 float32

// Abs abs
func (f Float32) Abs() Float32 {
	return Float32(math.Abs(float64(f)))
}

// Equals equals
func (f Float32) Equals(v any) bool {
	switch value := v.(type) {
	case float32:
		return value == float32(f)
	case Float32:
		return value == f
	default:
		return false
	}
}

// Less less
func (f Float32) Less(v any) bool {
	switch value := v.(type) {
	case float32:
		return float32(f) < value
	case Float32:
		return f < value
	default:
		return false
	}
}

func (f Float32) String() string {
	return strconv.FormatFloat(float64(f), 'f', 2, 64)
}

// Int to int
func (f Float32) Int() Int {
	return Int(f)
}

// Int8 to int8
func (f Float32) Int8() Int8 {
	return Int8(f)
}

// Int16 to int16
func (f Float32) Int16() Int16 {
	return Int16(f)
}

// Int32 to int32
func (f Float32) Int32() Int32 {
	return Int32(f)
}

// Int64 to int64
func (f Float32) Int64() Int64 {
	return Int64(f)
}

// Uint to uint
func (f Float32) Uint() Uint {
	return Uint(f)
}

// Uint8 to uint8
func (f Float32) Uint8() Uint8 {
	return Uint8(f)
}

// Uint16 to uint16
func (f Float32) Uint16() Uint16 {
	return Uint16(f)
}

// Uint32 to uint32
func (f Float32) Uint32() Uint32 {
	return Uint32(f)
}

// Uint64 to uint64
func (f Float32) Uint64() Uint64 {
	return Uint64(f)
}

// Float64 to float64
func (f Float32) Float64() Float64 {
	return Float64(f)
}
