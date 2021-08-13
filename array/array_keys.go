package array

import (
	"reflect"

	"github.com/spf13/cast"
)

// Keys array_keys()
func Keys(input interface{}) (result []interface{}) {
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
			result = append(result, key.Interface())
		}
	case reflect.Struct:
		if v.NumField() <= 0 {
			return
		}
		result = make([]interface{}, 0, v.NumField())
		for i := 0; i < v.NumField(); i++ {
			result = append(result, v.Type().Field(i).Name)
		}
	}
	return
}

// StringKeys array_keys()
func StringKeys(input interface{}) (result []string) {
	v := reflect.ValueOf(input)
	if (v.Kind() != reflect.Map) && (v.Kind() != reflect.Struct) {
		return
	}
	switch v.Kind() {
	case reflect.Map:
		if v.Len() <= 0 {
			return
		}
		result = make([]string, 0, v.Len())
		for _, key := range v.MapKeys() {
			s := cast.ToString(key.Interface())
			result = append(result, s)
		}
	case reflect.Struct:
		if v.NumField() <= 0 {
			return
		}
		result = make([]string, 0, v.NumField())
		for i := 0; i < v.NumField(); i++ {
			result = append(result, v.Type().Field(i).Name)
		}
	}
	return
}
