package main

import "fmt"

// 타입제한자 정의

type Number interface {
	int | int16 | int32 | int64 | float32 | float64
}

func main() {
	fmt.Println(" arth test process ")
	fmt.Println(" int add generic", add(100, 200))
	fmt.Println(" float add generic", add(100.2, 200.3))
}

func add[T Number](x, y T) T {
	return x + y
}
