package main

import (
	"fmt"
	"math"
)

func main() {
	const number = 7
	if number < 2 {
		printNotPrime(number)
	} else if number == 2 || number%2 == 0 {
		printPrime(number)
	} else {
		isPrime := findPrime(number)
		if isPrime {
			printPrime(number)
		} else {
			printNotPrime(number)
		}
	}
}

func findPrime(number int) bool {
	for i := 3; i <= int(math.Sqrt(float64(number))); i += 2 {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func printNotPrime(number int) {
	fmt.Printf("❌ Not prime → %d\n", number)
}

func printPrime(number int) {
	fmt.Printf("✅ Prime → %d\n", number)
}
