package main

import (
	"fmt"
)

func main() {
	numero1 := 10
	numero2 := 12
	numerosIguais := numero1 == numero2

	if numerosIguais {
		fmt.Println("Os números são iguais.")
	} else {
		fmt.Println("Os números são diferentes.")
	}
}
