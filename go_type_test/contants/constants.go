// Go supports _constants_ of character, string, boolean,
// and numeric values.

package main

import (
	"fmt"
	"math"
)

// `const` declares a constant value.
const s string = "constant"

func main() {
	fmt.Println(s)

	// A `const` statement can appear anywhere a `var`
	// statement can.
	const n = 500000000

	// Constant expressions perform arithmetic with
	// arbitrary precision.
	const d = 3e20 / n
	fmt.Println(" 타입 정의없는 상수 ", d)

	// A numeric constant has no type until it's given
	// one, such as by an explicit conversion.
	fmt.Println("타입변환한 상수 출력 ", int64(d))

	// A number can be given a type by using it in a
	// context that requires one, such as a variable
	// assignment or function call. For example, here
	// `math.Sin` expects a `float64`.
	fmt.Println("수학함수에 적용한 ", math.Sin(n))
}
