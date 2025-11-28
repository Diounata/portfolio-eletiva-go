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
		fmt.Println("Digite um número: ")

		texto, erro := reader.ReadString('\n')
		if erro != nil {
			fmt.Println("Digite um número válido!")
			continue
		}

		valor, _ := strconv.ParseInt(s.TrimSpace(texto), 10, 64)
		fatorial := int64(1)
		for i := valor; i > 1; i-- {
			fatorial *= i
		}

		fmt.Println("Fatorial de", valor, "=", fatorial)
	}
}
