package main

import (
	"fmt"
	gn "go_generic_test1/generic"
)

func main() {
	// 정수 계산
	var a int = 1
	var b int = 2
	fmt.Println(gn.Add(a, b))
	fmt.Println(" add Int  =", addInt(a, b))
	fmt.Println(" mul Int =", gn.Mul(a, b))

	fmt.Println(" mul Number =", gn.Product(a, b))
	// 실수 계산
	var f1 float64 = 3.14
	var f2 float64 = 1.43
	fmt.Println(gn.Add(f1, f2))

	fmt.Println(" add Float =", addFloat(f1, f2))
	fmt.Println(" mul Number =", gn.Product(f1, f2))
}

func addInt(x, y int) int {
	return x + y
}

func addFloat(x, y float64) float64 {
	return x + y
}
