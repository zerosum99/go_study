package main

import "fmt"

func main() {
	fmt.Println(" map/slcie test ")

	s := []int{1, 2, 3}
	m := map[int]int{1: 100, 2: 200}

	sp := &s[1] // 슬라이스 원소 즉 배열의 원소는 변수
	fmt.Printf(" %T %[1]v %[2]d \n", sp, *sp)

	// mp := &m[1]  맵의 요소는 변수가 아님
	fmt.Println(m)
}
