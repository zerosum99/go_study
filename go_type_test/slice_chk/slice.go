package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func Get(u *User) {
	fmt.Println("function call ", u.Name, u.Age)
}

func (u *User) Get() {
	fmt.Println("method call ", u.Name, u.Age)
}

func main() {

	a := User{"dahlmoon", 55}

	Get(&a)
	(&a).Get()
	a.Get()

}
