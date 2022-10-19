package main

import (
	"fmt"
	"reflect"
)

func typeSwitch(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println("type:", reflect.ValueOf(value).Type())
		fmt.Println(value, "is int.")
	case string:
		fmt.Println("type:", reflect.ValueOf(value).Type())
		fmt.Println(value, "is string.")
	default:
		fmt.Println("type:", reflect.ValueOf(value).Type())
	}
}

func main() {
	typeSwitch(42)
	typeSwitch("42")
	typeSwitch([]float32{3.14})
}
