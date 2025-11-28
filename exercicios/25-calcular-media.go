package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	s "strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	numeros := []int64{}
	for {
		fmt.Print("Insira número (0 para calcular média): ")

		numero, numeroErro := reader.ReadString('\n')
		numeroValor, numeroValorErro := strconv.ParseInt(s.TrimSpace(numero), 10, 64)

		if numeroValorErro != nil || numeroErro != nil {
			fmt.Println("Valor inválido! Tente novamente.")
			continue
		}

		if numeroValor != 0 {
			numeros = append(numeros, numeroValor)
			continue
		}

		if len(numeros) == 0 {
			fmt.Println("Nenhum número inserido para calcular a média.")
			continue
		}

		var somaNumeros int64 = 0
		for _, numero := range numeros {
			somaNumeros += numero
		}

		media := float64(somaNumeros) / float64(len(numeros))
		numeros = numeros[:0]
		fmt.Println("Média de números:", media)
	}
}
