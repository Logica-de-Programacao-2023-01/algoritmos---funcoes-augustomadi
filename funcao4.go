// Crie uma função que receba um slice de inteiros e retorne o segundo maior valor.
package main

import "fmt"

func main() {
	fmt.Print(segundo_maior([]int{5, 2, 6, 1, 7, 3, 2}))
}

func segundo_maior(lista []int) int {
	maior := lista[0]
	segundoMaior := lista[1]

	if segundoMaior > maior {
		maior, segundoMaior = segundoMaior, maior
	}

	for _, x := range lista[2:] {
		if x > maior {
			segundoMaior = maior
			maior = x
		} else if x > segundoMaior {
			segundoMaior = x
		}
	}
	return segundoMaior
}
