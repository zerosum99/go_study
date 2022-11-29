package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(" random package process ")
	fmt.Println(time.Now())
	rand.Seed(time.Now().Unix())

	in := []int{2, 3, 5, 8}
	rand.Shuffle(len(in), func(i, j int) {
		in[i], in[j] = in[j], in[i]
	})
	fmt.Println(in)

	rand.Shuffle(len(in), func(i, j int) {
		in[i], in[j] = in[j], in[i]
	})
	fmt.Println(in)

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	randomNumber := r.Intn(45)

	fmt.Println("A random number: ", randomNumber)
}
