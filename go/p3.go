// Largest prime factor
package main

import "fmt"

const (
	fact = 600851475143
)

func main() {
	largestPrimeFactor := fact

	for i := 2; i < largestPrimeFactor; i++ {
		if largestPrimeFactor%i == 0 {
			largestPrimeFactor = largestPrimeFactor / i
			i = 2
		}
	}

	fmt.Printf("Largest prime factor: %d", largestPrimeFactor)
}
