package main

import (
	"fmt"

	"golang.org/x/exp/constraints" // 제너릭 제한자를 처리하는 패키지
)

func main() {
	fmt.Println(" generic function process ")

	fmt.Println(add(100, 200))
	var a int8 = 100
	var b int8 = 22
	fmt.Println(add(a, b))

	c := 100.22
	d := 200.33
	fmt.Println(add(c, d))

	e := "Hello "
	f := "world"
	fmt.Println(add(e, f))

	fmt.Println(sum(e, f))

	var ua uint8 = 100
	var ub uint8 = 22
	fmt.Println(add(ua, ub))

	type Myint int16 // 별칭 타입을 지정

	var aa Myint = 100
	var ab Myint = 22
	fmt.Println(add(aa, ab))
}

type SInteger interface { // 틸트~ 는 별칭 타입도 포함해서 처리
	~int8 | ~int16 | ~int32 | ~int64 | ~int
}

type UInteger interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uint
}

type Float interface {
	float32 | float64
}

type Numeric interface {
	SInteger | UInteger | Float | string
}

func add[T Numeric](x, y T) T {
	return x + y
}

func sum[T constraints.Ordered](x, y T) T {
	return x + y
}
