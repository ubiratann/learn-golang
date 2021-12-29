Exercise description:

1. This exercise will reinforce your knowledge of method sets.
1. Create a type for a struct called "person"
1. Create a "talk" method for this type that has a receiver pointer (*person)
1. Create an interface, "humans", that is implemented by types with the "talk" method
1. Create a "saySomething" function whose parameter is of type "humans" and which calls the "talk" method
1. Demonstrate in your code:
    - That you can use a value of type "*person" in the "saySomething" function
    - That you cannot use a value of type "person" in the "saySomething" function

> Code:
```go
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

```

> Output:
```console
Hello Ubiratan
```