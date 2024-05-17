package types

import "strconv"

// Uint16 uint16
type Uint16 uint16

// Abs abs
func (i Uint16) Abs() Uint16 {
	return i
}

// Equals equals
func (i Uint16) Equals(v any) bool {
	switch value := v.(type) {
	case uint16:
		return value == uint16(i)
	case Uint16:
		return value == i
	default:
		return false
	}
}

// Less less
func (i Uint16) Less(v any) bool {
	switch value := v.(type) {
	case uint16:
		return uint16(i) < value
	case Uint16:
		return i < value
	default:
		return false
	}
}

func (i Uint16) String() string {
	return strconv.FormatUint(uint64(i), 10)
}

// Int to int
func (i Uint16) Int() Int {
	return Int(i)
}

// Int8 to int8
func (i Uint16) Int8() Int8 {
	return Int8(i)
}

// Int16 to int16
func (i Uint16) Int16() Int16 {
	return Int16(i)
}

// Int32 to int32
func (i Uint16) Int32() Int32 {
	return Int32(i)
}

// Int64 to int64
func (i Uint16) Int64() Int64 {
	return Int64(i)
}

// Uint to uint
func (i Uint16) Uint() Uint {
	return Uint(i)
}

// Uint8 to uint8
func (i Uint16) Uint8() Uint8 {
	return Uint8(i)
}

// Uint32 to uint32
func (i Uint16) Uint32() Uint32 {
	return Uint32(i)
}

// Uint64 to uint64
func (i Uint16) Uint64() Uint64 {
	return Uint64(i)
}

// Float32 to float32
func (i Uint16) Float32() Float32 {
	return Float32(i)
}

// Float64 to float64
func (i Uint16) Float64() Float64 {
	return Float64(i)
}
