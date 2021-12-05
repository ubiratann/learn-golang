package main

import "fmt"

const prefixoOla = "Olá "

func OlaMundo() string {
	return prefixoOla + "mundo"
}

func OlaVoce(nome string) string {
	return prefixoOla + nome
}

func main() {
	fmt.Println(OlaMundo())
	fmt.Println(OlaVoce("Ubiratan"))
}
