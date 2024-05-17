package types

import (
	"strconv"
	"time"
)

// String string
type String string

// Len length
func (s String) Len() int {
	return len(s)
}

// Equals equals
func (s String) Equals(v any) bool {
	switch value := v.(type) {
	case string:
		return string(s) == value
	case String:
		return s == value
	default:
		return false
	}
}

// Less less
func (s String) Less(v any) bool {
	switch value := v.(type) {
	case string:
		return string(s) < value
	case String:
		return s < value
	default:
		return false
	}
}

func (s String) String() string {
	return string(s)
}

// Int to int
func (s String) Int() Int {
	if v, err := strconv.Atoi(string(s)); err != nil {
		panic(err)
	} else {
		return Int(v)
	}
}

// Int8 to int8
func (s String) Int8() Int8 {
	if v, err := strconv.ParseInt(string(s), 10, 8); err != nil {
		panic(err)
	} else {
		return Int8(v)
	}
}

// Int16 to int16
func (s String) Int16() Int16 {
	if v, err := strconv.ParseInt(string(s), 10, 16); err != nil {
		panic(err)
	} else {
		return Int16(v)
	}
}

// Int32 to int32
func (s String) Int32() Int32 {
	if v, err := strconv.ParseInt(string(s), 10, 32); err != nil {
		panic(err)
	} else {
		return Int32(v)
	}
}

// Int64 to int64
func (s String) Int64() Int64 {
	if v, err := strconv.ParseInt(string(s), 10, 64); err != nil {
		panic(err)
	} else {
		return Int64(v)
	}
}

// Uint to uint
func (s String) Uint() Uint {
	if v, err := strconv.ParseUint(string(s), 10, 64); err != nil {
		panic(err)
	} else {
		return Uint(v)
	}
}

// Uint8 to uint8
func (s String) Uint8() Uint8 {
	if v, err := strconv.ParseUint(string(s), 10, 8); err != nil {
		panic(err)
	} else {
		return Uint8(v)
	}
}

// Uint16 to uint16
func (s String) Uint16() Uint16 {
	if v, err := strconv.ParseUint(string(s), 10, 16); err != nil {
		panic(err)
	} else {
		return Uint16(v)
	}
}

// Uint32 to uint32
func (s String) Uint32() Uint32 {
	if v, err := strconv.ParseUint(string(s), 10, 32); err != nil {
		panic(err)
	} else {
		return Uint32(v)
	}
}

// Uint64 to uint64
func (s String) Uint64() Uint64 {
	if v, err := strconv.ParseUint(string(s), 10, 64); err != nil {
		panic(err)
	} else {
		return Uint64(v)
	}
}

// Float32 to float32
func (s String) Float32() Float32 {
	if v, err := strconv.ParseFloat(string(s), 32); err != nil {
		panic(err)
	} else {
		return Float32(v)
	}
}

// Float64 to float64
func (s String) Float64() Float64 {
	if v, err := strconv.ParseFloat(string(s), 64); err != nil {
		panic(err)
	} else {
		return Float64(v)
	}
}

// Bool to bool
func (s String) Bool() Bool {
	if v, err := strconv.ParseBool(string(s)); err != nil {
		panic(err)
	} else {
		return Bool(v)
	}
}

// Bytes to []byte
func (s String) Bytes() []byte {
	return []byte(s)
}

// Runes to []rune
func (s String) Runes() []rune {
	return []rune(s)
}

// Duration to duration
func (s String) Duration() time.Duration {
	if v, err := time.ParseDuration(string(s)); err != nil {
		panic(err)
	} else {
		return v
	}
}

// Time to time
func (s String) Time(layout string) time.Time {
	if v, err := time.Parse(layout, string(s)); err != nil {
		panic(err)
	} else {
		return v
	}
}
