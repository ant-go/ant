package array

import (
	"reflect"
	"strconv"

	"github.com/spf13/cast"
)

// ---------------------------------------------------------------------------------------------------------------------
// int8 / int16 / int32 / int64 / int
// ---------------------------------------------------------------------------------------------------------------------

func InArrayInt8(a int8, array []int8) bool {
	for _, e := range array {
		if a == e {
			return true
		}
	}
	return false
}

func InArrayInt16(a int16, array []int16) bool {
	for _, e := range array {
		if a == e {
			return true
		}
	}
	return false
}

func InArrayInt32(a int32, array []int32) bool {
	for _, e := range array {
		if a == e {
			return true
		}
	}
	return false
}

func InArrayInt64(a int64, array []int64) bool {
	for _, e := range array {
		if a == e {
			return true
		}
	}
	return false
}

func InArrayInt(a int, array []int) bool {
	for _, e := range array {
		if a == e {
			return true
		}
	}
	return false
}

// ---------------------------------------------------------------------------------------------------------------------
// uint8 / uint16 / uint32 / uint64 / uint
// ---------------------------------------------------------------------------------------------------------------------

func InArrayUInt8(a uint8, array []uint8) bool {
	for _, e := range array {
		if a == e {
			return true
		}
	}
	return false
}

func InArrayUInt16(a uint16, array []uint16) bool {
	for _, e := range array {
		if a == e {
			return true
		}
	}
	return false
}

func InArrayUInt32(a uint32, array []uint32) bool {
	for _, e := range array {
		if a == e {
			return true
		}
	}
	return false
}

func InArrayUInt64(a uint64, array []uint64) bool {
	for _, e := range array {
		if a == e {
			return true
		}
	}
	return false
}

func InArrayUInt(a uint, array []uint) bool {
	for _, e := range array {
		if a == e {
			return true
		}
	}
	return false
}

// ---------------------------------------------------------------------------------------------------------------------
// string
// ---------------------------------------------------------------------------------------------------------------------

func InArrayString(a string, array []string) bool {
	for _, e := range array {
		if a == e {
			return true
		}
	}
	return false
}

// ---------------------------------------------------------------------------------------------------------------------
// interface{}
// ---------------------------------------------------------------------------------------------------------------------

func InArray(needle interface{}, haystack interface{}) bool {
	v := reflect.ValueOf(haystack)
	if (v.Kind() != reflect.Array) && (v.Kind() != reflect.Slice) && (v.Kind() != reflect.Map) {
		return false
	}

	s, err := cast.ToStringE(needle)
	if err != nil {
		return false
	}

	switch v.Kind() {
	case reflect.Array, reflect.Slice:
		if v.Len() <= 0 {
			return false
		}
		for i := 0; i < v.Len(); i++ {
			switch v.Index(i).Kind() {
			case reflect.Interface:
				if s == cast.ToString(v.Index(i).Interface()) {
					return true
				}
			case reflect.String:
				if s == v.Index(i).String() {
					return true
				}
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				if s == strconv.FormatInt(v.Index(i).Int(), 10) {
					return true
				}
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
				if s == strconv.FormatUint(v.Index(i).Uint(), 10) {
					return true
				}
			}
		}
	case reflect.Map:
		if v.Len() <= 0 {
			return false
		}
		for _, key := range v.MapKeys() {
			switch v.MapIndex(key).Kind() {
			case reflect.String:
				if s == v.MapIndex(key).String() {
					return true
				}
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				if s == strconv.FormatInt(v.MapIndex(key).Int(), 10) {
					return true
				}
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
				if s == strconv.FormatUint(v.MapIndex(key).Uint(), 10) {
					return true
				}
			}
		}
	}
	return false
}

func InArrayStrict(needle interface{}, haystack interface{}) bool {
	v := reflect.ValueOf(haystack)
	switch v.Kind() {
	case reflect.Array, reflect.Slice:
		if v.Len() <= 0 {
			return false
		}
		for i := 0; i < v.Len(); i++ {
			if reflect.DeepEqual(needle, v.Index(i).Interface()) {
				return true
			}
		}
	case reflect.Map:
		if v.Len() <= 0 {
			return false
		}
		for _, key := range v.MapKeys() {
			if reflect.DeepEqual(needle, v.MapIndex(key).Interface()) {
				return true
			}
		}
	}
	return false
}
