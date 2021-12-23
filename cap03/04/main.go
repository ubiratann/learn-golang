package main

import "fmt"

type ubiratan int

var x ubiratan

func main() {

	fmt.Printf("valor = %v\ntipo = %T\n", x, x)
	x = 42
	fmt.Println(x)
}
