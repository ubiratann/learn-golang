package main

import "fmt"

func main() {
	x := 2
	fmt.Printf("x : %d - %b %#x\n", x, x, x)

	y := x << 1
	fmt.Printf("y : %d - %b %#x\n", y, y, y)

}
