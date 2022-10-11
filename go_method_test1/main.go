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

}
