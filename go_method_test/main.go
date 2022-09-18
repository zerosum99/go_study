package main

import (
	"fmt"
	mt "go_method_test/method_test"
)

func main() {
	fmt.Println("go method test ")
	u := mt.User{"dahlmoon", "55"}
	fmt.Println("user name", u.Get())
}
