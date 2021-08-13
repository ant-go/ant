package array

import (
	"reflect"
	"testing"
)

func TestValues(t *testing.T) {
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
		{"array_values_invalid", args{input: 1}, nil},
		{"array_values_map_[]interface{}_empty", args{input: map[int]int{}}, nil},
		{"array_values_map_[]interface{}", args{input: m}, []interface{}{1, 2}},
		{"array_values_struct_[]interface{}_empty", args{input: struct{}{}}, nil},
		{"array_values_struct_[]interface{}", args{input: structAB{}}, []interface{}{"", 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := Values(tt.args.input); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("Values() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
