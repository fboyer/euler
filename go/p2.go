// Even Fibonacci numbers
package main

import "fmt"

const (
	limit = 4000000
)

func main() {
	var sum int
	f1 := 1
	f2 := 2
	for f2 < limit {
		if f2%2 == 0 {
			sum += f2
		}
		f1, f2 = f2, f1+f2
	}

	fmt.Printf("Sum: %d", sum)
}
