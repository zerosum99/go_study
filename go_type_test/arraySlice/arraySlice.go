package main

import "fmt"

func main() {
	fmt.Println("array vs slice ")

	var ar = [5]int{1, 2, 3, 4, 5}

	fmt.Printf(" %v \n ", ar)

	fmt.Printf(" %v ", arrayProc())

}

// 특정 배열 생성해서 반환하는 함수
func arrayProc() interface{} {
	const a int = 10
	b := [a]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// len=0, cap=10 인 슬라이스
	sliceA := make([]int, 0, 10)

	// 계속 한 요소씩 추가
	for i := 0; i < len(b); i++ {
		sliceA = append(sliceA, i)
		// 슬라이스 길이와 용량 확인
		fmt.Println(len(sliceA), cap(sliceA))
	}

	return sliceA

}
