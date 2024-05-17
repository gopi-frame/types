package types

import (
	"math"
	"strconv"
)

// Int64 int64
type Int64 int64

// Abs abs
func (i Int64) Abs() Int64 {
	return Int64(math.Abs(float64(i)))
}

// Equals equals
func (i Int64) Equals(v any) bool {
	switch value := v.(type) {
	case int64:
		return value == int64(i)
	case Int64:
		return value == i
	default:
		return false
	}
}

// Less less
func (i Int64) Less(v any) bool {
	switch value := v.(type) {
	case int64:
		return int64(i) < value
	case Int64:
		return i < value
	default:
		return false
	}
}

func (i Int64) String() string {
	return strconv.FormatInt(int64(i), 10)
}

// Int to int
func (i Int64) Int() Int {
	return Int(i)
}

// Int8 to int8
func (i Int64) Int8() Int8 {
	return Int8(i)
}

// Int16 to int16
func (i Int64) Int16() Int16 {
	return Int16(i)
}

// Int32 to int32
func (i Int64) Int32() Int32 {
	return Int32(i)
}

// Uint to uint
func (i Int64) Uint() Uint {
	return Uint(i)
}

// Uint8 to uint8
func (i Int64) Uint8() Uint8 {
	return Uint8(i)
}

// Uint16 to uint16
func (i Int64) Uint16() Uint16 {
	return Uint16(i)
}

// Uint32 to uint32
func (i Int64) Uint32() Uint32 {
	return Uint32(i)
}

// Uint64 to uint64
func (i Int64) Uint64() Uint64 {
	return Uint64(i)
}

// Float32 to float32
func (i Int64) Float32() Float32 {
	return Float32(i)
}

// Float64 to float64
func (i Int64) Float64() Float64 {
	return Float64(i)
}
