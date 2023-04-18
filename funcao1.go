// Crie uma função que receba um slice de inteiros e retorne a média aritmética dos valores.
package main

import "fmt"

func main() {
	fmt.Print(media([]int{5, 2, 7, 3}))
}

func media(lista []int) int {
	soma := 0
	for i := 0; i < len(lista); i++ {
		soma += lista[i]
	}
	return soma / len(lista)
}
