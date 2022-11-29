package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(" map 관련 처리 ")

	m := make(map[string]int)
	// name[key] = val로 key/value 쌍을 저장합니다.
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)

	v1 := m["k1"] // 맵의 값을 변수에 할당
	fmt.Println("v1: ", v1)
	//len을 맵에 사용하면 key/value 쌍의 갯수를 반환합니다.
	fmt.Println("len:", len(m))

	delete(m, "k2") // 맵의 키를 삭제
	fmt.Println("map:", m)

	_, ok := m["k2"]        // 삭제된 맵의 키를 조회하면 값은 초기값으로 리턴
	fmt.Println("prs:", ok) // 그래서 실제 값이 있는지 확인하기 위해 추가변수 설정 처리

	n := map[string]int{"foo": 1, "bar": 2} // 맵을 초기값을 지정해서 작성
	fmt.Println("map:", n)
	fmt.Println(" map type ", reflect.TypeOf(n)) // 리플렉션 모듈을 사용해서 맵의 자료형 확인

}
