package main

import (
	"fmt"
	"sort"
	"strings"
)

func saoAnagramas(palavra1, palavra2 string) bool {
	palavra1 = strings.ToLower(palavra1)
	palavra2 = strings.ToLower(palavra2)

	slice1 := strings.Split(palavra1, "")
	slice2 := strings.Split(palavra2, "")

	sort.Strings(slice1)
	sort.Strings(slice2)

	return strings.Join(slice1, "") == strings.Join(slice2, "")
}
func agruparAnagramas(palavras []string) map[string][]string {
	gruposAnagramas := make(map[string][]string)

	for _, palavra := range palavras {
		adicionado := false

		for chave, grupo := range gruposAnagramas {
			if saoAnagramas(grupo[0], palavra) {
				gruposAnagramas[chave] = append(grupo, palavra)
				adicionado = true
				break
			}
		}

		if !adicionado {
			gruposAnagramas[palavra] = []string{palavra}
		}
	}

	return gruposAnagramas
}

func main() {
	palavrasExemplo := []string{"amor", "roma", "carro", "corra", "março", "macró"}

	resultadoGruposAnagramas := agruparAnagramas(palavrasExemplo)

	fmt.Println("Grupos de palavras anagramas:")
	for _, grupo := range resultadoGruposAnagramas {
		fmt.Println(grupo)
	}
}
