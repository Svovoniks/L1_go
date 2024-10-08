package main

import (
	"fmt"
	"reflect"
)

func GetType(v interface{}) string {
	switch reflect.ValueOf(v).Kind() {
	case reflect.Int:
		return "int"
	case reflect.String:
		return "string"
	case reflect.Bool:
		return "bool"
	case reflect.Chan:
		return "channel"
	default:
		return "unknown"
	}
}

func main() {
	fmt.Println(GetType(1))
	fmt.Println(GetType("string"))
	fmt.Println(GetType(true))
	fmt.Println(GetType(make(chan int)))
}
