// Escreva uma função que receba um slice de inteiros como parâmetro e retorne um novo slice
// com apenas os números pares contidos no slice. Caso o slice esteja vazio, retorne um erro.
package main

import "fmt"

func main() {
	fmt.Print(num_pares([]int{5, 2, 16, 7, 8}))
}

func num_pares(numeros []int) ([]int, error) {
	if len(numeros) == 0 {
		return nil, fmt.Errorf("Sua lista está vazia")
	}
	numeros_pares := []int{}
	for i := 0; i < len(numeros); i++ {
		if numeros[i]%2 == 0 {
			numeros_pares = append(numeros_pares, numeros[i])
		}
	}
	return numeros_pares, nil
}
