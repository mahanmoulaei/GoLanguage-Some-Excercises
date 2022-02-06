package main

import (
	"fmt"
)

func main() {
	checkPrime()
}

func checkPrime() {
	isPrime := false
	max := 50000
	var primeArray []int
	for number := 0; number <= max; number++ {
		for i := number; i >= 1; i-- {
			if number%i == 0 {
				if i == 1 || i == number {
					isPrime = true
				} else {
					isPrime = false
					break
				}
			}
		}
		if isPrime {
			primeArray = append(primeArray, number)
		}
	}

	fmt.Println("Array of primes between 1 and", max, "is:\n", primeArray)
}
