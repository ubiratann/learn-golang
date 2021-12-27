Exercise description:

1. Create a slice containing string slices (`[][]string`). Assign values to this multi-demensional slice as follow:
	- "name", "surname", "hobby"
1. Include data for 3 people, use `range` to show this datas

> Code:
```go
package main

import "fmt"

func main() {
	slice := [][]string{
		{"Ubiratan", "Júnior", "listen trap"},
		{"Jose", "Pereira", "go to shopping"},
		{"Antonio", "Mendes", "play soccer"},
	}

	for _, value := range slice {
		fmt.Println(value[0])
		for _, value2 := range value {
			fmt.Printf("\t%v\n", value2)
		}
	}
}

```

> Output:
```console
Ubiratan
        Ubiratan
        Júnior
        listen trap
Jose
        Jose
        Pereira
        go to shopping
Antonio
        Antonio
        Mendes
        play soccer
```