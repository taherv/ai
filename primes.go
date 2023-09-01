package main

import (
	"fmt"
	"math"
)

// Foo generates all primes less than n.
func Foo(n int) []int {
	primes := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		primes[i] = true
	}

	for p := 2; p*p <= n; p++ {
		if primes[p] {
			for i := p * p; i <= n; i += p {
				primes[i] = false
			}
		}
	}

	var primeNumbers []int
	for p := 2; p <= n; p++ {
		if primes[p] {
			primeNumbers = append(primeNumbers, p)
		}
	}

	return primeNumbers
}

// FindLargest6DigitPrime finds the largest 6-digit prime number from the list.
func FindLargest6DigitPrime(primes []int) int {
	for i := len(primes) - 1; i >= 0; i-- {
		if primes[i] >= 100000 && primes[i] <= 999999 {
			return primes[i]
		}
	}
	return -1
}

// FindMiddle4DigitPrime finds the middle 4-digit prime number from the list.
func FindMiddle4DigitPrime(primes []int) int {
	var fourDigitPrimes []int
	for _, prime := range primes {
		if prime >= 1000 && prime <= 9999 {
			fourDigitPrimes = append(fourDigitPrimes, prime)
		}
	}

	middle := len(fourDigitPrimes) / 2
	if len(fourDigitPrimes)%2 == 0 {
		return (fourDigitPrimes[middle-1] + fourDigitPrimes[middle]) / 2
	}
	return fourDigitPrimes[middle]
}

func main() {
	primes := Foo(1000000)

	largest6DigitPrime := FindLargest6DigitPrime(primes)
	fmt.Println("Largest 6-digit prime:", largest6DigitPrime)

	middle4DigitPrime := FindMiddle4DigitPrime(primes)
	fmt.Println("Middle 4-digit prime:", middle4DigitPrime)
}

