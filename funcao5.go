// Crie uma função que receba um slice de inteiros e um valor inteiro, e retorne a posição do primeiro elemento igual ao valor no slice. Caso não encontre, retorne -1.
package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	value := 3
	fmt.Printf("O índice do valor %d no slice %v é %d\n", value, slice, indice(slice, value))
}

func indice(lista []int, valor int) int {
	for indice, val := range lista {
		if val == valor {
			return indice
		}
	}
	return -1
}
