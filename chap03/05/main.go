package main

import "fmt"

type ubiratan int

var x ubiratan
var y int

func main() {
	fmt.Printf("%v\t%v\n", x, y)
	x = 42
	fmt.Printf("%v\n", x)

	y = int(x)
	fmt.Printf("%v %T\n", y, y)
}
