package main

import (
	"fmt"
	"go_method_test1/test"
)

func main() {
	fmt.Println(" go method test 1 exec ")
	x := test.Add(100, 200)
	fmt.Println("Add =", x)

	ar := test.Arth{X: 100, Y: 200}
	fmt.Println("Arth add =", ar.Add())

	var add test.Adder = ar
	fmt.Println("interface = ", add.Add())

	pp := test.NewPerson("dahlmoon")
	fmt.Println("person =", pp.Name, pp.Age)

	fmt.Printf("person type  = %T \n", test.Person{"Jon", 55})
	fmt.Printf("person value = %v \n", test.Person{"Jon", 55})

	fmt.Println(test.Person{Name: "Alice", Age: 30})

	s := test.Person{Name: "Sean", Age: 50}
	fmt.Println(s.Name, s.Age)

	s.Age = 77
	fmt.Println(s.Name, s.Age)

	sp := &s
	fmt.Println(sp.Age)
	sp.Age = 51
	fmt.Println(sp.Age)
	fmt.Println(s.Name, s.Age)

	call() //익명함수 호출

}

// call 익명함수 처리
// 함수 내부에 익명함수 정의하고 이를 호출해서 처리 및 출력
func call() {
	sum := func(n ...int) int { //익명함수 정의
		s := 0
		for _, i := range n {
			s += i
		}
		return s
	}

	result := sum(1, 2, 3, 4, 5) //익명함수 호출
	println(result)
}
