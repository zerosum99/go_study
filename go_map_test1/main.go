package main

import "fmt"

func main() {
	fmt.Println(" call go map test 1 ")

	var m map[int]int
	fmt.Println(" map len = ", len(m))

	// m[0] = 111 nil 에는 키와 값을 추가할 수 없다.

	// 슬라이스와 다르게 초기화로 맵을 정의할 때는 항목들을 추가할 수 있다.
	mm := map[int]int{}
	mm[0] = 100
	fmt.Println(" map len = ", len(mm))

	mmm := make(map[int]int, 10)
	mmm[10] = 10000
	fmt.Println("mmm map len =", len(mmm))
	fmt.Println("mmm map =", mmm)
}
