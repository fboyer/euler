// Smallest multiple
package main

import "fmt"

const (
	limit = 20
)

func main() {
	fmt.Printf("Smallest multiple: %v", findSmallestMultiple())
}

func findSmallestMultiple() int {
	// Multiple must be at least equal to the limit
	for i := limit; ; i++ {
		for j := limit; j > 0; j-- {
			// As soon as we find a number that is not divisible we stop the inner loop
			if i%j != 0 {
				break
			}
			// If we passed all tests then we return the value found
			if j == 1 {
				return i
			}
		}
	}
}
