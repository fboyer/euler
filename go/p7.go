// 10001st prime
package main

import "fmt"

const (
	limit = 10001
)

func main() {
	var counter = 0

	var i, j int
	for i = 2; counter < limit; i++ {
		for j = 2; j < i; j++ {
			if i%j == 0 {
				break
			}
		}
		if i == j {
			counter++
		}
	}

	fmt.Println("10001st prime:", i-1)
}
