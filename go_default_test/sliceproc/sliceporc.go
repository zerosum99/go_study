package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(" slice process ")

	var s []int
	if s == nil {
		fmt.Println(" nil 초기값 ")
	}
	fmt.Println(" slienc len cap  0 ", len(s), cap(s))
	s = append(s, 1, 2, 3, 4)

	fmt.Println(" slienc len cap  1 ", len(s), cap(s))

	slice1 := []int{1, 2, 3, 4, 5} // 슬라이스 리터럴오 지정하기
	slice2 := make([]int, 2, 10)   // 슬라이스 길이에 따라서 자료형에 따른 초기값이 들어감

	fmt.Printf("slice1(%p): len=%d cap=%d %v\n", slice1, len(slice1), cap(slice1), slice1)
	fmt.Printf("slice2(%p): len=%d cap=%d %v\n", slice2, len(slice2), cap(slice2), slice2)

	array := [5]int{1, 2, 3, 4, 5} // 배열 선언
	slice3 := array[1:2]           // 배열의 부분 슬라이싱 처리하면 슬라이스가 생김

	fmt.Printf("array: len=%d %v\n", len(array), array)
	fmt.Printf("slice: len=%d cap=%d %v\n", len(slice3), cap(slice3), slice3)

	slice4 := []int{1, 2, 3, 4, 5}
	slice6 := slice4[1:3]
	slice5 := slice4[1:3:4] // 완전 슬라이스 처리해서 cap을 조정

	fmt.Printf("slice4: len=%d cap=%d %v\n", len(slice4), cap(slice4), slice4)
	fmt.Printf("slice5: len=%d cap=%d %v\n", len(slice5), cap(slice5), slice5)
	fmt.Printf("slice6: len=%d cap=%d %v\n", len(slice6), cap(slice6), slice6)
	slice7 := append(slice6, 1, 2, 3, 4) // 완전슬라이스는 추가되는 값이 생기면 새로운 배열을 만들어서 슬라이스 처리
	// 기존 배열과 다른 배열이 만들어져서 기존 배열을 변경하지 않음
	fmt.Printf("slice7: len=%d cap=%d %v\n", len(slice7), cap(slice7), slice7)
	fmt.Printf("slice4: len=%d cap=%d %v\n", len(slice4), cap(slice4), slice4)

	// 슬라이스 삭제
	// 슬라이스는 별도의 함수가 없어서 삭제하려면 실제 삭제될 슬라이스를 제거하고 나머지 요소들만 남겨야 함
	slice8 := []int{1, 2, 3, 4, 5, 6}
	deleteIdx := 2
	// append 함수를 통한 슬라이스 삭제하기
	fmt.Println(slice8)
	// 삭제하기 전까지 슬라이스를 만들고 나머지 요소를 원소로 처리 즉 슬라이스를 지정하고 ... 연산을 통해 개별 원소 처리하기
	slice9 := append(slice8[:deleteIdx], slice8[deleteIdx+1:]...)
	fmt.Println(slice9)

	for i := deleteIdx + 1; i < len(slice8); i++ {
		slice8[i-1] = slice8[i]
	}
	slice10 := slice8[:len(slice9)-1]
	fmt.Println(slice10)

	//슬라이스 정렬은 sort 패키지로 처리
	slice11 := []int{6, 3, 1, 5, 4, 2}

	fmt.Println(slice11)
	sort.Ints(slice11)
	fmt.Println(slice11)

	// 문자열일 경우 정렬하기
	slice12 := []string{"abc", "aac"}
	sort.Strings(slice12)
	fmt.Println(slice12)

}
