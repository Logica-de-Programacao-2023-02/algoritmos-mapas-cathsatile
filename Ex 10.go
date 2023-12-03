package main

import "fmt"

func contarParesFrequencia(slice []int) map[[2]int]int {
	frequenciaPares := make(map[[2]int]int)

	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			pair := [2]int{slice[i], slice[j]}
			frequenciaPares[pair]++
		}
	}

	return frequenciaPares
}

func main() {
	sliceExemplo := []int{1, 2, 3, 2, 4, 5, 1, 3, 4, 5}
	resultadoFrequenciaPares := contarParesFrequencia(sliceExemplo)
	fmt.Println("Frequência de pares de números:")
	for par, frequencia := range resultadoFrequenciaPares {
		fmt.Printf("%v: %d\n", par, frequencia)
	}
}
