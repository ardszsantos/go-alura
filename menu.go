package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {
	readFileSites()
	exibeIntroducao()
	for {

		exibeMenu()

		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs....")
		case 3:
			fmt.Println("Saindo do programa")
			os.Exit(0)

		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	nome := "Douglas"
	versao := 1.1
	fmt.Println("Olá, sr. ", nome)
	fmt.Println("Este programa está na versão: ", versao)
	fmt.Println(" ")
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("3- Sair do Programa")
	fmt.Println(" ")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	sites := readFileSites()

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando o", site, "na posição", i)
			testaSite(site)
			fmt.Println(" ")
		}

		time.Sleep(delay * time.Second)
		fmt.Println(" ")
	}

	fmt.Println(" ")
}

func testaSite(site string) {

	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site: ", site, "foi carregado corretamente")
	} else {
		fmt.Println("Site: ", site, "está com um problema. Status Code:", resp.StatusCode)
	}
}

func readFileSites() []string {
	var sites []string

	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	leitor := bufio.NewReader(arquivo)

	linha, err := leitor.ReadString('\n')

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	fmt.Println(linha)

	return sites
}
