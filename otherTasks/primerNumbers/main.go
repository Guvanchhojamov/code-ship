package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findPrimes(11, 20))
}

func findPrimes(start, end int) []int {
	primes := []int{}

	for i := start; i < end; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

func isPrime(n int) bool {
	if n <= 3 {
		return true
	}

	if n <= 1 || (n%2 == 0 || n%3 == 0) {
		return false
	}

	limit := int(math.Sqrt(float64(n)))

	for i := 5; i <= limit; i++ {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}
