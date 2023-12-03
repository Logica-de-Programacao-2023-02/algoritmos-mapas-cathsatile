package main

import (
	"fmt"
)

func contarFrequenciaCaracteres(texto string) map[rune]int {
	frequencia := make(map[rune]int)

	for _, char := range texto {
		frequencia[char]++
	}

	return frequencia
}

func main() {
	textoExemplo := "hello world"
	resultadoFrequencia := contarFrequenciaCaracteres(textoExemplo)

	fmt.Println("Frequência de caracteres:")
	for char, freq := range resultadoFrequencia {
		fmt.Printf("'%c': %d\n", char, freq)
	}
}
