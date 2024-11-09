package main

import (
	"fmt"
)

func main() {

	nome := "Douglas"
	idade := 24

	fmt.Println("Olá sr.", nome, "sua idade é ", idade)

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir os logs")
	fmt.Println("0- Sair do Programa")

	var comando int

	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi", comando)
}
