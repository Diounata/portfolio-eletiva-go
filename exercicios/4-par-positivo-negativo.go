package main

func main() {
	const number = 6
	const isEven = number%2 == 0
	const isPositive = number > 0
	if isEven {
		println("The number is even")
	} else if isPositive {
		println("The number is positive")
	} else {
		println("The number is not positive")
	}
}
