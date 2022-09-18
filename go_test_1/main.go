package main

import (
	"fmt"
	cfn "go_test_1/pkg_test"
)

func main() {
	fmt.Println("hello world")

	fmt.Println(" Add ", cfn.Add(100, 200))
	cfn.Count(777)
}
