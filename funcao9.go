// Crie uma função que receba uma string como parâmetro e retorne um novo slice com todas as palavras contidas na string.
// Considere que as palavras são separadas por espaços em branco. Caso a string seja vazia, retorne um erro.
package main

import (
	"fmt"
	"strings"
)

func main() {
	string := ""
	palavra, err := palavras(string)
	if err != nil {
		fmt.Print("Erro: ", err)
	} else {
		fmt.Print(palavra)
	}
}

func palavras(palavra string) ([]string, error) {
	if len(palavra) == 0 {
		return nil, fmt.Errorf("sua lista está vazia")
	}
	lista_de_palavras := strings.Split(palavra, " ")
	return lista_de_palavras, nil
}
