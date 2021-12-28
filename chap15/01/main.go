package main

import (
	"fmt"
)

func main() {
	x := "pointer"
	pointer := &x

	fmt.Println(*pointer)

}
