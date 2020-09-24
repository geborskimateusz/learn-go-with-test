package reflection

import (
	"fmt"
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

	if reflect.TypeOf(val).Kind() == reflect.Slice {
		fmt.Println(val.Kind())

		for i := 0; i < val.Len(); i++ {
			fmt.Println(val.Index(i))
		}
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			walk(field.Interface(), fn)
		}
	}
}

func getVal(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	fmt.Println(val)
	return val
}
