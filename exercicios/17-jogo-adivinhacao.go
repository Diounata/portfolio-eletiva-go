package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	numeroAleatorio := int64(rand.IntN(100))

	fmt.Println("Adivinha um valor entre 0 e 100:")

	for {
		valorTexto, _ := reader.ReadString('\n')
		valorTexto = strings.TrimSpace(valorTexto)

		valor, err := strconv.ParseInt(valorTexto, 10, 64)
		if err != nil {
			fmt.Println("Digite um número válido!")
			continue
		}

		if valor > numeroAleatorio {
			fmt.Println("Número é menor que", valor)
			continue
		}

		if valor < numeroAleatorio {
			fmt.Println("Número é maior que", valor)
			continue
		}

		fmt.Println("Você acertou o número!")
		break
	}
}
