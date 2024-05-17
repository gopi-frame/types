package types

import "strconv"

// Uint8 uint8
type Uint8 uint8

// Abs abs
func (i Uint8) Abs() Uint8 {
	return i
}

// Equals equals
func (i Uint8) Equals(v any) bool {
	switch value := v.(type) {
	case uint8:
		return value == uint8(i)
	case Uint8:
		return value == i
	default:
		return false
	}
}

// Less less
func (i Uint8) Less(v any) bool {
	switch value := v.(type) {
	case uint8:
		return uint8(i) < value
	case Uint8:
		return i < value
	default:
		return false
	}
}

func (i Uint8) String() string {
	return strconv.FormatUint(uint64(i), 10)
}

// Int to int
func (i Uint8) Int() Int {
	return Int(i)
}

// Int8 to int8
func (i Uint8) Int8() Int8 {
	return Int8(i)
}

// Int16 to int16
func (i Uint8) Int16() Int16 {
	return Int16(i)
}

// Int32 to int32
func (i Uint8) Int32() Int32 {
	return Int32(i)
}

// Int64 to int64
func (i Uint8) Int64() Int64 {
	return Int64(i)
}

// Uint to uint
func (i Uint8) Uint() Uint {
	return Uint(i)
}

// Uint16 to uint16
func (i Uint8) Uint16() Uint16 {
	return Uint16(i)
}

// Uint32 to uint32
func (i Uint8) Uint32() Uint32 {
	return Uint32(i)
}

// Uint64 to uint64
func (i Uint8) Uint64() Uint64 {
	return Uint64(i)
}

// Float32 to float32
func (i Uint8) Float32() Float32 {
	return Float32(i)
}

// Float64 to float64
func (i Uint8) Float64() Float64 {
	return Float64(i)
}
