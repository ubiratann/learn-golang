Exercise description:

1. Create constants typeds and untypeds
2. Show those constants values

> Code:
```go
package main

import "fmt"

const x int = 10
const y = 10

var b float32

func main() {
	a := x
	// 	b = x assignement error
	b = y

	fmt.Printf("a : %v - %T\n", a, a)
	fmt.Printf("b : %v - %T\n", b, b)

}

```

> Output:
```console
a : 10 - int
b : 10 - float32
```