// In Go, _variables_ are explicitly declared and used by
// the compiler to e.g. check type-correctness of function
// calls.

package main

import "fmt"

func main() {

	fmt.Println(" variables test ")

	// `var` declares 1 or more variables.
	var a = "initial"
	fmt.Println("변수와 초기화 ", a)

	// You can declare multiple variables at once.
	var b, c int = 1, 2
	fmt.Println("여러 변수 정의 및 초기화 ", b, c)

	// Go will infer the type of initialized variables.
	var d = true
	fmt.Println("변수와 초기화 ", d)

	// Variables declared without a corresponding
	// initialization are _zero-valued_. For example, the
	// zero value for an `int` is `0`.
	var e int
	fmt.Println("변수 정의, 초기값 세팅 ", e)

	// The `:=` syntax is shorthand for declaring and
	// initializing a variable, e.g. for
	// `var f string = "apple"` in this case.
	f := "apple"
	fmt.Println("짤른 변수 선언, 지역변수로만 정의 ", f)
}
