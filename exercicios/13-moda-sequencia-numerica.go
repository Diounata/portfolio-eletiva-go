package main

import (
	"fmt"
)

func main() {
	numeros := []int{1, 2, 2, 3, 4, 4, 4, 5, 5}
	numerosFrequencia := make(map[int]int)
	var numeroMaisFrequente int
	maiorSequencia := 0

	for _, numero := range numeros {
		numerosFrequencia[numero]++
		numeroFrequenciaAtual := numerosFrequencia[numero]

		if numeroFrequenciaAtual > maiorSequencia {
			maiorSequencia = numeroFrequenciaAtual
			numeroMaisFrequente = numero
		}
	}

	fmt.Println(numerosFrequencia)
	fmt.Println("Moda: ", numeroMaisFrequente)
}
