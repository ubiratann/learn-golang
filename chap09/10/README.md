Exercise description:

1. Using the previous exercise, remove an entry from map and show all values using `range`

> Code:
```go
package main

import "fmt"

func main() {
	map_ := map[string][]string{}

	map_["Ubiratan_Junior"] = []string{"play games", "listen music"}
	map_["Jose_Pereira"] = []string{"play soccer", "paint"}

	map_["Antonio_Mendes"] = []string{"eat", "dance"}

	for key, value := range map_ {
		fmt.Printf("%v:\n\t%v\n", key, value)
	}

	delete(map_, "Jose_Pereira")
	fmt.Printf("\n\n")
	for key, value := range map_ {
		fmt.Printf("%v:\n\t%v\n", key, value)
	}

}

```

> Output:
```console
Ubiratan_Junior:
        [play games listen music]
Jose_Pereira:
        [play soccer paint]
Antonio_Mendes:
        [eat dance]


Ubiratan_Junior:
        [play games listen music]
Antonio_Mendes:
        [eat dance]
```