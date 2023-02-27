package calculator

import "fmt"

var offset float64 = 1
var Offset float64 = 1

func Sum(a float64, b float64) float64 {
	fmt.Println("multiply: ", multiply(a, b))
	fmt.Println("Multiply: ", Multiply(a, b))
	return a + b + offset
}
