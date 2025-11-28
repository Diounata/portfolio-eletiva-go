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

	for {
		fmt.Print("1° número: ")

		n1, n1Erro := reader.ReadString('\n')

		fmt.Print("2° número: ")

		n2, n2Erro := reader.ReadString('\n')

		if n1Erro != nil || n2Erro != nil {
			fmt.Println("Valor digitado não é um número válido!")
			continue
		}

		n1Valor, errn1 := strconv.ParseInt(s.TrimSpace(n1), 10, 64)
		n2Valor, errn2 := strconv.ParseInt(s.TrimSpace(n2), 10, 64)

		if errn1 != nil || errn2 != nil {
			fmt.Println("Valores inválidos! Tente novamente.")
			continue
		}

		menorNumero := n1Valor
		maiorNumero := n2Valor
		if menorNumero > maiorNumero {
			temp := menorNumero
			menorNumero = maiorNumero
			maiorNumero = temp
		}

		mmc := int64(0)
		for n := maiorNumero; mmc == 0; n += maiorNumero {
			if n%menorNumero == 0 {
				mmc = n
				fmt.Println("MMC:", mmc)
				continue
			}
		}
	}
}
