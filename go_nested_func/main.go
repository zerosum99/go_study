package main

import (
	"fmt"
)

func main() {
	fmt.Println(" golang nested function ")

	var counter int = 1

	func(str string) {
		fmt.Println("Hi", str, "I'm an anonymous function")
	}("Ricky")

	funcVar := func(str string) {
		fmt.Println("Hi", str, "I'm an anonymous function assigned to a variable.")
	}

	funcVar("Johnny")

	closure := func(str string) {
		fmt.Println("Hi", str, "I'm a closure.")
		for i := 1; i < 5; i++ {
			fmt.Println("Counter incremented to:", counter)
			counter++
		}
	}

	fmt.Println("Counter is", counter, "before calling closure.")
	closure("Sandy")
	fmt.Println("Counter is", counter, "after calling closure.")
}
