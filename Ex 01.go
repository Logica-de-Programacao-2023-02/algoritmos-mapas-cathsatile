package main

import (
	"fmt"
	"strings"
)

func contarOcorrenciasPalavras(texto string) map[string]int {
	ocorrencias := make(map[string]int)

	palavras := strings.Fields(texto)

	for _, palavra := range palavras {
		ocorrencias[palavra]++
	}

	return ocorrencias
}

func main() {
	textoExemplo := "go é uma linguagem de programação go é eficiente go é poderosa"

	resultado := contarOcorrenciasPalavras(textoExemplo)

	fmt.Println("Ocorrências de palavras:")
	for palavra, ocorrencias := range resultado {
		fmt.Printf("%s: %d\n", palavra, ocorrencias)
	}
}
