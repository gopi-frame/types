package types

import "strconv"

// Uint uint
type Uint uint

// Abs abs
func (i Uint) Abs() Uint {
	return i
}

// Equals equals
func (i Uint) Equals(v any) bool {
	switch value := v.(type) {
	case uint:
		return value == uint(i)
	case Uint:
		return value == i
	default:
		return false
	}
}

// Less less
func (i Uint) Less(v any) bool {
	switch value := v.(type) {
	case uint:
		return uint(i) < value
	case Uint:
		return i < value
	default:
		return false
	}
}

func (i Uint) String() string {
	return strconv.FormatUint(uint64(i), 10)
}

// Int to int
func (i Uint) Int() Int {
	return Int(i)
}

// Int8 to int8
func (i Uint) Int8() Int8 {
	return Int8(i)
}

// Int16 to int16
func (i Uint) Int16() Int16 {
	return Int16(i)
}

// Int32 to int32
func (i Uint) Int32() Int32 {
	return Int32(i)
}

// Int64 to int64
func (i Uint) Int64() Int64 {
	return Int64(i)
}

// Uint8 to uint8
func (i Uint) Uint8() Uint8 {
	return Uint8(i)
}

// Uint16 to uint16
func (i Uint) Uint16() Uint16 {
	return Uint16(i)
}

// Uint32 to uint32
func (i Uint) Uint32() Uint32 {
	return Uint32(i)
}

// Uint64 to uint64
func (i Uint) Uint64() Uint64 {
	return Uint64(i)
}

// Float32 to float32
func (i Uint) Float32() Float32 {
	return Float32(i)
}

// Float64 to float64
func (i Uint) Float64() Float64 {
	return Float64(i)
}
