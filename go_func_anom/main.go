package main

import (
	"fmt"
	am "go_func_anom/anomfunc"
)

func main() {
	fmt.Println(" go func anomyous function ")

	defer func() {
		fmt.Println(" main end ")
	}()

	var f am.Function2 = add
	fmt.Println(" add call = ", f(100, 200))

	inner := outer(100)
	fmt.Println(" inner func call = ", inner(200))

	Inner := func(x int) func(int) int { // 익명함수로 내포된 함수를 만들어서 반환
		return func(y int) int {
			return x + y
		}
	}(100)
	fmt.Println(" anom func ", Inner(300)) // 반환된 결과를 실행하기
}

func add(x, y int) int {
	return x + y
}

func outer(x int) func(int) int { // 외부함수를 정의하고 내부 함수를 반환
	return func(y int) int { // 익명함수로 정의해서 함수를 반환하기
		return x + y
	}
}
