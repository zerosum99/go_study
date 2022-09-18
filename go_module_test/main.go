package main

//외부 모듈에 대해 go.mod 내에 require, replace 작성 필요
import (
	"fmt"

	"aaa.com/go_module_call/pkg_call"
)

func main() {
	fmt.Println(" go module test ")

	pkg_call.AAA()
}
