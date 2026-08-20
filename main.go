package main

import (
	"fmt"
)

func main() {

	// Exibe o menu principal
	showMenu()
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
