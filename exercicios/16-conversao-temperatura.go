package main

import "fmt"

func main() {
	temperaturaEmCelsius := float32(18)
	temperaturaEmFahrenheit := float32(60)

	conversaoParaFahrenheit := paraFahrenheit(temperaturaEmCelsius)
	conversaoParaCelsius := paraCelsius(temperaturaEmFahrenheit)

	fmt.Println("Temperatura de Celsius para Fahrenheit")
	fmt.Println(temperaturaEmCelsius, "°C →", conversaoParaFahrenheit, "°F")

	fmt.Println("Temperatura de Fahrenheit para Celsius")
	fmt.Println(temperaturaEmFahrenheit, "°F →", conversaoParaCelsius, "°C")

}

func paraFahrenheit(temperaturaEmCelsius float32) float32 {
	return (temperaturaEmCelsius * 9 / 5) + 32
}

func paraCelsius(temperaturaEmFahrenheit float32) float32 {
	return (temperaturaEmFahrenheit - 32) * 5 / 9
}
