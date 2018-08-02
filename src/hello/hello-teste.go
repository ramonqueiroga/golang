package main

import "fmt"

func main() {
	nomeDigitado := getNome()
	fmt.Println("O nome digitado foi", nomeDigitado)
}

func getNome() string {
	fmt.Println("Digite um nome")
	var nome string
	fmt.Scan(&nome)
	return nome
}
