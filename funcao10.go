// Escreva uma função que receba um slice de inteiros como parâmetro e retorne um novo slice com os valores ordenados de
// forma crescente. Caso o slice esteja vazio, retorne um erro.
package main

import (
	"fmt"
	"sort"
)

func main() {
	slice := []int{3, 45, 6, 7, 134, 26, 13, 5, 17}
	lista, err := ordem_cres(slice)
	if err != nil {
		fmt.Print("Erro: ", err)
	} else {
		fmt.Print(lista)
	}

}
func ordem_cres(numeros []int) ([]int, error) {
	if len(numeros) == 0 {
		return nil, fmt.Errorf("Sua slice esta vazia")
	}

	sort.Ints(numeros)

	return numeros, nil

}
