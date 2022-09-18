package main

import (
	"fmt"
	"math/cmplx"
)

var globalVar = 100

const MAX int32 = 9999

func main() {
	fmt.Println(" hello go_vs_2 package ")
	fmt.Println(" global variable = ", globalVar)
	fmt.Println(" global constant = ", MAX)

	fmt.Println(" add = ", add(100, 200))
	var x = 100
	y := 400
	fmt.Println(" add = ", add(x, y))

	a := complex(2.5, 3.1)
	b := complex(10.2, 2)
	fmt.Println(" a + b = ", a+b)
	fmt.Println(" a - b = ", a-b)
	fmt.Println(" a * b = ", a*b)
	fmt.Println(" a / b = ", a/b)
	fmt.Println(" a.real = ", real(a))
	fmt.Println(" a.image = ", imag(a))
	fmt.Println(" ABS a = ", cmplx.Abs(a))

}

func add(x int, y int) int {
	return x + y
}
