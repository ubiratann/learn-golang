Exercise description:

1. Using `iota`, create 4 constants that these values should be the 4 next years.
1. Show these values.

> Code:
```go
package main

import (
	"fmt"
)

const (
	_ = 2021 + iota
	a
	b
	c
	d
)

func main() {
	fmt.Printf("%v\n%v\n%v\n%v\n", a, b, c, d)
}

```

> Output:
```console
2022
2023
2024
2025
```