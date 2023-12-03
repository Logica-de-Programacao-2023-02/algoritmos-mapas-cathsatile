package main

import (
	"fmt"
)

func somarContagensPalavras(listaMapas []map[string]int) map[string]int {
	somaContagens := make(map[string]int)

	for _, mapa := range listaMapas {
		for palavra, contagem := range mapa {
			somaContagens[palavra] += contagem
		}
	}

	return somaContagens
}

func main() {
	listaMapasExemplo := []map[string]int{
		{"hello": 1, "world": 2},
		{"hello": 3, "example": 1},
		{"world": 1, "example": 2},
	}

	resultadoSomaContagens := somarContagensPalavras(listaMapasExemplo)
	fmt.Println("Soma das contagens de palavras:")
	for palavra, contagem := range resultadoSomaContagens {
		fmt.Printf("%s: %d\n", palavra, contagem)
	}
}
