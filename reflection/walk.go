package reflection

import (
	"reflect"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func walk(x interface{}, fn func(input string)) {

	val := getVal(x)
	walkByValue := func(val reflect.Value) {
		walk(val.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.Func:
		res := val.Call(nil)
		for _, el := range res {
			walk(el.Interface(), fn)
		}
	case reflect.Chan:
		for v, ok := val.Recv(); ok; v, ok = val.Recv() {
			walk(v.Interface(), fn)
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkByValue(val.MapIndex(key))
		}
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkByValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkByValue(val.Index(i))
		}
	case reflect.String:
		fn(val.String())
	}

}

func getVal(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return val
}
