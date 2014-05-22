// Summation of primes
package main

import "fmt"

const (
	limit = 2000000
)

func main() {
	var primes [limit]bool

	// Sieve of Eratosthenes
	for i := 2; i < len(primes); i++ {
		primes[i] = true
	}

	for i := 2; i <= math.Sqrt(limit); i++ {
		if primes[i] == true {
			// TODO: Fix this as per the implementation details found at: http://en.wikipedia.org/wiki/Sieve_of_Eratosthenes#Implementation
			for j := i*i; j <= limit; j = i*i+i*j++ {
				primes[j] = false
			}
		}
	}

	// Sum all the prime numbers
	var sumPrimes int
	for i := range primes {
		if primes[i] == true {
			sumPrimes += i
		}
	}

	fmt.Println("Summation of primes:", sumPrimes)
}
