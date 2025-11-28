package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	img, err := os.Open("imagem.ppm")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	defer img.Close()

	reader := bufio.NewReader(img)

	var formato string
	var largura, altura, valorMaximo int

	fmt.Fscan(reader, &formato)
	fmt.Fscan(reader, &largura, &altura)
	fmt.Fscan(reader, &valorMaximo)

	fmt.Printf("Formato: %s\n", formato)
	fmt.Printf("Dimensões: %dx%d\n", largura, altura)
	fmt.Printf("Valor máximo: %d\n", valorMaximo)
	fmt.Printf("Pixels:\n")

	var r, g, b int

	for i := 0; i < altura; i++ {
		for j := 0; j < largura; j++ {
			_, err := fmt.Fscan(reader, &r, &g, &b)
			if err != nil {
				fmt.Println("Erro lendo pixel:", err)
				return
			}
			fmt.Printf("%d %d %d\n", r, g, b)
		}
	}
}
