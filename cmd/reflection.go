package main

import (
	"fmt"
	"reflect"
)

/*
Reflection in computing is the ability of a program to examine its own structure,
particularly through types; it's a form of metaprogramming.
It's also a great source of confusion
*/

func walk(x interface{}, fn func(s string)) {
	val := getInterfaceValue(x)

	numOfVals := 0
	var getField func(i int) reflect.Value

	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		numOfVals = val.Len()
		getField = val.Index
	case reflect.Struct:
		numOfVals = val.NumField()
		getField = val.Field
	case reflect.String:
		fn(val.String())
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walk(val.MapIndex(key).Interface(), fn)
		}
	case reflect.Chan:
		for {
			if vl, ok := val.Recv(); ok {
				walk(vl.Interface(), fn)
			} else {
				break
			}
		}
	case reflect.Func:
		reflectFn := val.Call(nil)
		for _, res := range reflectFn {
			walk(res.Interface(), fn)
		}
	default:
		fmt.Printf("unhandled type kind -> %v: field -> %v\n", val.Kind(), val.Interface())
	}

	for i := 0; i < numOfVals; i++ {
		walk(getField(i).Interface(), fn)
	}
}

func getInterfaceValue(unknown interface{}) reflect.Value {
	val := reflect.ValueOf(unknown)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return val
}
