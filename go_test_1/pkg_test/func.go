package pkg_test

//함수, 변수, 상수 호출 처리 알아보기

import "fmt"

// 이름이 대문자는 퍼블릭
func Add(x int, y int) int {
	return x + y
}

var pVar int = 20

func Count(x int) {
	pVar += x

	fmt.Println(" Count", pVar)

}
