// Crie uma função que receba um número inteiro como parâmetro e retorne verdadeiro se ele for um número primo e falso
// caso contrário. Caso o número seja negativo, retorne um erro.
package main

import "fmt"

func main() {
	numero, err := primos(5)
	if err != nil {
		fmt.Print("Erro: ", err)
	} else {
		fmt.Print(numero)
	}
}
func primos(num int) (bool, error) {
	for i := 2; i < num; i++ {
		if num%i == 0 || num < 2 {
			return false, nil
		}
	}
	if num < 0 {
		return false, fmt.Errorf("Seu numero é menor que 0")
	} else {
		return true, nil
	}

}
