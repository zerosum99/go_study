package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second) // 1초간 중단
	fmt.Println("done")
	done <- true // 채널 입력
}

func main() {
	fmt.Println(" go chan sync test ")
	done := make(chan bool, 1) // 채널 생성
	go worker(done)            // 채널을 전달

	fmt.Println(" <-done ", <-done) // 채널의 값을 출력
	time.Sleep(3 * time.Second)     // 다른 고루틴 실행될 수 있도록 메인 함수 고루틴 처리시간을 지연시킴

}
