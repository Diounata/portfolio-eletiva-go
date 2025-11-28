package main

import (
	"fmt"
)

func registrarMovimento(disco int, origem string, destino string) {
	fmt.Println("Mover peça", disco, "de", origem, "para", destino)
}

func hanoi(n int, origem string, auxiliar string, destino string) {
	if n == 1 {
		registrarMovimento(n, origem, destino)
		return
	}
	hanoi(n-1, origem, destino, auxiliar)
	registrarMovimento(n, origem, destino)
	hanoi(n-1, auxiliar, origem, destino)
}

func main() {
	quantidadeDiscos := 3
	hanoi(quantidadeDiscos, "A", "B", "C")
}
