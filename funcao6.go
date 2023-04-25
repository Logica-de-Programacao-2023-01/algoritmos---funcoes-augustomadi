// Escreva uma função que receba um slice de strings como parâmetro e retorne uma string com todas as strings
// concatenadas e separadas por vírgulas. Caso o slice esteja vazio, retorne um erro.
package main

import "fmt"

func main() {
	fmt.Print(junção([]string{"ola", "como", "vai"}))
}
func junção(lista []string) string {
	junção := ""
	for i := 0; i < len(lista); i++ {
		junção += lista[i] + ","
	}
	return junção
}
