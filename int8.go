package types

import (
	"math"
	"strconv"
)

// Int8 int8
type Int8 int8

// Abs abs
func (i Int8) Abs() Int8 {
	return Int8(math.Abs(float64(i)))
}

// Equals equals
func (i Int8) Equals(v any) bool {
	switch value := v.(type) {
	case int8:
		return value == int8(i)
	case Int8:
		return value == i
	default:
		return false
	}
}

// Less less
func (i Int8) Less(v any) bool {
	switch value := v.(type) {
	case int8:
		return int8(i) < value
	case Int8:
		return i < value
	default:
		return false
	}
}

func (i Int8) String() string {
	return strconv.FormatInt(int64(i), 10)
}

// Int to int
func (i Int8) Int() Int {
	return Int(i)
}

// Int16 to int16
func (i Int8) Int16() Int16 {
	return Int16(i)
}

// Int32 to int32
func (i Int8) Int32() Int32 {
	return Int32(i)
}

// Int64 to int64
func (i Int8) Int64() Int64 {
	return Int64(i)
}

// Uint to uint
func (i Int8) Uint() Uint {
	return Uint(i)
}

// Uint8 to uint8
func (i Int8) Uint8() Uint8 {
	return Uint8(i)
}

// Uint16 to uint16
func (i Int8) Uint16() Uint16 {
	return Uint16(i)
}

// Uint32 to uint32
func (i Int8) Uint32() Uint32 {
	return Uint32(i)
}

// Uint64 to uint64
func (i Int8) Uint64() Uint64 {
	return Uint64(i)
}

// Float32 to float32
func (i Int8) Float32() Float32 {
	return Float32(i)
}

// Float64 to float64
func (i Int8) Float64() Float64 {
	return Float64(i)
}
