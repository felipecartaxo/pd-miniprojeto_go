package main

import (
	"fmt"
)

func main() {

	// slice de inteiros inicialmente vazio
	// que iremos manipular
	numbers := []int{1, 2, 3}

	// opcao informada pelo usuario
	var option int

	// exibe o menu principal
	showMenu()

	// leitura da opcao do usuario
	fmt.Scan(&option) // & atribui o valor lido a variavel

	if option == 2 {
		listNumbers(numbers)
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

func listNumbers(list []int) {
	fmt.Println("Números: ", list)
}
