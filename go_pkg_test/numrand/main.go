package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var stdin = bufio.NewReader(os.Stdin)

func inputIntValue() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		stdin.ReadString('\n')

	}
	return n, err
}

func main() {
	rand.Seed(time.Now().UnixNano())

	n := rand.Intn(100)
	fmt.Println(" num n = ", n)

	for {
		fmt.Printf("숫자를 입력하세요")
		n, err := inputIntValue()
		if err != nil {
			fmt.Println(" 숫자만 입력하세요")
		} else {
			fmt.Println("입력하신 숫자는 ", n, "입니다")
		}

		if n == 9999 {
			break
		}

	}
}
