package main

import (
	"fmt"
)

func observe(i interface{}) {

	// using the format specifier
	// %T to check type in interface
	fmt.Printf("The type passed is: %T\n", i)

	// using the format specifier %#v
	// to check value in interface
	fmt.Printf("The value passed is: %#v \n", i)
	fmt.Println("-------------------------------------")
}

func main() {
	var value float64 = 15
	value2 := "GeeksForGeeks"
	observe(value)
	observe(value2)
}
