package main

import "fmt"

func main() {
	fmt.Println(" 초기값 처리 ")

	var s []int
	// var mm map[int]int  nil 로 처리되어 맵을 추가할 수 없음
	var m map[int]int = make(map[int]int) // map을 생성할 때는 반드시 make로 지정해서 기본값을 가지고 있어야 함

	s1 := append(s, 10) // nil 값이라도 슬라이스는 append 함수로 원소를 추가할 수 있음

	fmt.Println(s1)

	m[1] = 100
	fmt.Println(m[1])

	m3 := map[string]float64{ // Map literal
		"e":  2.71828,
		"pi": 3.1416,
	}
	fmt.Println(len(m3)) // 맵의 길이를 len 함수로 확인할 수 있음

	object := map[string]string{
		"name": "john doe",
	}

	arrayObject := []map[string]string{
		{"name": "john doe"},
		{"name": "jane doe"},
	}

	usingMake := make(map[string]string)
	usingMake["name"] = "max cavalera"

	fmt.Printf("this is object %v \n", object)
	fmt.Printf("this is array object %v \n", arrayObject)
	fmt.Printf("using make method %v \n", usingMake)
}
