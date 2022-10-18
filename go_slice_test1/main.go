package main

import "fmt"

func main() {
	// var a []int  즉 nil 값일 경우는 요소를 추가할 수 없다
	// var a []int
	b := []int{}
	fmt.Println("slice len = ", len(b), cap(b)) // cap이 0일때는 추가되지 않음

	c := []int{0}
	c[0] = 100

	// c[1] = 999  인덱스로 값을 추가 하면 예외 발생함
	c = append(c, 999) // 슬라이스일 때 추가는 append 함수로 처리해야함

	fmt.Println(" slice len = ", len(c), cap(c))
	d := make([]int, 0, 10)
	d = append(d, 300)
	fmt.Println(" slice len = ", len(d), cap(d))
}
