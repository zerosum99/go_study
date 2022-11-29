package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(" os args process ")
	s := strings.Join(os.Args[1:], " ")
	fmt.Println(s)

	s2 := strings.Split("Hello Wordl", " ")

	fmt.Println(s2)

}
