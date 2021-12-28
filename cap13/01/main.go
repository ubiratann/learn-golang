package main

import "fmt"

func main() {
	x := returnInt()
	y, z := returnStringInt()

	fmt.Println(x)
	fmt.Println(y, z)
}

func returnInt() int {
	return 1
}

func returnStringInt() (string, int) {
	return "string", 1
}
