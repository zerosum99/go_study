package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(" 함수 매개변수 확인 ")

	para(100)
	v := 100
	w := &v
	para(&v)
	para(w)

	para2(v, w)
}

func para(x interface{}) {
	fmt.Println(" ", x)
	fmt.Println(" 타입 ", reflect.TypeOf(x))
}

func para2(x ...interface{}) {
	fmt.Println(" ", x)
	fmt.Println(" 타입 ", reflect.TypeOf(x))
}
