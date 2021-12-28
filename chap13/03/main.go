package main

import "fmt"

func main() {

	fmt.Println("Executing with defer")
	withDefer()
}

func first() {
	fmt.Println("first")
}

func second() {
	fmt.Println("second")
}

func withDefer() {
	first()
	defer second()
	withoutDefer()
}

func withoutDefer() {
	first()
	second()
}
