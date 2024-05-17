package types

import (
	"math"
	"strconv"
)

// Int int
type Int int

// Abs abs
func (i Int) Abs() Int {
	return Int(math.Abs(float64(i)))
}

// Equals equals
func (i Int) Equals(v any) bool {
	switch value := v.(type) {
	case int:
		return value == int(i)
	case Int:
		return value == i
	default:
		return false
	}
}

// Less less
func (i Int) Less(v any) bool {
	switch value := v.(type) {
	case int:
		return int(i) < value
	case Int:
		return i < value
	default:
		return false
	}
}

func (i Int) String() string {
	return strconv.Itoa(int(i))
}

// Int8 to int8
func (i Int) Int8() Int8 {
	return Int8(i)
}

// Int16 to int16
func (i Int) Int16() Int16 {
	return Int16(i)
}

// Int32 to int32
func (i Int) Int32() Int32 {
	return Int32(i)
}

// Int64 to int64
func (i Int) Int64() Int64 {
	return Int64(i)
}

// Uint to uint
func (i Int) Uint() Uint {
	return Uint(i)
}

// Uint8 to uint8
func (i Int) Uint8() Uint8 {
	return Uint8(i)
}

// Uint16 to uint16
func (i Int) Uint16() Uint16 {
	return Uint16(i)
}

// Uint32 to uint32
func (i Int) Uint32() Uint32 {
	return Uint32(i)
}

// Uint64 to uint64
func (i Int) Uint64() Uint64 {
	return Uint64(i)
}

// Float32 to float32
func (i Int) Float32() Float32 {
	return Float32(i)
}

// Float64 to float64
func (i Int) Float64() Float64 {
	return Float64(i)
}
