package declaration

import (
	"fmt"
)

const BoilingF = 212.0

func FtoC(f float64) float64 {

	fmt.Println("declaration package")
	return (f - 32) * 5 / 9
}
