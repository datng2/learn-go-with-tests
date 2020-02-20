package main

import "reflect"

// Golang challenge:
//    Write a func:
//          walk(x interface{}, fn func(string)),
// which takes a struct x and calls fn for all *strings fields* found inside.
// difficulty level: recursively.

// Background:

// An interface:
// =============
// - Stores a pair: the concrete value assigned to the variable, and that
//   value's type descriptor.
// 		Example:
//          var r io.Reader
// 			tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
// 			r = tty // Schematically, (value, type) => (tty, *os.File)
//
// - Store any concrete (non-interface) value as long as that value
//   implements the interface's methods.
// 		Example:
//          var r io.Reader
// 			r = os.Stdin
// 			r = bufio.NewReader(r)

// interface{}
// ===========
//      Empty set of methods and is satisfied by **any value** at all,
//      since any value has zero or more methods.

func walk(obj interface{}, fn func(input string)) {
	v := reflect.ValueOf(obj)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	switch v.Kind() {

	case reflect.String:
		fn(v.String())

	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			// Not working: walk(val.Field(i), fn) why?
			walk(v.Field(i).Interface(), fn)
		}

	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			walk(v.Index(i).Interface(), fn)
		}

	case reflect.Map:
		for _, key := range v.MapKeys() {
			walk(v.MapIndex(key).Interface(), fn)
		}
	}
}
