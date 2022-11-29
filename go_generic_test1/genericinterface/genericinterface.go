package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(" generic interface constraints process ")

	u := User{}
	p := People{}

	Call(u)
	Call(p)

	var x MyInt = 100
	var y MyInt = 90

	PrintMin(x, y)

}

type Adder interface {
	Add()
}

type User struct {
}

func (u User) Add() {
	fmt.Println("User Add ")
}

type People struct{}

func (p People) Add() {
	fmt.Println("People Add ")
}

func Call[T Adder](a T) { // 타입 제한자 대신 인터페이스로 사용이 가능
	a.Add()
}

type Stringer interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~int
	String() string
}

func PrintMin[T Stringer](a, b T) {
	if a < b {
		fmt.Println(b.String())
	} else {
		fmt.Println(a.String())
	}
}

type MyInt int

func (m MyInt) String() string {
	return strconv.Itoa(int(m))
}
