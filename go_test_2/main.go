package main

import (
	"fmt"
	"go_test_2/test"
)

func main() {
	fmt.Println("Hello world...")
	fmt.Println("package ", test.AAA)
	p := test.Person{"dahlmoon", 55}
	fmt.Println("peoele", p.Get())
}
