package main

import "fmt"

func ping(pings chan<- string, msg string) { // 함수에 매개변수로 채널의 저장할 수 있도록  전달..
	fmt.Println(" ping call ")
	pings <- msg // 채널에 전달된 인자를 넣기
}

func pong(pings <-chan string, pongs chan<- string) { // 함수의 매개변수에 채널의 정보를 저장 및 꺼낼 수 있도록 전달
	fmt.Println(" pong call ")
	msg := <-pings // 채널에서 정보를 꺼내고 변수에 할당
	pongs <- msg   // 다른 채널에 입력으로 저장
}

func main() {
	pings := make(chan string, 1) // 채널 버퍼가 1일 두개의 채널을 생성
	pongs := make(chan string, 1)

	ping(pings, "passed message") // 먼저 채널에 값을 저장하는 함수 호출
	pong(pings, pongs)            // 채널의 값을 꺼내서 다른 채널에 전달
	fmt.Println(<-pongs)          // 최종 결과를 채널에서 꺼내서 출력
}
