package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("slice test process")

	s := []int{1, 2, 3, 4, 5}
	s[2] = 777 // 슬라이스 요소값 변경
	ss := s[2]
	fmt.Println("slece element", ss)
	fmt.Println("slice", s)
	fmt.Println("slice type ", reflect.TypeOf(s))
	fmt.Println("slice value ", reflect.ValueOf(s))
	sss := reflect.ValueOf(s)
	fmt.Printf(" %T \n", sss)
	// sss[1] = 999 리플렉션으로 처리하면 실제 reflect.Value 자료형이라서 인덱스 검색을 지원하지 않음
	fmt.Println(sss.Interface().([]int))
	s11 := sss.Interface().([]int)
	s11[1] = 999
	fmt.Println("reflect -> []int ", s11)
	fmt.Println("slice value type ", reflect.ValueOf(s).Type())

	reverse(s)

	fmt.Println("reverse slice", s)

}

func reverse(r []int) {
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i] // 튜플 변환처리
	}
}
