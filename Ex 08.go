package main

import (
	"fmt"
)

func calcularEquilibrioContas(gastos map[string]float64) map[string]float64 {
	var totalGastos float64
	for _, valor := range gastos {
		totalGastos += valor
	}
	mediaGastos := totalGastos / float64(len(gastos))

	equilibrio := make(map[string]float64)
	for pessoa, gasto := range gastos {
		equilibrio[pessoa] = mediaGastos - gasto
	}

	return equilibrio
}

func main() {
	gastosExemplo := map[string]float64{
		"Alice":   30.0,
		"Bob":     20.0,
		"Charlie": 25.0,
		"David":   15.0,
	}
	resultadoEquilibrio := calcularEquilibrioContas(gastosExemplo)
	fmt.Println("Valores que cada pessoa deve receber ou pagar para igualar as despesas:")
	for pessoa, valor := range resultadoEquilibrio {
		fmt.Printf("%s: %.2f\n", pessoa, valor)
	}
}
