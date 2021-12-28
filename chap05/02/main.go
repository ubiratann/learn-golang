package main

import "fmt"

func main() {
	x := 2
	y := 2

	a := x != y
	b := x == y
	fmt.Printf("a:%v - b:%v\n", a, b)

	a = x <= y
	b = x >= y
	fmt.Printf("a:%v - b:%v\n", a, b)

	a = x < y
	b = x > y
	fmt.Printf("a:%v - b:%v\n", a, b)

}
