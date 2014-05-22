// Sum square difference
package main

import "fmt"

const (
	limit = 100
)

func main() {
	var sumSquare, squareSum int

	for i := 1; i <= limit; i++ {
		sumSquare += i * i
		squareSum += i
	}
	squareSum *= squareSum

	fmt.Println("Sum square difference:", squareSum-sumSquare)
}
