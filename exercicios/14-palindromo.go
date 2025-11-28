package main

import (
	"fmt"
	s "strings"
)

func main() {
	texto := "abcbaa"
	textoSeparado := s.Split(texto, "")
	palindromo := true

	i := 0
	j := len(textoSeparado) - 1
	for i < j {
		if textoSeparado[i] != textoSeparado[j] {
			palindromo = false
			break
		}
		i++
		j--
	}

	if palindromo {
		fmt.Println("Palíndromo: ", texto)
	} else {
		fmt.Println("Não palíndromo: ", texto)
	}
}
