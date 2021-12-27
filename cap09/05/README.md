Exercise description:

1. Start with the following slice:
	- `x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}`
1. Using slicing and `append`, create an `y` slice wich have the following values:
	- `[42, 43, 44, 48, 49, 50, 51]`

> Code:
```go
package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y := append(x[:3], x[6:]...)

	fmt.Println(y)
}

```

> Output:
```console
[42 43 44 48 49 50 51]
```