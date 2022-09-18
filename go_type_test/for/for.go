// `for` is Go's only looping construct. Here are
// some basic types of `for` loops.

package main

import "fmt"

func main() {

	// The most basic type, with a single condition.
	// 조건식만 있는 순환
	i := 1
	fmt.Println(" conditional for ")
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// A classic initial/condition/after `for` loop.
	// 일반 순환
	fmt.Println("  for ")
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// `for` without a condition will loop repeatedly
	// until you `break` out of the loop or `return` from
	// the enclosing function.
	// 무한순환 이때는 내부적으로 break를 통해서 순환 중단
	fmt.Println("  loop for ")
	for {
		fmt.Println("loop")
		break
	}

	// You can also `continue` to the next iteration of
	// the loop.
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	// range 순환
	l := []int{1, 2, 3, 4, 5, 6}
	for i, v := range l {
		fmt.Println(i, v)
	}
}
