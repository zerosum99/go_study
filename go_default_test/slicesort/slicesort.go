package main

import (
	"fmt"
	"sort"
)

type employee struct {
	name   string
	salary int
}

// 구조체를 관리하느 슬라이스 정의
type employeeList []employee

// 구조체에 대한 정렬을 위한 Len, Less, Swap 메소드를 정의
func (e employeeList) Len() int { // 구초체의 길이를 처리
	return len(e)
}

func (e employeeList) Less(i, j int) bool { // 연봉을 비교해서 참과거짓을 비교
	return e[i].salary < e[j].salary // 내림차순은 > 이고 올림차순은 < 로 비교
}

func (e employeeList) Swap(i, j int) { // 두 개의 구조체를 순서를 바꾸기
	e[i], e[j] = e[j], e[i]
}

func main() {

	fmt.Println(" 구조체 정렬하기 ")
	// 구조체의 슬라이스 생성,  구조체 3개의 원소를 정의
	eList := []employee{
		{name: "John", salary: 3000},
		{name: "Bill", salary: 4000},
		{name: "Sam", salary: 1000},
	}
	// 정렬은 구조체 자료형으로 변환하기
	sort.Sort(employeeList(eList))

	// 정렬한 결과를 출력하기
	for _, employee := range eList {
		fmt.Printf("Name: %s Salary %d\n", employee.name, employee.salary)
	}
}
