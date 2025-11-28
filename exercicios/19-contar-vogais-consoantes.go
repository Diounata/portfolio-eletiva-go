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
	textoSeparado := s.Split(texto, "")

	vogais := 0
	consoantes := 0
	for _, char := range textoSeparado {
		char = s.ToLower(char)
		if char == "a" || char == "e" || char == "i" || char == "o" || char == "u" {
			vogais++
			continue
		}
		consoantes++
	}

	fmt.Println("Total de letras:", len(textoSeparado))
	fmt.Println("Total de vogais:", vogais)
	fmt.Println("Total de consoantes:", consoantes)
}
