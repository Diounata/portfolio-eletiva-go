package main

import (
	"fmt"
	"time"
)

func main() {
	dataNascimento := "2004-02-12"
	data, _ := time.Parse("2006-01-02", dataNascimento)

	fmt.Println("Data de nascimento: ", dataNascimento)
	fmt.Println("Dia da semana: ", data.Weekday())
}
