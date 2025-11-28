package main

import (
	"fmt"
	"slices"
)

func main() {
	texto := "lorem ipsum dolor sit amet"
	textoRuna := []rune(texto)
	slices.Sort(textoRuna)
	textoOrdenado := string(textoRuna)

	fmt.Println("Texto ordenado:", textoOrdenado)
}
