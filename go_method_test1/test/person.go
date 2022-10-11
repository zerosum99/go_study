package test

type Person struct {
	Name string
	Age  int
}

func NewPerson(name string) Person {
	p := Person{Name: name}
	p.Age = 42
	return p
}
