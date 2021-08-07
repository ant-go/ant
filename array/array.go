package array

import (
	"reflect"
)

func ChangeLowerCase(v interface{}) (m map[string]interface{}) {
	return
}

func Merge() {
	//
}

func getFields(v reflect.Value, tagName string) {
	if v.Kind() == reflect.Ptr || v.Kind() == reflect.UnsafePointer {
		v = v.Elem()
	}
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if tag := field.Tag.Get(tagName); tag == "-" {
			continue
		}
		// field.Type
	}
}
