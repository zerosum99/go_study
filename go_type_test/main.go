package main

import "fmt"

func main() {
	fmt.Println(" type test ")

	var ui uint16 = 100
	fmt.Println(" uint 16 = ", ui)
	var ui32 uint32 = 200
	// 타입이 다르므로 타입을 변환한 후에 덧셈을 실행함
	uadd := ui + uint16(ui32)
	fmt.Println(" type convert add ", uadd)
}
