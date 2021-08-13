package array

import (
	"reflect"
)

// Values array_values()
func Values(input interface{}) (result []interface{}) {
	v := reflect.ValueOf(input)
	if (v.Kind() != reflect.Map) && (v.Kind() != reflect.Struct) {
		return
	}
	switch v.Kind() {
	case reflect.Map:
		if v.Len() <= 0 {
			return
		}
		result = make([]interface{}, 0, v.Len())
		for _, key := range v.MapKeys() {
			result = append(result, v.MapIndex(key).Interface())
		}
	case reflect.Struct:
		if v.NumField() <= 0 {
			return
		}
		result = make([]interface{}, 0, v.NumField())
		for i := 0; i < v.NumField(); i++ {
			result = append(result, v.Field(i).Interface())
		}
	}
	return
}
