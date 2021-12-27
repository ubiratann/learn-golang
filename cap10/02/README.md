Exercise description:

1. Using the previous solution, put the values of "people" in a map using the `last_name` as `key`
1. Show the values of the map using `range`
1. Show the ice cream lists using another `range` inside the first one

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

	map_ := map[string]people{
		"Junior": people{"Ubiratan", "Junior", []string{"chocolate", "vanilla"}},
		"Costa":  people{"Jose", "Carlos", []string{"vanilla", "flakes"}},
	}

	for _, people := range map_ {
		fmt.Printf("%v:\n", people.name)

		for _, favourite_ice_cream := range people.favourite_ice_cream {
			fmt.Printf("\t%v\n", favourite_ice_cream)
		}
	}
}

```

> Output:
```console
Ubiratan:
        chocolate
        vanilla
Jose:
        vanilla
        flakes
```