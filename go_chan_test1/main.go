package main

import (
	"fmt"
)

func main() {
	fmt.Println(" go chan test ")

	messages := make(chan string) // 채널을 생성, 채널도 포인터로 생성됨
	fmt.Printf(" chan type = %T \n", messages)

	fmt.Println("messages = ", messages)

	go func() { messages <- "ping" }() // 고루틴 함수 실행하고 채널에 데이터 입력

	msg := <-messages // 메인 고루틴에서 채녈의 데이터를 변수에 할당
	fmt.Println(msg)  // 변수에 할당된 것을 출력

}
