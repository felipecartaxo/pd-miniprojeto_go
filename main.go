package main

import (
	"errors"
	"fmt"
)

func main() {

	// slice de inteiros inicialmente vazio
	// que iremos manipular
	numbers := []int{}

	// opcao informada pelo usuario
	var option int

	for {
		// exibe o menu principal
		showMenu()

		// leitura da opcao do usuario
		fmt.Scan(&option) // & atribui o valor lido a variavel

		if option == 1 {
			// le o numero a ser adicionado
			var currentNumber int
			fmt.Print("Digite o número que deseja adicionar: ")
			fmt.Scan(&currentNumber)

			// adiciona o valor ao final do slice
			// e atualiza o proprio slice
			numbers = addNumber(numbers, currentNumber)
		}

		if option == 2 {
			listNumbers(numbers)
		}

		if option == 3 {
			// le o indice a ser removido
			var index int
			fmt.Print("Informe o índice do elemento que deseja remover: ")
			fmt.Scan(&index)

			// faz a remocao do valor e atualiza o slice
			numbers = removeNumber(numbers, index)
		}

		if option == 5 {
			var a, b int
			fmt.Println("Digite dois valores: ")
			fmt.Scan(&a, &b)

			fmt.Println(safeDivision(a, b))
		}

		if option == 6 {
			numbers = clearList(numbers)
		}
		if option == 0 {
			break
		}
	}
}

func showMenu() {

	fmt.Println("Digite um número de 0 a 6: ")
	fmt.Println("[1] Adicionar um número")
	fmt.Println("[2] Listar os números")
	fmt.Println("[3] Remover um número a partir do índice")
	fmt.Println("[4] Exibir estatísticas")
	fmt.Println("[5] Realizar divisão segura")
	fmt.Println("[6] Resetar a lista")
	fmt.Println("[0] Sair da aplicação")
}

func addNumber(list []int, num int) []int {
	return append(list, num)
}

func listNumbers(list []int) {
	fmt.Println("Números: ", list)
}

func removeNumber(list []int, index int) []int {
	// list[:index] = todos os elementos ate o indice
	// list[index+1:] = todos os elementos do indice em diante
	list = append(list[:index], list[index+1:]...)

	return list
}

// TODO: func minValue() {}
// TODO: func maxValue() {}
// TODO: func avg()      {}

func safeDivision(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Divisão por zero")
	}

	return a / b, nil
}

func clearList(list []int) []int {
	list = nil // remove todos os elementos do slice

	return list
}
