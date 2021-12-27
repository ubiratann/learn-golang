Exercise description:

1. Using the previous exercise, add a new entry on map and show all values using `range`

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
```