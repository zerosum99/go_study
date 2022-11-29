package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type KFunc2 func(int, int) int

func main() {
	fmt.Println(" function test call ")
	fmt.Println(" call add : ", add(100, 200))

	callFunc(add, 10, 30)
	var fc KFunc2 = add
	fc.String()

	FuncAnalyse(add)

}

func add(x, y int) int {
	return x + y
}

func callFunc(x KFunc2, y int, z int) {
	fmt.Println(" call Function  : ", x(y, z))
}

func (f KFunc2) String() {
	fmt.Println(" argument ", reflect.ValueOf(f))
}

func FuncAnalyse(m interface{}) {

	// 함수의 타임을 확인한다.
	x := reflect.TypeOf(m)

	numIn := x.NumIn()   //함수 내부의 매개변수 갯수
	numOut := x.NumOut() //함수의 반환값 갯수

	fmt.Println("함수나 메소드 :", x.String())
	fmt.Println("가변인자여부 :", x.IsVariadic()) // Used (<type> ...) ?
	fmt.Println("소속 패키지의 경로 :", x.PkgPath())

	fmt.Println("입력 매개변수 갯수  :", numIn)

	for i := 0; i < numIn; i++ {

		inV := x.In(i)
		in_Kind := inV.Kind() //func
		fmt.Printf("\n 매개변수 IN: "+strconv.Itoa(i)+"\nKind: %v\nName: %v\n-----------\n", in_Kind, inV.Name())
	}

	fmt.Println("반환값 갯수  :", numOut)
	for o := 0; o < numOut; o++ {

		returnV := x.Out(0)
		return_Kind := returnV.Kind()
		fmt.Printf("\n 반환값 OUT: "+strconv.Itoa(o)+"\nKind: %v\nName: %v\n", return_Kind, returnV.Name())
	}
}
