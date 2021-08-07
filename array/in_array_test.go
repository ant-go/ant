package array

import "testing"

func TestInArray(t *testing.T) {
	type args struct {
		needle   interface{}
		haystack interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// true-slice
		{"in_array_int_[]int_true", args{1, []int{0, 1, 2, 3}}, true},
		{"in_array_int_[]string_true", args{1, []string{"0", "1", "2", "3"}}, true},
		{"in_array_string_[]int_true", args{"1", []int{0, 1, 2, 3}}, true},
		{"in_array_string_[]string_true", args{"1", []string{"0", "1", "2", "3"}}, true},
		{"in_array_interface_int_[]int_true", args{interface{}(1), []int{0, 1, 2, 3}}, true},
		{"in_array_interface_int_[]string_true", args{interface{}(1), []string{"0", "1", "2", "3"}}, true},
		{"in_array_interface_int_[]interface_int_true", args{interface{}(1), []interface{}{"0", 1, "2", "3"}}, true},
		{"in_array_interface_int_[]interface_string_true", args{interface{}(1), []interface{}{"0", "1", "2", "3"}}, true},
		{"in_array_interface_string_[]int_true", args{interface{}("1"), []int{0, 1, 2, 3}}, true},
		{"in_array_interface_string_[]string_true", args{interface{}("1"), []string{"0", "1", "2", "3"}}, true},
		{"in_array_interface_string_[]interface_int_true", args{interface{}("1"), []interface{}{"0", 1, "2", "3"}}, true},
		{"in_array_interface_string_[]interface_string_true", args{interface{}("1"), []interface{}{"0", "1", "2", "3"}}, true},

		{"in_array_uint_[]uint_true", args{1, []uint{0, 1, 2, 3}}, true},
		{"in_array_uint_[]string_true", args{1, []string{"0", "1", "2", "3"}}, true},
		{"in_array_string_[]uint_true", args{"1", []uint{0, 1, 2, 3}}, true},
		{"in_array_interface_uint_[]uint_true", args{interface{}(uint(1)), []uint{0, 1, 2, 3}}, true},
		{"in_array_interface_uint_[]string_true", args{interface{}(uint(1)), []string{"0", "1", "2", "3"}}, true},
		{"in_array_interface_uint_[]interface_uint_true", args{interface{}(uint(1)), []interface{}{"0", uint(1), "2", "3"}}, true},
		{"in_array_interface_uint_[]interface_string_true", args{interface{}(uint(1)), []interface{}{"0", "1", "2", "3"}}, true},
		{"in_array_interface_string_[]uint_true", args{interface{}("1"), []uint{0, 1, 2, 3}}, true},
		{"in_array_interface_string_[]interface_uint_true", args{interface{}("1"), []interface{}{"0", uint(1), "2", "3"}}, true},

		// false-slice
		{"in_array_int_[]int_false", args{4, []int{0, 1, 2, 3}}, false},
		{"in_array_int_[]string_false", args{4, []string{"0", "1", "2", "3"}}, false},
		{"in_array_string_[]int_false", args{"4", []int{0, 1, 2, 3}}, false},
		{"in_array_string_[]string_false", args{"4", []string{"0", "1", "2", "3"}}, false},
		{"in_array_interface_int_[]int_false", args{interface{}(4), []int{0, 1, 2, 3}}, false},
		{"in_array_interface_int_[]string_false", args{interface{}(4), []string{"0", "1", "2", "3"}}, false},
		{"in_array_interface_int_[]interface_int_false", args{interface{}(4), []interface{}{"0", 1, "2", "3"}}, false},
		{"in_array_interface_int_[]interface_string_false", args{interface{}(4), []interface{}{"0", "1", "2", "3"}}, false},
		{"in_array_interface_string_[]int_false", args{interface{}("4"), []int{0, 1, 2, 3}}, false},
		{"in_array_interface_string_[]string_false", args{interface{}("4"), []string{"0", "1", "2", "3"}}, false},
		{"in_array_interface_string_[]interface_int_false", args{interface{}("4"), []interface{}{"0", 1, "2", "3"}}, false},
		{"in_array_interface_string_[]interface_string_false", args{interface{}("4"), []interface{}{"0", "1", "2", "3"}}, false},

		{"in_array_uint_[]uint_false", args{4, []uint{0, 1, 2, 3}}, false},
		{"in_array_uint_[]string_false", args{4, []string{"0", "1", "2", "3"}}, false},
		{"in_array_string_[]uint_false", args{"4", []uint{0, 1, 2, 3}}, false},
		{"in_array_interface_uint_[]uint_false", args{interface{}(uint(4)), []uint{0, 1, 2, 3}}, false},
		{"in_array_interface_uint_[]string_false", args{interface{}(uint(4)), []string{"0", "1", "2", "3"}}, false},
		{"in_array_interface_uint_[]interface_uint_false", args{interface{}(uint(4)), []interface{}{"0", uint(1), "2", "3"}}, false},
		{"in_array_interface_uint_[]interface_string_false", args{interface{}(uint(4)), []interface{}{"0", "1", "2", "3"}}, false},
		{"in_array_interface_string_[]uint_false", args{interface{}("4"), []uint{0, 1, 2, 3}}, false},
		{"in_array_interface_string_[]interface_uint_false", args{interface{}("4"), []interface{}{"0", uint(1), "2", "3"}}, false},

		// true-map
		{"in_array_int_map[int]int_true", args{1, map[int]int{0:0, 1:1, 2:2, 3:3}}, true},
		{"in_array_int_map[int]string_true", args{1, map[int]string{0:"0", 1:"1", 2:"2", 3:"3"}}, true},
		{"in_array_string_map[int]int_true", args{"1", map[int]int{0:0, 1:1, 2:2, 3:3}}, true},
		{"in_array_string_map[int]string_true", args{"1", map[int]string{0:"0", 1:"1", 2:"2", 3:"3"}}, true},
		{"in_array_interface_int_map[int]int_true", args{interface{}(1), map[int]int{0:0, 1:1, 2:2, 3:3}}, true},
		{"in_array_interface_int_map[int]string_true", args{interface{}(1), map[int]string{0:"0", 1:"1", 2:"2", 3:"3"}}, true},
		{"in_array_interface_int_map[int]int_true", args{interface{}(1), map[int]int{0:0, 1:1, 2:2, 3:3}}, true},
		{"in_array_interface_int_map[int]string_true", args{interface{}(1), map[int]string{0:"0", 1:"1", 2:"2", 3:"3"}}, true},
		{"in_array_interface_string_map[int]int_true", args{interface{}("1"), map[int]int{0:0, 1:1, 2:2, 3:3}}, true},
		{"in_array_interface_string_map[int]string_true", args{interface{}("1"), map[int]string{0:"0", 1:"1", 2:"2", 3:"3"}}, true},
		{"in_array_interface_string_map[int]int_true", args{interface{}("1"), map[int]int{0:0, 1:1, 2:2, 3:3}}, true},
		{"in_array_interface_string_map[int]string_true", args{interface{}("1"), map[int]string{0:"0", 1:"1", 2:"2", 3:"3"}}, true},

		{"in_array_uint_map[int]uint_true", args{1, map[int]uint{0:0, 1:1, 2:2, 3:3}}, true},
		{"in_array_uint_map[int]string_true", args{1, map[int]string{0:"0", 1:"1", 2:"2", 3:"3"}}, true},
		{"in_array_string_map[int]uint_true", args{"1", map[int]uint{0:0, 1:1, 2:2, 3:3}}, true},
		{"in_array_string_map[int]string_true", args{"1", map[int]string{0:"0", 1:"1", 2:"2", 3:"3"}}, true},
		{"in_array_interface_uint_map[int]uint_true", args{interface{}(uint(1)), map[int]uint{0:0, 1:1, 2:2, 3:3}}, true},
		{"in_array_interface_uint_map[int]string_true", args{interface{}(uint(1)), map[int]string{0:"0", 1:"1", 2:"2", 3:"3"}}, true},
		{"in_array_interface_uint_map[int]uint_true", args{interface{}(uint(1)), map[int]uint{0:0, 1:1, 2:2, 3:3}}, true},
		{"in_array_interface_uint_map[int]string_true", args{interface{}(uint(1)), map[int]string{0:"0", 1:"1", 2:"2", 3:"3"}}, true},
		{"in_array_interface_string_map[int]uint_true", args{interface{}("1"), map[int]uint{0:0, 1:1, 2:2, 3:3}}, true},
		{"in_array_interface_string_map[int]uint_true", args{interface{}("1"), map[int]uint{0:0, 1:1, 2:2, 3:3}}, true},

		// false-map
		{"in_array_int_map[int]int_false", args{4, map[int]int{0:0, 1:1, 2:2, 3:3}}, false},
		{"in_array_int_map[int]string_false", args{4, map[int]string{0:"0", 1:"1", 2:"2", 3:"3"}}, false},
		{"in_array_string_map[int]int_false", args{"4", map[int]int{0:0, 1:1, 2:2, 3:3}}, false},
		{"in_array_string_map[int]string_false", args{"4", map[int]string{0:"0", 1:"1", 2:"2", 3:"3"}}, false},
		{"in_array_interface_int_map[int]int_false", args{interface{}(4), map[int]int{0:0, 1:1, 2:2, 3:3}}, false},
		{"in_array_interface_int_map[int]string_false", args{interface{}(4), map[int]string{0:"0", 1:"1", 2:"2", 3:"3"}}, false},
		{"in_array_interface_int_map[int]int_false", args{interface{}(4), map[int]int{0:0, 1:1, 2:2, 3:3}}, false},
		{"in_array_interface_int_map[int]string_false", args{interface{}(4), map[int]string{0:"0", 1:"1", 2:"2", 3:"3"}}, false},
		{"in_array_interface_string_map[int]int_false", args{interface{}("4"), map[int]int{0:0, 1:1, 2:2, 3:3}}, false},
		{"in_array_interface_string_map[int]string_false", args{interface{}("4"), map[int]string{0:"0", 1:"1", 2:"2", 3:"3"}}, false},
		{"in_array_interface_string_map[int]int_false", args{interface{}("4"), map[int]int{0:0, 1:1, 2:2, 3:3}}, false},
		{"in_array_interface_string_map[int]string_false", args{interface{}("4"), map[int]string{0:"0", 1:"1", 2:"2", 3:"3"}}, false},

		{"in_array_uint_map[int]uint_false", args{4, map[int]uint{0:0, 1:1, 2:2, 3:3}}, false},
		{"in_array_uint_map[int]string_false", args{4, map[int]string{0:"0", 1:"1", 2:"2", 3:"3"}}, false},
		{"in_array_string_map[int]uint_false", args{"4", map[int]uint{0:0, 1:1, 2:2, 3:3}}, false},
		{"in_array_string_map[int]string_false", args{"4", map[int]string{0:"0", 1:"1", 2:"2", 3:"3"}}, false},
		{"in_array_interface_uint_map[int]uint_false", args{interface{}(uint(4)), map[int]uint{0:0, 1:1, 2:2, 3:3}}, false},
		{"in_array_interface_uint_map[int]string_false", args{interface{}(uint(4)), map[int]string{0:"0", 1:"1", 2:"2", 3:"3"}}, false},
		{"in_array_interface_uint_map[int]uint_false", args{interface{}(uint(4)), map[int]uint{0:0, 1:1, 2:2, 3:3}}, false},
		{"in_array_interface_uint_map[int]string_false", args{interface{}(uint(4)), map[int]string{0:"0", 1:"1", 2:"2", 3:"3"}}, false},
		{"in_array_interface_string_map[int]uint_false", args{interface{}("4"), map[int]uint{0:0, 1:1, 2:2, 3:3}}, false},
		{"in_array_interface_string_map[int]uint_false", args{interface{}("4"), map[int]uint{0:0, 1:1, 2:2, 3:3}}, false},

		// others
		{"in_array_int_nil_false", args{1, nil}, false},
		{"in_array_int_[]int_false", args{1, []int{}}, false},
		{"in_array_int_map_false", args{1, map[int]int{}}, false},
		{"in_array_nil_nil_false", args{nil, []int{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InArray(tt.args.needle, tt.args.haystack); got != tt.want {
				t.Errorf("InArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInArrayInt(t *testing.T) {
	type args struct {
		a     int
		array []int
	}
	var x, y = 1, 4
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"in_array_int_true", args{x, []int{0, 1, 2, 3}}, true},
		{"in_array_int_false", args{y, []int{0, 1, 2, 3}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InArrayInt(tt.args.a, tt.args.array); got != tt.want {
				t.Errorf("InArrayInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInArrayInt16(t *testing.T) {
	type args struct {
		a     int16
		array []int16
	}
	var x, y int16 = 1, 4
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"in_array_int16_true", args{x, []int16{0, 1, 2, 3}}, true},
		{"in_array_int16_false", args{y, []int16{0, 1, 2, 3}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InArrayInt16(tt.args.a, tt.args.array); got != tt.want {
				t.Errorf("InArrayInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInArrayInt32(t *testing.T) {
	type args struct {
		a     int32
		array []int32
	}
	var x, y int32 = 1, 4
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"in_array_int32_true", args{x, []int32{0, 1, 2, 3}}, true},
		{"in_array_int32_false", args{y, []int32{0, 1, 2, 3}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InArrayInt32(tt.args.a, tt.args.array); got != tt.want {
				t.Errorf("InArrayInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInArrayInt64(t *testing.T) {
	type args struct {
		a     int64
		array []int64
	}
	var x, y int64 = 1, 4
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"in_array_int64_true", args{x, []int64{0, 1, 2, 3}}, true},
		{"in_array_int64_false", args{y, []int64{0, 1, 2, 3}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InArrayInt64(tt.args.a, tt.args.array); got != tt.want {
				t.Errorf("InArrayInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInArrayInt8(t *testing.T) {
	type args struct {
		a     int8
		array []int8
	}
	var x, y int8 = 1, 4
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"in_array_int8_true", args{x, []int8{0, 1, 2, 3}}, true},
		{"in_array_int8_false", args{y, []int8{0, 1, 2, 3}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InArrayInt8(tt.args.a, tt.args.array); got != tt.want {
				t.Errorf("InArrayInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInArrayStrict(t *testing.T) {
	type args struct {
		needle   interface{}
		haystack interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"in_array_strict_int_true", args{1, []int{0, 1, 2, 3}}, true},
		{"in_array_strict_int_false", args{"1", []int{0, 1, 2, 3}}, false},
		{"in_array_strict_interface_true", args{1, []interface{}{0, 1, "2", 3}}, true},
		{"in_array_strict_interface_false", args{1, []interface{}{0, "1", "2", 3}}, false},

		{"in_array_strict_map[int]int_true", args{1, map[int]int{0:0, 1:1, 2:2, 3:3}}, true},
		{"in_array_strict_map[int]int_false", args{"1", map[int]int{0:0, 1:1, 2:2, 3:3}}, false},
		{"in_array_strict_map[int]interface_true", args{1, map[int]interface{}{0:0, 1:1, 2:"2", 3:3}}, true},
		{"in_array_strict_map[int]interface_false", args{1, map[int]interface{}{0:0, 1:"1", 2:"2", 3:3}}, false},

		{"in_array_strict_[]int_false", args{1, []int{}}, false},
		{"in_array_strict_map[int]int_false", args{1, map[int]int{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func (t *testing.T) {
			if got := InArrayStrict(tt.args.needle, tt.args.haystack); got != tt.want {
				t.Errorf("InArrayStrict() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInArrayString(t *testing.T) {
	type args struct {
		a     string
		array []string
	}
	var x, y = "1", "4"
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"in_array_string_true", args{x, []string{"0", "1", "2", "3"}}, true},
		{"in_array_string_false", args{y, []string{"0", "1", "2", "3"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InArrayString(tt.args.a, tt.args.array); got != tt.want {
				t.Errorf("InArrayString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInArrayUInt(t *testing.T) {
	type args struct {
		a     uint
		array []uint
	}
	var x, y uint = 1, 4
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"in_array_uint_true", args{x, []uint{0, 1, 2, 3}}, true},
		{"in_array_uint_false", args{y, []uint{0, 1, 2, 3}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InArrayUInt(tt.args.a, tt.args.array); got != tt.want {
				t.Errorf("InArrayUInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInArrayUInt16(t *testing.T) {
	type args struct {
		a     uint16
		array []uint16
	}
	var x, y uint16 = 1, 4
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"in_array_uint16_true", args{x, []uint16{0, 1, 2, 3}}, true},
		{"in_array_uint16_false", args{y, []uint16{0, 1, 2, 3}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InArrayUInt16(tt.args.a, tt.args.array); got != tt.want {
				t.Errorf("InArrayUInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInArrayUInt32(t *testing.T) {
	type args struct {
		a     uint32
		array []uint32
	}
	var x, y uint32 = 1, 4
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"in_array_uint32_true", args{x, []uint32{0, 1, 2, 3}}, true},
		{"in_array_uint32_false", args{y, []uint32{0, 1, 2, 3}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InArrayUInt32(tt.args.a, tt.args.array); got != tt.want {
				t.Errorf("InArrayUInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInArrayUInt64(t *testing.T) {
	type args struct {
		a     uint64
		array []uint64
	}
	var x, y uint64 = 1, 4
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"in_array_uint64_true", args{x, []uint64{0, 1, 2, 3}}, true},
		{"in_array_uint64_false", args{y, []uint64{0, 1, 2, 3}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InArrayUInt64(tt.args.a, tt.args.array); got != tt.want {
				t.Errorf("InArrayUInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInArrayUInt8(t *testing.T) {
	type args struct {
		a     uint8
		array []uint8
	}
	var x, y uint8 = 1, 4
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"in_array_uint8_true", args{x, []uint8{0, 1, 2, 3}}, true},
		{"in_array_uint8_false", args{y, []uint8{0, 1, 2, 3}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InArrayUInt8(tt.args.a, tt.args.array); got != tt.want {
				t.Errorf("InArrayUInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}
