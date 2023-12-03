package main

import "fmt"

func mesclarMapas(mapa1, mapa2 map[string]string) map[string]string {
	resultado := make(map[string]string)

	for chave, valor := range mapa1 {
		resultado[chave] = valor
	}

	for chave, valor := range mapa2 {
		resultado[chave] = valor
	}

	return resultado
}

func main() {

	mapa1 := map[string]string{"a": "apple", "b": "banana", "c": "cherry"}
	mapa2 := map[string]string{"b": "blueberry", "d": "date", "e": "elderberry"}

	resultado := mesclarMapas(mapa1, mapa2)

	fmt.Println("Mapa Mesclado:")
	for chave, valor := range resultado {
		fmt.Printf("%s: %s\n", chave, valor)
	}
}
