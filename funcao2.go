// Crie uma função que receba uma string e retorne a quantidade de vogais presentes.
package main

import "strings"

func main() {

}
func vogais_count(palavra string) int {
	vogais := "aeiouAEIOU"
	count := 0

	for _, carac := range palavra {
		if strings.ContainsRune(vogais, carac) {
			count++
		}
	}
	return count
}
