package main

import "fmt"

type Living interface {
	GetName() string
	GetAge() int
}

func Get(l Living) {
	fmt.Println("Name = ", l.GetName())
	fmt.Println("Age = ", l.GetAge())
}

func main() {
	fmt.Println("Hello, 世界")
	p := Person{"Dahl", 55}
	Get(&p)

}

type Person struct {
	Name string
	Age  int
}

func (p *Person) GetName() string {
	return p.Name
}

func (p *Person) GetAge() int {
	return p.Age
}
