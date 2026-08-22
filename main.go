package main

import (
	"errors"
	"fmt"
)

func main() {
	numbers := []int{}

	var option int

	for {
		showMenu()

		fmt.Scan(&option)
		fmt.Println()

		if option == 1 {
			var currentNumber int
			fmt.Print("Digite o numero que deseja adicionar: ")
			fmt.Scan(&currentNumber)

			numbers = addNumber(numbers, currentNumber)
		}

		if option == 2 {
			listNumbers(numbers)
		}

		if option == 3 {
			var index int
			fmt.Print("Informe o indice do elemento que deseja remover: ")
			fmt.Scan(&index)

			var err error
			numbers, err = removeNumber(numbers, index)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}

		if option == 4 {
			if err := showStatistics(numbers); err != nil {
				fmt.Println(err)
			}
		}

		if option == 5 {
			var a, b int
			fmt.Print("Digite dois valores: ")
			fmt.Scan(&a, &b)

			result, err := safeDivision(a, b)
			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Println(result)
		}

		if option == 6 {
			numbers = clearList(numbers)
		}

		if option < 0 || option > 6 {
			fmt.Println("Informe um valor valido")
		}

		if option == 0 {
			fmt.Println("Saindo...")
			fmt.Println()

			break
		}
	}
}

func showMenu() {
	fmt.Println("\n======================================")
	fmt.Println("[1] Adicionar um numero")
	fmt.Println("[2] Listar os numeros")
	fmt.Println("[3] Remover numero a partir de um indice")
	fmt.Println("[4] Exibir estatisticas")
	fmt.Println("[5] Realizar divisao segura")
	fmt.Println("[6] Limpar a lista")
	fmt.Println("[0] Sair da aplicacao")

	fmt.Print("\nDigite um numero de 0 a 6: ")
}

func addNumber(list []int, num int) []int {
	return append(list, num)
}

func listNumbers(list []int) {
	fmt.Println("Numeros: ", list)
}

func removeNumber(list []int, index int) ([]int, error) {
	if index < 0 || index >= len(list) {
		return list, errors.New("indice invalido")
	}

	// list[:index] = todos os elementos ate o indice
	// list[index+1:] = todos os elementos do indice em diante
	list = append(list[:index], list[index+1:]...)

	return list, nil
}

func showStatistics(numbers []int) error {
	min, err := minValue(numbers)
	if err != nil {
		return err
	}

	max, err := maxValue(numbers)
	if err != nil {
		return err
	}

	average, err := avg(numbers)
	if err != nil {
		return err
	}

	fmt.Println("Menor valor: ", min)
	fmt.Println("Maior valor: ", max)
	fmt.Println("Media: ", average)

	return nil
}

func minValue(list []int) (int, error) {
	if len(list) == 0 {
		return 0, errors.New("lista vazia")
	}

	var min int = list[0]

	for i := 1; i < len(list); i++ {
		if list[i] < min {
			min = list[i]
		}
	}

	return min, nil
}

func maxValue(list []int) (int, error) {
	if len(list) == 0 {
		return 0, errors.New("lista vazia")
	}

	var max int = list[0]

	for i := 1; i < len(list); i++ {
		if list[i] > max {
			max = list[i]
		}
	}

	return max, nil
}

func avg(list []int) (int, error) {
	if len(list) == 0 {
		return 0, errors.New("lista vazia")
	}

	total := 0
	for i := 0; i < len(list); i++ {
		total += list[i]
	}

	average := total / len(list)

	return average, nil
}

func safeDivision(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divisao por zero")
	}

	return a / b, nil
}

func clearList(list []int) []int {
	list = nil

	return list
}
