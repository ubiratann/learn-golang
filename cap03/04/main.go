package main

import "fmt"

type ubiratan int

var x ubiratan

func main() {

	fmt.Printf("value = %v\ntype = %T\n", x, x)
	x = 42
	fmt.Println(x)
}
