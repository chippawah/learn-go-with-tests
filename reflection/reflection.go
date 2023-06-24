package main

import "reflect"

func walk(x interface{}, fn func(string)) {
	val := getValue(x)
	numValues := 0
	var getField func(int) reflect.Value
	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		numValues = val.NumField()
		getField = val.Field
	case reflect.Slice:
		numValues = val.Len()
		getField = val.Index
	}
	for i := 0; i < numValues; i++ {
		walk(getField(i).Interface(), fn)
	}

}

func getValue(x interface{}) (val reflect.Value) {
	val = reflect.ValueOf(x)
	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}
	return
}
