Exercise description:

1. Create a anonymous struct
1. Chalenge: the struct must have a slice and map field

> Code:
```go
package main

import "fmt"

func main() {
	x := struct {
		name             string
		parents          map[string]string
		favourite_colors []string
	}{"Ubiratan Jr.", map[string]string{"Father": "Ubiratan", "Mother": "Fatima"}, []string{"red", "blue", "green"}}

	fmt.Printf("Name: %v\n", x.name)

	fmt.Printf("\tParents:\n")
	for key, value := range x.parents {
		fmt.Printf("\t\t%v: %v\n", key, value)
	}

	fmt.Println("Favourite colors:")
	for index, color := range x.favourite_colors {
		fmt.Printf("\t%v. %v\n", index, color)
	}
}

```

> Output:
```console
Name: Ubiratan Jr.
        Parents:
                Father: Ubiratan
                Mother: Fatima
Favourite colors:
        0. red
        1. blue
        2. green
```