package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println("Qual vamor de i: ", i)
	}

	valor := 0
	teste := true

	for teste {
		valor ++
		if valor%5 == 0 {
			teste = false
		}
		fmt.Println("o numero é: ", valor)
	}

	for {
		valor--
		fmt.Println("o numero é: ", valor)
		if valor == 0 {
			break
		}
	}

	texto := "Eu adoro escrever programas usando Go"

	for indice, letra := range texto{
		fmt.Printf("Texto(%d) = %q\r\n", indice, letra)
	}

}