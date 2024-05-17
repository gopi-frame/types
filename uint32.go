package types

import "strconv"

// Uint32 uint32
type Uint32 uint32

// Abs abs
func (i Uint32) Abs() Uint32 {
	return i
}

// Equals equals
func (i Uint32) Equals(v any) bool {
	switch value := v.(type) {
	case uint32:
		return value == uint32(i)
	case Uint32:
		return value == i
	default:
		return false
	}
}

// Less less
func (i Uint32) Less(v any) bool {
	switch value := v.(type) {
	case uint32:
		return uint32(i) < value
	case Uint32:
		return i < value
	default:
		return false
	}
}

func (i Uint32) String() string {
	return strconv.FormatUint(uint64(i), 10)
}

// Int to int
func (i Uint32) Int() Int {
	return Int(i)
}

// Int8 to int8
func (i Uint32) Int8() Int8 {
	return Int8(i)
}

// Int16 to int16
func (i Uint32) Int16() Int16 {
	return Int16(i)
}

// Int32 to int32
func (i Uint32) Int32() Int32 {
	return Int32(i)
}

// Int64 to int64
func (i Uint32) Int64() Int64 {
	return Int64(i)
}

// Uint to uint
func (i Uint32) Uint() Uint {
	return Uint(i)
}

// Uint8 to uint8
func (i Uint32) Uint8() Uint8 {
	return Uint8(i)
}

// Uint16 to uint16
func (i Uint32) Uint16() Uint16 {
	return Uint16(i)
}

// Uint64 to uint64
func (i Uint32) Uint64() Uint64 {
	return Uint64(i)
}

// Float32 to float32
func (i Uint32) Float32() Float32 {
	return Float32(i)
}

// Float64 to float64
func (i Uint32) Float64() Float64 {
	return Float64(i)
}
