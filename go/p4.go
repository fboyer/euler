// Largest palindrome product
package main

import (
	"fmt"
	"strconv"
)

const (
	limit = 999
)

func main() {
	var largestPalindrome, product int

	for i := limit; i > 0; i-- {
		for j := limit; j > 0; j-- {
			product = i * j
			if isPalindrome(product) {
				if product > largestPalindrome {
					largestPalindrome = product
				}
			}
		}
	}

	fmt.Printf("Largest palindrome product: %v", largestPalindrome)
}

func isPalindrome(num int) bool {
	numArray := []byte(strconv.Itoa(num))
	lenNum := len(numArray)
	for i := 0; i < lenNum; i++ {
		if i > lenNum-i-1 {
			break
		}
		if numArray[i] != numArray[lenNum-i-1] {
			return false
		}
	}

	return true
}
