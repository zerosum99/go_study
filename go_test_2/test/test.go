package test

import "strconv"

var AAA int32 = 999

type Person struct {
	Name string
	Age  int
}

func (p Person) Get() string {
	aaa := strconv.Itoa(p.Age)
	return p.Name + aaa
}
