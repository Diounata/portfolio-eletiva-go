package main

import (
	"fmt"
	"slices"
)

func main() {
	numeros := []int{3, 5, 2, 8, 6}
	slices.Sort(numeros)

	fmt.Println("Números ordenados: ", numeros)
}
