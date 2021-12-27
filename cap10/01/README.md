Exercise description:

1. Create a `struct` named people with the following fields:
    - name
    - last_name
    - favourite_ice_cream

1. Create 2 values of people type and show the values, use `range` over the slice of ice creams to print the values

> Code:
```go
package main

import "fmt"

type people struct {
	name             string
	last_name        string
	favourite_ice_cream []string
}

func main() {
	p1 := people{"Ubiratan", "Junior", []string{"chocolate", "flakes"}}
	p2 := people{"Jose", "Costa", []string{"chocolate", "strawberry", "vanilla"}}

	fmt.Println(p1.name, p1.last_name, ":")
	for _, value := range p1.favourite_ice_cream {
		fmt.Printf("\t%v\n", value)
	}

	fmt.Println(p2.name, p2.last_name, ":")
	for _, value := range p2.favourite_ice_cream {
		fmt.Printf("\t%v\n", value)
	}
}
```

> Output:
```console
Ubiratan Junior :
        chocolate
        flakes
Jose Costa :
        chocolate
        strawberry
        vanilla
```