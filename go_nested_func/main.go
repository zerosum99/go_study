package main

import (
	"fmt"
)

func main() {
	fmt.Println(" golang nested function ")

	var counter int = 1

	func(str string) {                                       // 익명함수 정의 및 실행 
		fmt.Println("Hi", str, "I'm an anonymous function")
	}("Ricky")

	funcVar := func(str string) {                            // 익명하수 정의 및 변수 할당 
		fmt.Println("Hi", str, "I'm an anonymous function assigned to a variable.")
	}

	funcVar("Johnny")                // 익명함수 실행

	closure := func(str string) {                // 익명함수 정의 및 변수 할당 
		fmt.Println("Hi", str, "I'm a closure.")
		for i := 1; i < 5; i++ {
			fmt.Println("Counter incremented to:", counter)   // 외부함수의 지역변수인 자유변수 사용 
			counter++                            // 외부 함수의 지역변수를 사용해서 클로저 구성 
		}
	}

	fmt.Println("Counter is", counter, "before calling closure.")
	closure("Sandy")
	fmt.Println("Counter is", counter, "after calling closure.")
}
