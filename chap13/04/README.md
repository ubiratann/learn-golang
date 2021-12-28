Exercise description:

1. Create a struct type "person" that contains the fields:
    - Name
    - Last name
    - Age

1. Create a method for "person" that demonstrates full name and age;
1. Create a value of type "person";
1. Use the method created to demonstrate this value. 

> Code:
```go
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

```

> Output:
```console
Name: Ubiratan Junior    Age: 22
```