package main

import "fmt"

func main() {
	fmt.Println("go function test")

	msg := "Hello"
	say(msg)
	call(msg)
	refSay(&msg)
	println("changed msg : ", msg) //변경된 메시지 출력
}

func say(msg string) {
	fmt.Println("say func call : ", msg)
}

func refSay(msg *string) {
	fmt.Println("refSay call :", *msg)
	*msg = "Changed" //메시지 변경
}

func call(msg string) {
	fmt.Println("call func call : ", msg)
}
