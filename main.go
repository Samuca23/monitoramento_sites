package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	for {
		exibeIntroducao()
		comando := leComando()
		validaComando(comando)
	}

}

func exibeIntroducao() {
	nome := "Samuel"
	versao := 1.1

	fmt.Println("Olá, Sr.", nome)
	fmt.Println("Este programa está na versão", versao)
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir os Logs")
	fmt.Println("0 - Sair")
}

func leComando() int {
	var comando int
	fmt.Scan(&comando)

	return comando
}

func validaComando(comando int) {

	switch comando {
	case 1:
		iniciarMonitoramento()
	case 2:
		fmt.Printf("Exibindo Logs")
	case 0:
		sair()
	default:
		fmt.Print("Não conheço esse comando")
		os.Exit(-1)
	}
}

func iniciarMonitoramento() {
	fmt.Printf("Monitorando...")
	sites := []string{"https://www.alura.com.br", "https://www.unidavi.com.br"}

	for i := 0; i < 5; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
	}

}

func testaSite(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "esta com problemas. Status Code:", resp.StatusCode)
	}
}

func sair() {
	fmt.Printf("Saindo")
	os.Exit(0)
}
