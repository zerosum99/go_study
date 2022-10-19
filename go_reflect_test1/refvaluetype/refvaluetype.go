package main

import (
	"fmt"
	"reflect"
)

func typeOf(i interface{}) {
	fmt.Println(" typeOf call ", reflect.TypeOf(i))
}

func valueOf(i interface{}) {
	fmt.Println(" valueof call ", reflect.ValueOf(i))
}

func main() {
	fmt.Println("reflect value type ")

	var i = 100
	valueOf(i)
	typeOf(i)

	var j int16 = 100
	valueOf(j)
	typeOf(j)
}
