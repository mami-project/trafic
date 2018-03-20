package config

import (
	"fmt"
	"reflect"
)

func AppendKeyVal(args []string, key string, val interface{}) []string {
	// Assumption is that val is of a comparable type.
	// If so, we can check whether it is the zero value for its type,
	// in which case we do nothing.  Otherwise, we append the
	// supplied key and the value to the string array.
	if val != reflect.Zero(reflect.TypeOf(val)).Interface() {
		args = append(args, key)
		args = append(args, fmt.Sprintf("%v", val))
	}

	return args
}

func AppendKey(args []string, key string, val interface{}) []string {
	if val != reflect.Zero(reflect.TypeOf(val)).Interface() {
		args = append(args, key)
	}

	return args
}
