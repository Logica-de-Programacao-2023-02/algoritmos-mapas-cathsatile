package main

import (
	"fmt"
)

func fibonacciAteN(n int) map[int]int {
	resultado := make(map[int]int)

	a, b := 0, 1
	indice := 0

	for a <= n {
		resultado[indice] = a
		a, b = b, a+b
		indice++
	}

	return resultado
}

func main() {
	numeroLimite := 50
	resultadoFibonacci := fibonacciAteN(numeroLimite)
	fmt.Printf("Sequência de Fibonacci até %d:\n", numeroLimite)
	for indice, valor := range resultadoFibonacci {
		fmt.Printf("%d: %d\n", indice, valor)
	}
}
