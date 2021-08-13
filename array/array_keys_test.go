package array

import (
	"reflect"
	"testing"
)

type structAB struct {
	A string
	B int
}

func TestKeys(t *testing.T) {
	m := map[interface{}]interface{}{
		"a": 1,
		0:   2,
	}
	type args struct {
		input interface{}
	}
	tests := []struct {
		name       string
		args       args
		wantResult []interface{}
	}{
		{"array_keys_invalid", args{input: 1}, nil},
		{"array_keys_map_[]interface{}_empty", args{input: map[int]int{}}, nil},
		{"array_keys_map_[]interface{}", args{input: m}, []interface{}{"a", 0}},
		{"array_keys_struct_[]interface{}_empty", args{input: struct{}{}}, nil},
		{"array_keys_struct_[]interface{}", args{input: structAB{}}, []interface{}{"A", "B"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := Keys(tt.args.input); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("Keys() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestStringKeys(t *testing.T) {
	m := map[interface{}]interface{}{
		"a": 1,
		0:   2,
	}
	type args struct {
		input interface{}
	}
	tests := []struct {
		name       string
		args       args
		wantResult []string
	}{
		{"array_keys_invalid", args{input: 1}, nil},
		{"array_keys_map_[]string{}_empty", args{input: map[int]int{}}, nil},
		{"array_keys_map_[]string{}", args{input: m}, []string{"a", "0"}},
		{"array_keys_struct_[]string{}_empty", args{input: struct{}{}}, nil},
		{"array_keys_struct_[]string{}", args{input: structAB{}}, []string{"A", "B"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := StringKeys(tt.args.input); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("StringKeys() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
