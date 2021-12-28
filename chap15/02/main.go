package main

import "fmt"

type person struct {
	name string
	age  int
}

func changeMe(p *person, newAge int, newName string) {
	p.age = newAge
	p.name = newName
}

func main() {
	ubiratan := person{"Ubiratan", 22}

	fmt.Printf("Initial person values:\n\tAge: %v\n\tName: %v\n", ubiratan.age, ubiratan.name)
	changeMe(&ubiratan, 23, "Bira")
	fmt.Printf("Final person values:\n\tAge: %v\n\tName: %v\n", ubiratan.age, ubiratan.name)
}
