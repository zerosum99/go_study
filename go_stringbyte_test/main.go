package main

import "fmt"

func main() {
	fmt.Println(" string byte process ")

	s := "Hello World"
	b := []byte(s) // 문자열을 바이트 슬라이스로 변환할 때 복사가 생김
	fmt.Println("string ;", string(b))
	b[1] = 65

	fmt.Println("string ;", s, "byte:", string(b))

	fmt.Printf("first=%[1]T,%[1]v  second =%[2]T %[2]v \n", s, b)
}
