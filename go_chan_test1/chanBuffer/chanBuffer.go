package main

import (
	"fmt"
)

func main() {
	fmt.Println(" channel buffer process ")

	messages := make(chan string, 2) // 채널 버퍼 지정
	messages <- "first msg "         // 채널에 메시지 전달
	messages <- "second msg "

	fmt.Println(len(messages))
	fmt.Println(<-messages) // 채널에서 메시지 꺼내기
	fmt.Println(<-messages)

	fmt.Println(len(messages)) // 현재 채녈의 길이 확인
}
