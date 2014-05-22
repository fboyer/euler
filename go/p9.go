// Special Pythagorean triplet
package main

import "fmt"

const (
	limit = 1000
)

func main() {
	var a, b, c = 3, 4, 5
	for i := 0; i < limit; i++ {
		for j := 0; j < i; j++ {
			for k := 0; k < j; k++ {
				if k*k+j*j == i*i && i+j+k == limit {
					a = k
					b = j
					c = i
				}
			}
		}
	}

	fmt.Println("Special Pythagorean triplet:", a*b*c)
}
