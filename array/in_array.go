package array

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

func InArray(a interface{}, array ...interface{}) bool {
	for _, e := range array {
		if a == e {
			return true
		}
	}
	return false
}

func InArrayStrict(a interface{}, array ...interface{}) bool {
	for _, e := range array {
		if a == e {
			return true
		}
	}
	return false
}
