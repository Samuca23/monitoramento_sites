package main

import (
	"fmt"
)

func main() {
	nome := "Samuel"
	versao := 1.1

	fmt.Println("Olá, Sr.", nome)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir os Logs")
	fmt.Println("0 - Sair")
	var comando int
	fmt.Scan(&comando)
}
