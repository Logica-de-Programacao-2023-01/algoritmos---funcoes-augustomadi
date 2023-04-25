// Crie uma função que receba um slice de inteiros e uma função como parâmetros. A função deve aplicar a
//função recebida em cada elemento do slice e retornar um novo slice com os resultados. Caso o slice esteja vazio,
//retorne um erro.

package main

import "fmt"

func função(num int) int {
	return num * 2
}

func multiplicação(lista []int, função func(int) int) ([]int, error) {
	if len(lista) == 0 {
		return nil, fmt.Errorf("Sua lista está vazia")
	}
	nova_lista := []int{}
	for i := 0; i < len(lista); i++ {
		nova_lista = append(nova_lista, função(lista[i]))
	}
	return nova_lista, nil
}
func main() {
	numeros := []int{42, 19, 7, 85, 63}
	lista, err := multiplicação(numeros, função)
	if err != nil {
		fmt.Print("Erro", err)
	} else {
		fmt.Print("Sua lista é: ", lista)
	}
}
