package main

import "fmt"

type person struct {
	name      string
	last_name string
	age       int
}

func (p person) showNameAndAge() (string, int) {
	return fmt.Sprintf("%v %v", p.name, p.last_name), p.age
}

func main() {
	ubiratan := person{"Ubiratan", "Junior", 22}
	complete_name, age := ubiratan.showNameAndAge()

	fmt.Printf("Name: %v \t Age: %v\n", complete_name, age)
}
