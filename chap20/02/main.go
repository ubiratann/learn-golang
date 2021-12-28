package main

import "fmt"

type person struct {
	name string
	age  int
}

type human interface {
	talk()
}

func (p *person) talk() {
	fmt.Println("Hello", p.name)
}

func saySomething(h human) {
	h.talk()
}

func main() {
	ubiratan := person{"Ubiratan", 22}
	saySomething(&ubiratan)
	// saySomething(ubiratan)
	// Error: person does not implement human (talk method has pointer receiver)
}
