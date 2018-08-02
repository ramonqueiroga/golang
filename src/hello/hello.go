package main

import "fmt"
import "os"

func main() {
	exibeIntroducao()
	existeMenu()

	comando := leComando()

	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo logs...")
	case 0:
		fmt.Println("Saindo do programa...")
		os.Exit(-1)
	default:
		fmt.Println("Não existe essa opção")
		os.Exit(-1)
	}

}

func existeMenu() {
	fmt.Println("0- Sair do programa")
	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir logs")
}

func exibeIntroducao() {
	nome := "Ramon"
	versao := 1.1
	fmt.Println("Olá. Sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func leComando() int {
	var comando int
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi", comando)
	return comando
}
