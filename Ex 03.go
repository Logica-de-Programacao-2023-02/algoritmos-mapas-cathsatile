package main

import "fmt"

func somarValoresMapa(mapaInteiros map[string]int) int {
	soma := 0

	for _, valor := range mapaInteiros {
		soma += valor
	}

	return soma
}

func main() {

	mapaExemplo := map[string]int{
		"um":     1,
		"dois":   2,
		"três":   3,
		"quatro": 4,
	}

	resultadoSoma := somarValoresMapa(mapaExemplo)

	fmt.Printf("Soma dos valores no mapa: %d\n", resultadoSoma)
}
