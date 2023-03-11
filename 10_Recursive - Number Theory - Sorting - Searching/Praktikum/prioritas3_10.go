package main

import (
	"fmt"
)

func isPrime(number int) bool {
	if number < 2 {
		return false
	}
	for i := 2; i*i <= number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func getPrime(n, current, primeCount int) int {
	if primeCount == n {
		return current - 1
	}
	if isPrime(current) {
		return getPrime(n, current+1, primeCount+1)
	}
	return getPrime(n, current+1, primeCount)
}

func primeX(number int) int {
	return getPrime(number, 2, 0)
}

func main() {
	fmt.Println(primeX(1))
	fmt.Println(primeX(5))
	fmt.Println(primeX(8))
	fmt.Println(primeX(9))
	fmt.Println(primeX(10))
}
