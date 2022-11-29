package main

import (
	"fmt"
	"math/rand"
)

// Main function
func main() {

	// Finding random numbers of int type
	// Using Intn() function
	res_1 := rand.Intn(45)
	res_2 := rand.Intn(45)
	res_3 := rand.Intn(45)
	res_4 := rand.Intn(45)
	res_5 := rand.Intn(45)
	res_6 := rand.Intn(45)

	// Displaying the result
	fmt.Println("Random Number 1: ", res_1)
	fmt.Println("Random Number 2: ", res_2)
	fmt.Println("Random Number 3: ", res_3)
	fmt.Println("Random Number 1: ", res_4)
	fmt.Println("Random Number 2: ", res_5)
	fmt.Println("Random Number 3: ", res_6)
}
