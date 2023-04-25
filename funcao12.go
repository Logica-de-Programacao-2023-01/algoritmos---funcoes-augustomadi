// Escreva uma função que receba um slice de strings como parâmetro e retorne uma string contendo apenas as strings que
// começam com uma letra maiúscula. Caso o slice esteja vazio, retorne um erro.
package main

import (
	"fmt"
	"strings"
)

func main() {
	paravras := []string{"Abacaxi", "banana", "Cachorro", "dente", "Elefante"}
	lista, err := caps(paravras)

	if err != nil {
		fmt.Print("Erro: ", err)
	} else {
		fmt.Print("Suas palavras com letras maiusculas são: ", lista)
	}
}

func caps(lista []string) (string, error) {
	if len(lista) == 0 {
		return "", fmt.Errorf("Sua lista esta vazia")
	}

	maiusculas := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lista_corrigida := []string{}
	for _, indices := range lista {
		if strings.Contains(maiusculas, string(indices[0])) {
			lista_corrigida = append(lista_corrigida, indices)
		}
	}

	return strings.Join(lista_corrigida, ""), nil
}
