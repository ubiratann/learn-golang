Exercise description:

1. Create a value and assign it to a variable
1. Demonstrate the address of this value in memory

> Code:
```go
package main

import (
	"fmt"
)

func main() {
	x := "pointer"
	pointer := &x

	fmt.Println(*pointer)

}

```

> Output:
```console
pointer
```