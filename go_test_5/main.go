package main

import (
	"fmt"

	"test.example.com/go_test_5/declaration"
	"test.example.com/go_test_5/packageTest"
)

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf(" boiling point = %g 화씨 or %g 섭씨 ", f, c)
	fmt.Println()
	fmt.Println("package", packageTest.A)

	fmt.Println("declaration", declaration.BoilingF)
	ftoc := declaration.FtoC(f)
	fmt.Println("declaration FtoC call ", ftoc)

}
