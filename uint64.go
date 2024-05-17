package types

import "strconv"

// Uint64 uint64
type Uint64 uint64

// Abs abs
func (i Uint64) Abs() Uint64 {
	return i
}

// Equals equals
func (i Uint64) Equals(v any) bool {
	switch value := v.(type) {
	case uint64:
		return value == uint64(i)
	case Uint64:
		return value == i
	default:
		return false
	}
}

// Less less
func (i Uint64) Less(v any) bool {
	switch value := v.(type) {
	case uint64:
		return uint64(i) < value
	case Uint64:
		return i < value
	default:
		return false
	}
}

func (i Uint64) String() string {
	return strconv.FormatUint(uint64(i), 10)
}

// Int to int
func (i Uint64) Int() Int {
	return Int(i)
}

// Int8 to int8
func (i Uint64) Int8() Int8 {
	return Int8(i)
}

// Int16 to int16
func (i Uint64) Int16() Int16 {
	return Int16(i)
}

// Int32 to int32
func (i Uint64) Int32() Int32 {
	return Int32(i)
}

// Int64 to int64
func (i Uint64) Int64() Int64 {
	return Int64(i)
}

// Uint to uint
func (i Uint64) Uint() Uint {
	return Uint(i)
}

// Uint8 to uint8
func (i Uint64) Uint8() Uint8 {
	return Uint8(i)
}

// Uint16 to uint16
func (i Uint64) Uint16() Uint16 {
	return Uint16(i)
}

// Uint32 to uint32
func (i Uint64) Uint32() Uint32 {
	return Uint32(i)
}

// Float32 to float32
func (i Uint64) Float32() Float32 {
	return Float32(i)
}

// Float64 to float64
func (i Uint64) Float64() Float64 {
	return Float64(i)
}
