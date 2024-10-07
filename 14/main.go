package main

import (
	"fmt"
	"reflect"
)

func getType(v interface{}) string {
	switch reflect.ValueOf(v).Kind() {
	case reflect.Int:
		return "int"
	case reflect.String:
		return "string"
	case reflect.Bool:
		return "bool"
	case reflect.Chan:
		return "chan"
	default:
		return "unknown"
	}
}

func main() {
    fmt.Println(getType(1))
    fmt.Println(getType("string"))
    fmt.Println(getType(true))
    fmt.Println(getType(make(chan int)))
}
