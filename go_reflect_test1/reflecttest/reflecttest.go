package main

import (
	"fmt"
	"reflect"
)

func main() {

	i := 100
	j := "문자열"
	t1 := reflect.TypeOf(i)
	k1 := t1.Kind()
	v1 := reflect.ValueOf(i)
	t2 := reflect.TypeOf(j)
	k2 := t2.Kind()
	v2 := reflect.ValueOf(j)

	fmt.Println("Type of first interface:", t1)
	fmt.Println("Kind of first interface:", k1)
	fmt.Printf("Kind of type: %T \n", k1.String())
	fmt.Println("Value of first interface:", v1)
	fmt.Println("Type of second interface:", t2)
	fmt.Println("Kind of second interface:", k2)
	fmt.Printf("Kind of type: %T  \n", k2.String())
	fmt.Println("Value of first interface:", v2)

	f1 := reflect.TypeOf(add)
	fk1 := f1.Kind()
	fv1 := reflect.ValueOf(add)
	fmt.Println("Type of function interface:", f1)
	fmt.Println("Kind of function interface:", fk1)
	fmt.Println("Value of function interface:", fv1)

	s1 := reflect.TypeOf([]int{1, 2, 3})
	sk1 := s1.Kind()
	sv1 := reflect.ValueOf([]int{1, 2, 3})
	fmt.Println("Type of slice interface:", s1)
	fmt.Println("Kind of slice interface:", sk1)
	fmt.Println("Value of slice interface:", sv1)
}

func add(x, y int) int {
	return x + y
}
