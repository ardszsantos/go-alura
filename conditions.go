package main

import (
	"fmt"
)

func main() {
	fmt.Println("**** Conditions ****")

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("3- Sair do Programa")

	var comando int
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi", comando)

	if comando == 1 {
		fmt.Println("Monitorando...")
	} else if comando == 2 {
		fmt.Println("EXibindo Logs...")
	} else if comando == 3 {
		fmt.Println("Saindo do programa...")
	} else {
		fmt.Println("NÃO conheço esse comando.")
	}

}
