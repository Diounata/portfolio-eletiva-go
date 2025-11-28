package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"os"
)

func converterRGBParaJPG(r, g, b [][]uint8, nomeArquivo string) error {
	altura := len(r)
	largura := len(r[0])

	imagem := image.NewRGBA(image.Rect(0, 0, largura, altura))

	for y := 0; y < altura; y++ {
		for x := 0; x < largura; x++ {
			imagem.Set(x, y, color.RGBA{
				R: r[y][x],
				G: g[y][x],
				B: b[y][x],
				A: 255,
			})
		}
	}

	arquivo, err := os.Create(nomeArquivo)
	if err != nil {
		return err
	}
	defer arquivo.Close()

	return jpeg.Encode(arquivo, imagem, &jpeg.Options{Quality: 90})
}

func main() {
	r := [][]uint8{{255, 0}, {0, 255}}
	g := [][]uint8{{0, 255}, {255, 0}}
	b := [][]uint8{{0, 0}, {255, 255}}

	if err := converterRGBParaJPG(r, g, b, "saida.jpg"); err != nil {
		panic(err)
	}
}
