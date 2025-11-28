package main

import (
	"bufio"
	"fmt"
	"os"
	s "strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Digite um texto: ")
	texto, _ := reader.ReadString('\n')
	texto = s.TrimSpace(texto)
	textoSeparado := s.Split(texto, " ")

	palavras := make(map[string]int)
	for _, palavra := range textoSeparado {
		palavras[palavra] += 1
	}

	fmt.Println("Texto:", texto)
	fmt.Println("Padrões de ocorrência:")

	for chave := range palavras {
		fmt.Println(chave, "→", palavras[chave], "vezes")
	}
}
