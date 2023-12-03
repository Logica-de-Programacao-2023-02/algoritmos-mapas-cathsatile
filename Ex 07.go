package main

import (
	"fmt"
	"strings"
)

func contarLetras(palavra string) map[rune]int {
	contagemLetras := make(map[rune]int)

	for _, char := range palavra {
		contagemLetras[char]++
	}

	return contagemLetras
}

func contagensLetrasPorPalavra(frase string) map[string]map[rune]int {
	contagensPalavras := make(map[string]map[rune]int)

	palavras := strings.Fields(frase)
	for _, palavra := range palavras {
		contagensPalavras[palavra] = contarLetras(palavra)
	}

	return contagensPalavras
}

func main() {
	fraseExemplo := "go é uma linguagem de programação"
	resultadoContagens := contagensLetrasPorPalavra(fraseExemplo)

	fmt.Println("Contagens de letras por palavra:")
	for palavra, contagemLetras := range resultadoContagens {
		fmt.Printf("%s: %v\n", palavra, contagemLetras)
	}
}
