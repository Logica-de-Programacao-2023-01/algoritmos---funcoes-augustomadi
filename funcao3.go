// Crie uma função que receba um slice de strings e retorne a concatenação de todas as strings.
package main

import "fmt"

func main() {
	fmt.Print(frase([]string{"ola", "como", "vai"}))
}
func frase(lista []string) string {
	junção := ""
	for i := 0; i < len(lista); i++ {
		junção += lista[i] + " "
	}
	return junção
}
