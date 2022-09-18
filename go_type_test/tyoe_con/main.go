package main

import "fmt"

func main() {
	fmt.Println(" 타입 변환 ")
	ii := 100
	fmt.Println(" type check")
	typeConv(ii)

}

func typeConv(s interface{}) {
	switch s.(type) {
	case int:
		fmt.Println(" int ")
	default:
		fmt.Println(" other type ")
	}
}
