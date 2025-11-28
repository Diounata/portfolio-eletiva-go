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
		fmt.Println("Digite seu peso: ")
		fmt.Print("Peso (kg): ")

		peso, pesoErro := reader.ReadString('\n')

		fmt.Println("Digite sua altura: ")
		fmt.Print("Altura (m): ")

		altura, alturaErro := reader.ReadString('\n')

		if pesoErro != nil || alturaErro != nil {
			fmt.Println("Valor digitado não é um número válido!")
			continue
		}

		pesoValor, errPeso := strconv.ParseFloat(s.TrimSpace(peso), 64)
		alturaValor, errAltura := strconv.ParseFloat(s.TrimSpace(altura), 64)

		if errPeso != nil || errAltura != nil || alturaValor <= 0 {
			fmt.Println("Valores inválidos! Tente novamente.")
			continue
		}

		imc := pesoValor / (alturaValor * alturaValor)

		fmt.Printf("IMC: %.2f\n", imc)

		// Categorias
		categoria := ""

		switch {
		case imc < 18.5:
			categoria = "Abaixo do peso"
		case imc < 25:
			categoria = "Peso normal"
		case imc < 30:
			categoria = "Sobrepeso"
		case imc < 35:
			categoria = "Obesidade grau 1"
		case imc < 40:
			categoria = "Obesidade grau 2"
		default:
			categoria = "Obesidade grau 3"
		}

		fmt.Println("Categoria:", categoria, "\n")
	}
}
